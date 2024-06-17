
> 利用 ElasticSearch 本身，可以做监控，但是告警是写到 ES 的（ 邮件什么的需要企业认证的 ）
> 创建一个 Connector ，指定一个 Index，作告警的输出

![](assets/Stack%20Monitoring%20告警/Stack%20Monitoring%20告警_image_1.png)


> 我们配置好了 Stack Monitoring 后，就可以配置默认的告警规则

![](assets/Stack%20Monitoring%20告警/Stack%20Monitoring%20告警_image_2.png)


> 可以选择编辑其中一个告警规则

![](assets/Stack%20Monitoring%20告警/Stack%20Monitoring%20告警_image_3.png)


> 然后选择 Index ， 只有前两个不要钱，不要 Gold License

![](assets/Stack%20Monitoring%20告警/Stack%20Monitoring%20告警_image_4.png)


> Action frequency 可以自由编辑，Document to Index 也可以自由编辑，也可以点击 Index document example 链接，进去看看

![](assets/Stack%20Monitoring%20告警/Stack%20Monitoring%20告警_image_5.png)

```json
{
	"@timestamp": "{{date}}",
	"rule_id": "{{rule.id}}",
	"rule_name": "{{rule.name}}",
	"alert_id": "{{alert.id}}",
	"context_message": "内存使用率超过85% ，请保持关注",
	"rule_url": "{{rule.url}}"
}
```


> 至此，其实告警方式，只是写了文档到一个指定的索引中，至于怎么告警，可以通过 elastalert2 发出，或者自己写一个，此处不再详细介绍

![](assets/Stack%20Monitoring%20告警/Stack%20Monitoring%20告警_image_6.png)

