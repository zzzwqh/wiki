### 一. 创建测试数据
> 创建测试索引
```bash
curl "http://127.0.0.1:9200/test-idx" \
  -u user:password \
  -s \
  -H "Content-Type: application/json" \
  -X PUT \
  -d '{
    "settings": {
      "number_of_shards": 1
    },
    "mappings": {
      "properties": {
        "@timestamp": { "type": "date" },
        "source": { "type": "keyword" },
        "system": {
          "properties": {
            "syslog": {
              "properties": {
                "hostname": { "type" : "keyword" },
                "message": { "type" : "keyword" }
              }
            }
          }
        }
      }
    }
  }'
```
> 执行后，输出如下
```bash
{"acknowledged":true,"shards_acknowledged":true,"index":"test-idx"}#
```
