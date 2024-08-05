
# 安装 OTEL-Operator


```bash
helm repo add open-telemetry https://open-telemetry.github.io/opentelemetry-helm-charts
helm fetch open-telemetry/opentelemetry-operator --untar

# 先安装依赖项
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.13.2/cert-manager.yaml


# 我放在了 ~/otel/ 下面
helm upgrade --install opentelemetry-operator  ~/otel/opentelemetry-operator -f ~/otel/opentelemetry-operator/values.yaml --namespace kube-otel --create-namespace
```


# 安装 Jaejer
```bash
# jaejer.yaml 简易版本
apiVersion: apps/v1
kind: Deployment
metadata:
  name: jaeger
  namespace: kube-otel
spec:
  replicas: 1
  selector:
    matchLabels:
      app: jaeger
  template:
    metadata:
      labels:
        app: jaeger
    spec:
      containers:
      - name: jaeger
        image: mirrors-docker.gastudio.cn/jaegertracing/all-in-one:latest
        env:
        - name: COLLECTOR_OTLP_ENABLED
          value: "true"
        ports:
        - containerPort: 16686
        - containerPort: 14268
---
apiVersion: v1
kind: Service
metadata:
  name: jaeger
  namespace: kube-otel
spec:
  selector:
    app: jaeger
  type: ClusterIP
  ports:
    - name: ui
      port: 16686
      targetPort: 16686
    - name: collector
      port: 14268
      targetPort: 14268
    - name: http
      protocol: TCP
      port: 4318
      targetPort: 4318
    - name: grpc
      protocol: TCP
      port: 4317
      targetPort: 4317
```


# 安装 otelcol（colletor）

```bash
$ cat otelcol.yaml
apiVersion: opentelemetry.io/v1beta1
kind: OpenTelemetryCollector
metadata:
  name: otel
spec:
  config:
    receivers:
      otlp:
        protocols:
          grpc:
            endpoint: 0.0.0.0:4317
          http:
            endpoint: 0.0.0.0:4318
    processors:
      memory_limiter:
        check_interval: 1s
        limit_percentage: 75
        spike_limit_percentage: 15
      batch:
        send_batch_size: 10000
        timeout: 10s

    exporters:
      debug: {}
      otlp/jaeger:
        endpoint: "jaeger.kube-otel:4317"
        tls:
          insecure: true
          #    service:
          #      pipelines:
          #        traces:
          #          receivers: [otlp]
          #          processors: [memory_limiter, batch]
          #          exporters: [debug]

    service:
      pipelines:
        traces:
          receivers: [otlp]
          processors: []
          exporters: [debug,otlp/jaeger]
```


# 安装 otelinst

> 用途就是，在指定的名称空间下，加了 annotaion 对应字段的 pod，都会自动注入 initContainer，收集 Traces 到 Collector
>
```bash
$ cat otel-inst.yaml
apiVersion: opentelemetry.io/v1alpha1
kind: Instrumentation
metadata:
  name: otelinst
  namespace: test-ns # 只在 test-ns 名称空间中生效
spec:
  propagators:
    - tracecontext
    - baggage
    - b3
  sampler:
    type: parentbased_traceidratio
    argument: "1"
  env:
    - name: OTEL_EXPORTER_OTLP_ENDPOINT
      value: otel-collector.middleware:4318
  java:
    env:
      - name: OTEL_EXPORTER_OTLP_ENDPOINT
        value: http://otel-collector.middleware:4317 # collector 地址
```


查看 Instrumentation 资源，能看到 java 程序配置的自动注入的镜像

```bash
$ kubectl describe otelinst otelinst -n test-ns | grep Java -A4
  Java:
    Env:
      Name:   OTEL_EXPORTER_OTLP_ENDPOINT
      Value:  http://otel-collector.middleware:4317
    Image:    ghcr.io/open-telemetry/opentelemetry-operator/autoinstrumentation-java:1.32.1
```


# 业务容器加入注解

