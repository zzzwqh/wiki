对于 decode_json_fields 的字段理解可能比较容易，但是 expand_keys 我配置感觉没有什么效果，官网文档链接和解释如下：
https://www.elastic.co/guide/en/beats/filebeat/8.13/decode-json-fields.html
(Optional) A Boolean value that specifies whether keys in the decoded JSON should be recursively de-dotted and expanded into a hierarchical object structure. For example, `{"a.b.c": 123}` would be expanded into `{"a":{"b":{"c":123}}}`.



## 开启 expand_keys 的场景

此处收集的日志格式案例如下所示：

```json
{"@timestamp":"2024-05-26T13:41:44.080155025+08:00","@version":"2","message":"auto flush","message_json.customize.test.field":{"events1": {"eventID":"1","eventdescription":"test"},"user": "john_doe","status": "successful","ip": "192.168.1.1"},"logger_name":"com.gameale.gateway.filter.GatewayLogFilter","thread_name":"reactor-http-epoll-3","level":"INFO","level_value":20000,"sid":"gateway","service":"gateway"}
```


我在 filebeat.yaml 中配置这个字段后，一直感觉没什么效果，其实官网给的解释很清晰，但是在 Kibana 查询后文档中显示其实没什么出入

![](assets/expand_keys%20作用/expand_keys%20作用_image_1.png)


但是在索引 Mapping 中，我们就能观察到这个配置生效后的作用，已经将字段分割，如下所见

![](assets/expand_keys%20作用/expand_keys%20作用_image_2.png)

## 未开启 expand_keys 的场景

此处收集的日志格式如下所示：

```bash
{"@timestamp":"2024-05-26T13:51:44.080155025+08:00","@version":"2","message":"auto flush","message_test.customize.test.field":{"events1": {"eventID":"1","eventdescription":"test"},"user": "john_doe","status": "successful","ip": "192.168.1.1"},"logger_name":"com.gameale.gateway.filter.GatewayLogFilter","thread_name":"reactor-http-epoll-3","level":"INFO","level_value":20000,"sid":"gateway","service":"gateway"}
```