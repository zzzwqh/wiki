Elasticsearch 8.0 版本以后，是必须要配置证书的
查看集群证书过期时间可以用如下命令，elasticsearch operator 的默认维护证书都是 1 年，这不行，证书过期了不是会导致集群不可用吗

```bash
GET /_ssl/certificates
```


----

后来发现，ECK 是实现了证书自动更新的，在 operator.yaml 中，有如下字段配置：


```yaml
...
...
---
# Source: eck-operator/templates/configmap.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: elastic-operator
  namespace: elastic-system
  labels:
    control-plane: elastic-operator
    app.kubernetes.io/version: "2.12.0"
data:
  eck.yaml: |-
    log-verbosity: 0
    metrics-port: 0
    container-registry: docker.elastic.co
    max-concurrent-reconciles: 3
    ca-cert-validity: 8760h
    ca-cert-rotate-before: 24h
    cert-validity: 8760h
    cert-rotate-before: 24h
...
...
```

把上面的证书有效期替换下

- 8760h => 876000h 
- 24h => 8760h


这样证书就是 100 年的了，提前 1 年更新证书，公司黄了，我化成灰，证书都不会出问题


手动更新的方式
https://github.com/elastic/cloud-on-k8s/issues/4675

记录一个好用的工具，自动识别 configmap 变更重载滚动更新 pod 
https://github.com/stakater/Reloader