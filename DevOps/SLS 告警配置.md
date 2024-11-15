因为研发一直使用 elasticsearch / kibana 查询日志，所以配置了 kibana 查询 SLS 日志，日志告警需求配置 Kibana 跳转链接，所以需要自定义 Kibana 域名 / Index Pattern Name / KQL，然后在告警卡片中，呈现出来

告警配置，其中 kibana_url 配置如下，镶嵌了几个变量，这几个变量在添加标签中定义，如图所示
```
kibana_url = https://${kibana_domain}/app/discover#/?_g=(filters:!(),refreshInterval:(pause:!t,value:0),time:(from:now-1h,to:now))&_a=(columns:!(),filters:!(),index:${index_name},interval:auto,query:(language:kuery,query:'${kql}'),sort:!(!('@timestamp',desc)))
``` 

![](assets/SLS%20告警配置/SLS%20告警配置_image_1.png)


内容模板如下，可以根据具体查询的字段，配置告警模板，查询到的字段都包含在 annotations 中：
```bash
**告警名称**：**<font color="blue">{{ alert.alert_name }}</font>**
**触发时间**：{{ alert.fire_time | format_date }}
**告警时间**：{{ alert.alert_time | format_date }}
**告警服务**：{{ alert.annotations.service }} {{ alert.annotations.sid }}
**告警模块**：{{ alert.annotations.logger_name }}
**命中数量**：{{ alert.annotations.__count__ }}
**告警详细**：{{ alert.annotations.message }}
**其他链接**：{% if alert.query_url -%} [[登陆查询]]({{ alert.query_url }}){% endif -%}  {% if alert.alert_url -%}｜[[告警设置]]({{ alert.alert_url }}) {% endif -%}｜[[Kibana 查询链接]]({{ alert.annotations.kibana_url }})

```
![](assets/SLS%20告警配置/SLS%20告警配置_image_2.png)


最终效果：
![](assets/SLS%20告警配置/SLS%20告警配置_image_3.png)

