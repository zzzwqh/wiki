
> 拉取数据大于 10000 条需要分页，分页拿到的数据容易漏掉第一页，当前代码只负责打印到 Console 
```python
from elasticsearch import Elasticsearch
from datetime import datetime, timedelta
import pytz

# 创建 Elasticsearch 客户端
es = Elasticsearch(
    "http://10.72.1.101:9200",
    basic_auth=("exporter", "exporter@xxxxxx")
)

# 定义 UTC+8 时区
shanghai_tz = pytz.timezone('Asia/Shanghai')

# 定义 UTC+8 时间范围的开始时间和结束时间
start_time_shanghai = datetime(2024, 3, 26, 0, 0, 0, tzinfo=shanghai_tz)
end_time_shanghai = datetime(2024, 3, 27, 0, 0, 0, tzinfo=shanghai_tz)

# 转换为 UTC 时间
start_time_utc = start_time_shanghai - timedelta(hours=8)
end_time_utc = end_time_shanghai - timedelta(hours=8)

# 格式化为 Elasticsearch 识别的 ISO8601 格式
start_time_str = start_time_utc.strftime('%Y-%m-%dT%H:%M:%SZ')
end_time_str = end_time_utc.strftime('%Y-%m-%dT%H:%M:%SZ')

# 创建查询体
body = {
    "query": {
        "bool": {
            "must": [
                {"term": {"fields.logfile_type": "clientlog"}},
                {"match": {"message": "content:NullReferenceException: "}},
                {"range": {"@timestamp": {"gte": start_time_str, "lte": end_time_str}}}
            ]
        }
    },
    "sort": [
        {"@timestamp": {"order": "asc"}}  # 按照时间戳升序排序
    ],
    "from": 0,
    "size": 10000  # 每次滚动检索 1000 条文档
}

# 定义滚动搜索的索引和时间
index = "journey_alpha-2024.03.26"
scroll_time = '1m'

# 发起第一个滚动搜索
search_result = es.search(index=index, body=body, scroll=scroll_time)
scroll_id = search_result['_scroll_id']
retrieved_hits = 0
initial_total_hits = search_result['hits']['total']['value']
print("Total hits:", initial_total_hits)

# 打印第一页的内容
hits = search_result['hits']['hits']
for hit in hits:
    timestamp_str = hit['_source']['@timestamp']
    timestamp_dt = datetime.strptime(timestamp_str, '%Y-%m-%dT%H:%M:%S.%fZ')
    utc_dt = timestamp_dt.replace(tzinfo=pytz.utc)
    shanghai_dt = utc_dt.astimezone(shanghai_tz)
    retrieved_hits += 1
    print("Local Time (Asia/Shanghai):", shanghai_dt, "Document:", hit['_source'])

# 更新滚动上下文
scroll_id = search_result['_scroll_id']

# 继续滚动查询并打印剩余的内容
while retrieved_hits < initial_total_hits:
    scroll_result = es.scroll(scroll_id=scroll_id, scroll=scroll_time)
    hits = scroll_result['hits']['hits']
    if not hits:
        break  # 如果没有更多的文档，退出循环
    for hit in hits:
        timestamp_str = hit['_source']['@timestamp']
        timestamp_dt = datetime.strptime(timestamp_str, '%Y-%m-%dT%H:%M:%S.%fZ')
        utc_dt = timestamp_dt.replace(tzinfo=pytz.utc)
        shanghai_dt = utc_dt.astimezone(shanghai_tz)
        retrieved_hits += 1
        print("Local Time (Asia/Shanghai):", shanghai_dt, "Document:", hit['_source'])
    # 更新滚动上下文
    scroll_id = scroll_result['_scroll_id']
```