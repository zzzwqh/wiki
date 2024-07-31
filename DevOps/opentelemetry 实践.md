
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



查看