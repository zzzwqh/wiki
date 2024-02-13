### 一. 语法
- 新增文档
```bash
PUT /megacorp/_doc/1
{
    "first_name" : "John",
    "last_name" :  "Smith",
    "age" :        25,
    "about" :      "I love to go rock climbing",
    "interests": [ "sports", "music" ]
}
PUT /megacorp/_doc/2
{
    "first_name" :  "Jane",
    "last_name" :   "Smith",
    "age" :         32,
    "about" :       "I like to collect rock albums",
    "interests":  [ "music" ]
}

PUT /megacorp/_doc/3
{
    "first_name" :  "Douglas",
    "last_name" :   "Fir",
    "age" :         35,
    "about":        "I like to build cabinets",
    "interests":  [ "forestry" ]
}
```

- 根据 ID 获取文档
```
GET /megacorp/_doc/1
```
- 删除文档
```bash
Delete /megacorp/_doc/2
```
- 搜索当前 Index 所有文档
```bash
GET /megacorp/_search
```