## 场景一
研发侧提出一个需求，如果 Json 日志中的 Message 字段为 Json 格式，解析 Message 字段的 Json 字段，到 elasticsearch doc 字段中，比如下面格式的日志

```json
{"@timestamp":"2024-05-26T12:32:44.080155025+08:00","@version":"2","message":{"event": "login","user": "john_doe","status": "successful","ip": "192.168.1.1"},"logger_name":"com.gameale.gateway.filter.GatewayLogFilter","thread_name":"reactor-http-epoll-3","level":"INFO","level_value":20000,"sid":"gateway","service":"gateway"}
```

想把 Message 中的 event / login 等等字段解析到文档的根路径下
其实正常情况下，我们配置 filebeat 如下就能够拿到相应的字段到根路径的

```yaml
    # 是否开启 filebeat debug
    #logging.level: debug
    #logging.selectors: ["*"]
    - type: log
      paths:
        - /data/logs/*_game_*.log
      tail_files: false
      fields:
        logfile_type: mail
      close_inactive: 5m
      ignore_older: 10m
      close_timeout: 1h
      symlinks: false
      # 解析 Json 配置
      json.keys_under_root: true
      json.overwrite_keys: true
      json.expand_keys: true
      json.add_error_key: true

```

但是我们服务日志的 Message 字段并非都是 Json 类型，大部分业务 Message 字段都是文本类型，所以导致了冲突，当我讲上面的日志，重定向到日志文件时，filebeat 产生了 droping event 报错日志（ 需要打开 filebeat debug 才能看到日志 ）：

```bash
2024-05-26T12:24:22.175512769+08:00 {"log.level":"debug","@timestamp":"2024-05-26T04:24:22.175Z","log.logger":"elasticsearch","log.origin":{"function":"[github.com/elastic/beats/v7/libbeat/outputs/elasticsearch.(*Client).bulkCollectPublishFails](http://github.com/elastic/beats/v7/libbeat/outputs/elasticsearch.(*Client).bulkCollectPublishFails)","file.name":"elasticsearch/client.go","file.line":455},"message":"Cannot index event publisher.Event{Content:beat.Event{Timestamp:time.Date(2024, time.May, 26, 12, 21, 44, 80155025, time.Location(\"\")), Meta:null, Fields:{\"@version\":\"2\",\"agent\":{\"ephemeral_id\":\"36a8b56d-6555-4b0a-aa10-d1e4b8494e7a\",\"id\":\"1d90cb5a-f1b9-4945-a3f7-9ee5cbef31c8\",\"name\":\"k8s-node-1\",\"type\":\"filebeat\",\"version\":\"8.13.4\"},\"ecs\":{\"version\":\"8.0.0\"},\"fields\":{\"logfile_type\":\"gateway\"},\"host\":{\"name\":\"k8s-node-1\"},\"input\":{\"type\":\"log\"},\"level\":\"INFO\",\"level_value\":20000,\"log\":{\"file\":{\"path\":\"/data/logs/gateway_2024052612.log\"},\"offset\":63981},\"logger_name\":\"com.gameale.gateway.filter.GatewayLogFilter\",\"message\":{\"event\":\"login\",\"ip\":\"[192.168.1.1](http://192.168.1.1)\",\"status\":\"successful\",\"user\":\"john_doe\"},\"service_name\":\"gateway\",\"sid\":\"gateway\",\"thread_name\":\"reactor-http-epoll-3\"}, Private:file.State{Id:\"native::4719267-64529\", PrevId:\"\", Finished:false, Fileinfo:(*os.fileStat)(0xc0013cb5f0), Source:\"/data/logs/gateway_2024052612.log\", Offset:64308, Timestamp:time.Date(2024, time.May, 26, 4, 3, 38, 461957640, [time.Local](http://time.Local)), TTL:-1, Type:\"log\", Meta:map[string]string(nil), FileStateOS:file.StateOS{Inode:0x4802a3, Device:0xfc11}, IdentifierName:\"native\"}, TimeSeries:false}, Flags:0x1, Cache:publisher.EventCache{m:mapstr.M(nil)}} (status=400): {\"type\":\"document_parsing_exception\",\"reason\":\"[1:376] failed to parse field [message] of type [match_only_text] in document with id 'o5Eks48BDOdyS-HhdfqU'. Preview of field's value: '{ip=[192.168.1.1](http://192.168.1.1), event=login, user=john_doe, status=successful}'\",\"caused_by\":{\"type\":\"illegal_state_exception\",\"reason\":\"Can't get text on a START_OBJECT at 1:301\"}}, dropping event!","service.name":"filebeat","ecs.version":"1.6.0"}
```

