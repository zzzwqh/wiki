
> 在 ECK 中有 MietricBeat CRD 资源类型

```yaml
---
# metricbeat resources
apiVersion: beat.k8s.elastic.co/v1beta1
kind: Beat
metadata:
  name: metricbeat
  namespace: middleware
spec:
  type: metricbeat
  version: 8.13.4
  elasticsearchRef:
    # 填写好对应 es 资源的名字
    name: elasticsearch
  config:
    metricbeat:
      autodiscover:
        providers:
          - type: kubernetes
            scope: cluster
            hints.enabled: true
            templates:
              - condition:
                  contains:
                    kubernetes.labels.scrape: es
                config:
                  - module: elasticsearch
                    metricsets:
                      - ccr
                      - cluster_stats
                      - enrich
                      - index
                      - index_recovery
                      - index_summary
                      - node_stats
                      - shard
                    period: 10s
                    hosts: "https://${data.host}:${data.ports.https}"
                    username: ${MONITORED_ES_USERNAME}
                    password: ${MONITORED_ES_PASSWORD}
                    # WARNING: disables TLS as the default certificate is not valid for the pod FQDN
                    # TODO: switch this to "certificate" when available: https://github.com/elastic/beats/issues/8164
                    # 这里需要改成 certificate 
                    ssl.verification_mode: "certificate"
                    # 证书路径需要写一下，使用 ECK CRD 创建，这个是固定路径
                    ssl.certificate_authorities:
                      - /mnt/elastic-internal/elasticsearch-certs/ca.crt
                    xpack.enabled: true
              - condition:
                  contains:
                    kubernetes.labels.scrape: kb
                config:
                  - module: kibana
                    metricsets:
                      - stats
                    period: 10s
                    hosts: "https://${data.host}:${data.ports.https}"
                    username: ${MONITORED_ES_USERNAME}
                    password: ${MONITORED_ES_PASSWORD}
                    # WARNING: disables TLS as the default certificate is not valid for the pod FQDN
                    # TODO: switch this to "certificate" when available: https://github.com/elastic/beats/issues/8164
                    ssl.verification_mode: "certificate"
                    ssl.certificate_authorities:
                      - /mnt/elastic-internal/elasticsearch-certs/ca.crt
                    xpack.enabled: true
    processors:
    - add_cloud_metadata: {}
    logging.json: true
  deployment:
    podTemplate:
      spec:
        serviceAccountName: metricbeat
        automountServiceAccountToken: true
        # required to read /etc/beat.yml
        securityContext:
          runAsUser: 0
        containers:
        - name: metricbeat
          env:
          - name: MONITORED_ES_USERNAME
            value: elastic
          - name: MONITORED_ES_PASSWORD
            valueFrom:
              secretKeyRef:
                key: elastic
                name: elasticsearch-es-elastic-user
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: metricbeat
rules:
- apiGroups: [""] # "" indicates the core API group
  resources:
  - namespaces
  - pods
  - nodes
  verbs:
  - get
  - watch
  - list
- apiGroups: ["apps"]
  resources:
  - replicasets
  verbs:
  - get
  - list
  - watch
- apiGroups: ["batch"]
  resources:
  - jobs
  verbs:
  - get
  - list
  - watch
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: metricbeat
  namespace: middleware
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: metricbeat
subjects:
- kind: ServiceAccount
  name: metricbeat
  namespace: middleware
roleRef:
  kind: ClusterRole
  name: metricbeat
  apiGroup: rbac.authorization.k8s.io
```

![](assets/MetricBeat%20云原生监控/MetricBeat%20云原生监控_image_1.png)

