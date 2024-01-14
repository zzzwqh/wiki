> **背景**：因需要做滚动升级 Gateway ，滚动升级时 Gateway 前端 SLB 需要调整权重，避免手动调整，使用脚本做调整

[负载均衡虚拟服务器组 API 调试](https://next.api.aliyun.com/api/Slb/2014-05-15/SetVServerGroupAttribute?params={%22RegionId%22:%22cn-qingdao%22,%22VServerGroupId%22:%22rsp-j6cl3dg3gxdn8%22}&tab=DEBUG)
[负载均衡虚拟服务器组代码示例](https://api.alibabacloud.com/api-tools/demo/Slb/8fc8c098-f76e-4fec-a233-e64822f5e70a)
> 根据上面两个地址可以写个简单的工具，每次滚动更新时，使用脚本执行操控 SLB 权重

```go

```





> 配置需要填写如下：
```yaml
accessKeyId: ""  
accessKeySecret: ""  
regionId: "cn-hongkong"  
vServerGroupId: "rsp-abcde"  
backendServers:  
  - serverId: "i-123456789"  
    weight: "100"  
    port: "8081"  
#    description: "" # 选填配置  
  - serverId: "i-123456789"  
    weight: "100"  
    port: "8082"  
#    description: "" # 选填配置
```