根据报错得知，Message 在 ES Index Mapping 中是 match_only_text 类型的字段，不能解析 Object ，我们的日志都是收集到一个索引的，这个索引的 Message 已经被自动映射为 text 类型了，所以如果想解析 message 为 Json 字段，除非新建另外一个索引（这样不太合适），或者研发侧将日志的 message 字段换一个名字，比如 message_json ，那么生成的日志格式就会变成如下样例：

```json
{"@timestamp":"2024-05-26T12:46:44.080155025+08:00","@version":"2","message":"auto flush","message_json":{"event": "login","user": "john_doe","status": "successful","ip": "192.168.1.1"},"logger_name":"com.gameale.gateway.filter.GatewayLogFilter","thread_name":"reactor-http-epoll-3","level":"INFO","level_value":20000,"sid":"gateway","service":"gateway"}
```

因为 message_json 是在索引中不存在的，此时 filebeat 将数据传输到 elasticsearch 后会自动识别创建新字段的映射，也就可以在文档的根路径拿到对应值了

![](assets/Json%20日志字段映射冲突/Json%20日志字段映射冲突_image_1.png)




## 场景二

Filebeat 收集日志入 Elasticsearch 时，报错日志如下

```bash
2024-05-26T13:23:42.869300259+08:00 {"log.level":"debug","@timestamp":"2024-05-26T05:23:42.869Z","log.logger":"elasticsearch","log.origin":{"function":"github.com/elastic/beats/v7/libbeat/outputs/elasticsearch.(*Client).bulkCollectPublishFails","file.name":"elasticsearch/client.go","file.line":455},"message":"Cannot index event publisher.Event{Content:beat.Event{Timestamp:time.Date(2024, time.May, 26, 13, 14, 44, 80155025, time.Location(\"\")), Meta:null, Fields:{\"@version\":\"2\",\"agent\":{\"ephemeral_id\":\"9a8a4d00-b475-4ba1-b2bb-3d31da82d841\",\"id\":\"1d90cb5a-f1b9-4945-a3f7-9ee5cbef31c8\",\"name\":\"k8s-node-1\",\"type\":\"filebeat\",\"version\":\"8.13.4\"},\"ecs\":{\"version\":\"8.0.0\"},\"fields\":{\"logfile_type\":\"gateway\"},\"host\":{\"name\":\"k8s-node-1\"},\"input\":{\"type\":\"log\"},\"level\":\"INFO\",\"level_value\":20000,\"log\":{\"file\":{\"path\":\"/data/logs/gateway_2024052613.log\"},\"offset\":83799},\"logger_name\":\"com.gameale.gateway.filter.GatewayLogFilter\",\"message\":\"auto flush\",\"message_json\":{\"event\":{\"eventID\":\"1\",\"eventdescription\":\"test\"},\"ip\":\"192.168.1.1\",\"status\":\"successful\",\"user\":\"john_doe\"},\"service_name\":\"gateway\",\"sid\":\"gateway\",\"thread_name\":\"reactor-http-epoll-3\"}, Private:file.State{Id:\"native::4719269-64529\", PrevId:\"\", Finished:false, Fileinfo:(*os.fileStat)(0xc0011d4410), Source:\"/data/logs/gateway_2024052613.log\", Offset:84188, Timestamp:time.Date(2024, time.May, 26, 5, 21, 32, 867856772, time.Local), TTL:-1, Type:\"log\", Meta:map[string]string(nil), FileStateOS:file.StateOS{Inode:0x4802a5, Device:0xfc11}, IdentifierName:\"native\"}, TimeSeries:false}, Flags:0x1, Cache:publisher.EventCache{m:mapstr.M(nil)}} (status=400): {\"type\":\"document_parsing_exception\",\"reason\":\"[1:540] failed to parse field [message_json.event] of type [keyword] in document with id 'Po5as48BV7hhFmdQyjBC'. Preview of field's value: '{eventID=1, eventdescription=test}'\",\"caused_by\":{\"type\":\"illegal_state_exception\",\"reason\":\"Can't get text on a START_OBJECT at 1:500\"}}, dropping event!","service.name":"filebeat","ecs.version":"1.6.0"}
```

其实我们没有 event 这个字段那为什么 event 会冲突？因为 Filebeat 采用了 ECS（Elasticsearch Common Schema）创建索引