```bash
apiVersion: apps/v1
kind: Deployment
metadata:
  name: account
  namespace: {{ .Release.Namespace }}
  labels:
    app: account
spec:
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 5%
      maxSurge: 25%
  selector:
    matchLabels:
      app: account
  template:
    metadata:
      annotations:
        instrumentation.opentelemetry.io/inject-java: "true" # 注解
        prometheus.io/path: "/actuator/prometheus"
        prometheus.io/port: "9102"
        prometheus.io/scrape: "true"
        # cloudnativegame.io/fake-time: "2024-01-01 00:00:00"  # 此处还可以配置时分秒组合的时间间隔，如'3h40s'和'-7h20m40s'， '-'表示过去的时间。
      labels:
        app: account
    spec:
      terminationGracePeriodSeconds: 300
      containers:
        - name: account
          image: "{{ .Values.global.repository }}/account:{{ .Values.account.imageTag  }}"
          imagePullPolicy: Always
          ports:
          - containerPort: 8090
            name: account
          - containerPort: 9102
            name: prometheus
          readinessProbe:
            httpGet:
              path: /account/system/health-check
              port: 8090
            initialDelaySeconds: 30
            periodSeconds: 15
          livenessProbe:
            httpGet:
              path: /account/system/health-check
              port: 8090
            initialDelaySeconds: 90
            periodSeconds: 30
          envFrom:
            - configMapRef:
                name: octopus-cm
          env:
          - name: PODNAME # 用于 jvm_args 的 GC 日志，Kubernetes 自带的 HOSTNAME 变量无效
            valueFrom:
              fieldRef:
                fieldPath: metadata.name
          {{ if eq .Values.global.jvmLevel "sandbox" }}
          - name: xmx
            value: "2g"
          - name: xms
            value: "2g"
          - name: jvm_args
            value: "-XX:MetaspaceSize=512m -XX:MaxMetaspaceSize=512m -XX:MaxDirectMemorySize=512m -XX:+PrintCommandLineFlags -verbose:gc -Xloggc:heap/$(PODNAME)-gc.log -Dio.netty.noPreferDirect=true -XX:+PrintGCDetails -XX:+HeapDumpOnOutOfMemoryError -XX:HeapDumpPath=heap/"
          resources:
            requests:
              cpu: 2
              memory: "6Gi"
            limits:
              cpu: 2
              memory: "6Gi"
          {{ else if eq .Values.global.jvmLevel "low" }}
          - name: xmx
            value: "4g"
          - name: xms
            value: "4g"
          - name: jvm_args
            value: "-XX:MetaspaceSize=512m -XX:MaxMetaspaceSize=512m -XX:MaxDirectMemorySize=1024m -XX:+PrintCommandLineFlags -verbose:gc -Xloggc:heap/$(PODNAME)-gc.log -Dio.netty.noPreferDirect=true -XX:+PrintGCDetails -XX:+HeapDumpOnOutOfMemoryError -XX:HeapDumpPath=heap/"
          resources:
            requests:
              cpu: 4
              memory: "10Gi"
            limits:
              cpu: 4
              memory: "10Gi"
          {{ else if eq .Values.global.jvmLevel "medium" }}
          - name: xmx
            value: "8g"
          - name: xms
            value: "8g"
          - name: jvm_args
            value: "-XX:MetaspaceSize=1024m -XX:MaxMetaspaceSize=1024m -XX:MaxDirectMemorySize=2048m -XX:+PrintCommandLineFlags -verbose:gc -Xloggc:heap/$(PODNAME)-gc.log -Dio.netty.noPreferDirect=true -XX:+PrintGCDetails -XX:+HeapDumpOnOutOfMemoryError -XX:HeapDumpPath=heap/"
          resources:
            requests:
              cpu: 8
              memory: "18Gi"
            limits:
              cpu: 8
              memory: "18Gi"
          {{ else if eq .Values.global.jvmLevel "high" }}
          - name: xmx
            value: "16g"
          - name: xms
            value: "16g"
          - name: jvm_args
            value: "-XX:MetaspaceSize=2048m -XX:MaxMetaspaceSize=2048m -XX:MaxDirectMemorySize=4096m -XX:+PrintCommandLineFlags -verbose:gc -Xloggc:heap/$(PODNAME)-gc.log -Dio.netty.noPreferDirect=true -XX:+PrintGCDetails -XX:+HeapDumpOnOutOfMemoryError -XX:HeapDumpPath=heap/"
          resources:
            requests:
              cpu: 16
              memory: "34Gi"
            limits:
              cpu: 16
              memory: "34Gi"
          {{ else if eq .Values.global.jvmLevel "superhigh" }}
          - name: xmx
            value: "32g"
          - name: xms
            value: "32g"
          - name: jvm_args
            value: "-XX:MetaspaceSize=4096m -XX:MaxMetaspaceSize=4096m -XX:MaxDirectMemorySize=8192m -XX:+PrintCommandLineFlags -verbose:gc -Xloggc:heap/$(PODNAME)-gc.log -Dio.netty.noPreferDirect=true -XX:+PrintGCDetails -XX:+HeapDumpOnOutOfMemoryError -XX:HeapDumpPath=heap/"
          resources:
            requests:
              cpu: 32
              memory: "66Gi"
            limits:
              cpu: 32
              memory: "66Gi"
          {{ else }}
          {{ end }}
          lifecycle:
            preStop:
              exec:
                command: ["/bin/bash","-c","cd /octopus/heap && tar -zcvf ${HOSTNAME}-$(date +%Y%m%d%H%M%S).tar.gz ${HOSTNAME}-gc.log* && rm -rf /octopus/heap/${HOSTNAME}-gc.log* && curl http://localhost:8090/account/internal/deregister && sleep 15 "]
          volumeMounts:
            - name: localtime
              mountPath: /etc/localtime
            - name: timezone
              mountPath: /etc/timezone
            - name: logs
              mountPath: /octopus/logs
            - name: heap
              mountPath: /octopus/heap
            - name: sa
              mountPath: /octopus/sa
            - name: sa
              mountPath: /octopus/logs/default
      volumes:
        - name: localtime
          hostPath:
            path: /etc/localtime
        - name: timezone
          hostPath:
            path: /etc/timezone
        - name: logs
          hostPath:
            path: /data/logs
        - name: heap
          hostPath:
            path: /data/gc
        - name: sa
          hostPath:
            path: /data/Analytic
      restartPolicy: Always
      imagePullSecrets:
        - name: {{ .Values.global.imagePullSecrets }}
      # Pod 反亲和，避免单点故障
      affinity:
        podAntiAffinity:
          preferredDuringSchedulingIgnoredDuringExecution: # 软策略
            - weight: 1
              podAffinityTerm:
                topologyKey: kubernetes.io/hostname # 拓扑域
                labelSelector:
                  matchExpressions:
                    - key: app
                      operator: In
                      values:
                        - account
```


# Zipkin 作为 Exporter 的 Otel collector 相关配置


```yaml
apiVersion: opentelemetry.io/v1beta1
kind: OpenTelemetryCollector
metadata:
  name: otel
spec:
  config:
    receivers:
      otlp:
        protocols:
          grpc:
            endpoint: 0.0.0.0:4317
          http:
            endpoint: 0.0.0.0:4318
    processors:
      memory_limiter:
        check_interval: 1s
        limit_percentage: 75
        spike_limit_percentage: 15
      batch:
        send_batch_size: 10000
        timeout: 10s

    exporters:
      debug: {}
      zipkin:
        endpoint: "http://zipkin.kube-otel:9411/api/v2/spans"
        format: proto
        tls:
          insecure: true

    service:
      pipelines:
        traces:
          receivers: [otlp]
          processors: []
          exporters: [debug,zipkin]
```

![](assets/opentelemetry%20实践/opentelemetry%20实践_image_1.png)

