
```bash
# 创建一个测试索引
curl "http://127.0.0.1:9200/test_index" \  
  -u elastic:gameale \  
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

```json
{
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
  }
```