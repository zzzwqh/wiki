Elasticsearch 8.0 版本以后，是必须要配置证书的
查看集群证书过期时间可以用如下命令，elasticsearch operator 的默认维护证书都是 1 年，这不行，证书过期了不是会导致集群不可用吗

```bash
GET /_ssl/certificates
```


----

后来发现，ECK 是实现了

鄙人认为，自建 Elasticsearch 需要一个 100 年的证书，保证公司黄了都不会出现证书问题