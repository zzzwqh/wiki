
模板：
```bash
{{ $var := .externalURL}}{{ range $k, $v := .alerts }}{{ if eq $v.status "resolved" }}**<font color="green">Prometheus 恢复通知</font>  **
告警名称：{{$v.labels.alertname}}
告警级别：{{$v.labels.level}}
告警状态：{{ $v.status }}
开始时间：{{GetCSTtime $v.startsAt}}
结束时间：{{GetCSTtime $v.endsAt}}
告警地址：{{$v.labels.instance}}
告警描述：**<font color="green">{{$v.annotations.description}}**{{ else }}</font>**<font color="red">Prometheus 告警通知</font>  **
告警名称：{{$v.labels.alertname}}
告警级别：{{$v.labels.level}}
告警状态：{{ $v.status }}
开始时间：{{GetCSTtime $v.startsAt}}
告警地址：{{$v.labels.instance}}
告警描述：**<font color="red">{{$v.annotations.description}}</font>**{{ end }}{{ end }}
```

配置好后保存
![](assets/prometheusalert/prometheusalert_image_1.png)

可以看到当前模板的 URL，这个 URL 要配置在 alertmanager 里面

![](assets/prometheusalert/prometheusalert_image_2.png)

