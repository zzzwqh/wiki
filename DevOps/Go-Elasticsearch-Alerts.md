> 公司的 elastalert2 被人魔改过，不太好用，想用 go 重新写一个，找到一个 github 上的项目，需要做下飞书通知等改进，先测试下，项目地址：

```bash
 git clone git@github.com:morningconsult/go-elasticsearch-alerts.git 
```


## 一. 创建测试数据
#### 1. 创建测试索引 test-idx

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
#### 2. 插入测试文档
```bash
# 清除测试数据
rm /tmp/gea-payload-*.json
NOW="$( date +%s )000"  
cat <<EOF > /tmp/gea-payload-1.json  
{  
  "@timestamp": "${NOW}",  
  "source": "/var/log/system.log",  
  "system": {  
    "syslog": {  
      "hostname": "ip-127-0-0-1",  
      "message": "You got an error buddy!",  
      "queue_size": {  
        "value": 60  
      }  
    }  
  }  
}  
EOF
  
cat <<EOF > /tmp/gea-payload-2.json  
{  
  "@timestamp": "${NOW}",  
  "source": "/var/log/errors.log",  
  "system": {  
    "syslog": {  
      "hostname": "ip-127-0-0-1",  
      "message": "Another error!",  
      "queue_size": {  
        "value": 59  
      }  
    }  
  }  
}  
EOF

for f in /tmp/gea-payload-*.json; do
  # Write a document to the new index
  curl "http://127.0.0.1:9200/test-idx/_doc" \
    -s \
    -H "Content-Type: application/json" \
    -X POST \
    -u elastic:gameale \
    -d "@${f}"
done
```
> 正常返回输出如下
```bash
{"_index":"test-idx","_type":"_doc","_id":"2IV-kI0BOEpJpMsp4jQW","_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":6,"_primary_term":1}{"_index":"test-idx","_type":"_doc","_id":"2YV-kI0BOEpJpMsp4jQf","_version":1,"result":"created","_shards":{"total":2,"successful":1,"failed":0},"_seq_no":7,"_primary_term":1}#
```
## 