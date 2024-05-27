Elasticsearch 8.0 版本以后，是必须要配置证书的

ECK Operator 默认自动生成的证书有效期是 1 年， 临近 24 小时会更替，ECK 管理的资源（ 例如 Beat ）会更替对应证书，但是非 ECK 管理的资源（ 例如集群外的 ElasticSearch 的客户端 ），不会自动更换证书

查看集群证书过期时间可以用如下命令，elasticsearch operator 的默认维护证书都是 1 年

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


这样证书就是 100 年的了，提前 1 年更新证书，公司黄了，你化成灰，证书都不会出问题 ~~
![](assets/Elasticsearch%20证书问题/Elasticsearch%20证书问题_image_1.png)

集群外的 filebeat 访问集群内 Elasticsearch 需要配置证书
```bash
# 进入 pod 查看 beat.yaml ，外部的 filebeat 可以参考集群内的 filebeat 配置即可
...
output:
    elasticsearch:
        hosts:
            # 此处要注意一点，不能用 IP ，必须要域名，否则证书仍然会报错
            - https://project-es-http.project.svc:9200
        index: project_dev-%{+yyyy.MM.dd}
        password: x
        ssl:
            certificate_authorities:
                - /mnt/elastic-internal/elasticsearch-certs/ca.crt
        username: elastic
...


# 获取 beat 所挂载的证书，拷贝到集群外的 filebeat 部署环境中
kubectl get secret project-beat-es-ca -o jsonpath="{.data.ca\.crt}" | base64 --decode
```




手动更新的方式
https://github.com/elastic/cloud-on-k8s/issues/4675

记录一个好用的工具，自动识别 configmap 变更重载滚动更新 pod 

https://github.com/stakater/Reloader