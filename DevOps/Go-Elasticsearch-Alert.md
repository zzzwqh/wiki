## 一. 创建测试数据
- 创建测试索引 test-idx
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
> 执行后，正常返回输出如下
```bash
{"acknowledged":true,"shards_acknowledged":true,"index":"test-idx"}
```
- 插入测试文档
```bash
# 清除测试数据
rm /tmp/gea-payload-*.json
```

## 