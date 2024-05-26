研发侧提出一个需求，如果 Json 日志中的 Message 字段为 Json 格式，解析 Message 字段的 Json 字段，到 elasticsearch doc 字段中，比如下面格式的日志：
```json
{"@timestamp":"2024-05-26T12:32:44.080155025+08:00","@version":"2","message":{"event": "login","user": "john_doe","status": "successful","ip": "192.168.1.1"},"logger_name":"com.gameale.gateway.filter.GatewayLogFilter","thread_name":"reactor-http-epoll-3","level":"INFO","level_value":20000,"sid":"gateway","service":"gateway"}
```
想把 message 中的 event / login 等等字段解析到文档的根路径下

但是我们服务日志的 message 字段并非都是 Json 类型，也有文本类型（text），所以使用
