Elasticsearch 8.0 版本以后，是必须要配置证书的
查看集群证书过期时间可以用如下命令，elasticsearch operator 的默认维护证书都是 1 年，这不行，证书过期了会导致集群不可用

```bash
GET /_ssl/certificates
```


----

自建 Elasticsearch 需要创建一个 100 年的证书，保证公司黄了都不会出现证书问题

https://www.elastic.co/guide/en/cloud-on-k8s/2.12/k8s-webhook.html