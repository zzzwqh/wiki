> **背景**：因需要做滚动升级 Gateway ，滚动升级时 Gateway 前端 SLB 需要调整权重，避免手动调整，使用脚本做调整

[负载均衡虚拟服务器组 API 调试](https://next.api.aliyun.com/api/Slb/2014-05-15/SetVServerGroupAttribute?params={%22RegionId%22:%22cn-qingdao%22,%22VServerGroupId%22:%22rsp-j6cl3dg3gxdn8%22}&tab=DEBUG)
[负载均衡虚拟服务器组代码示例](https://api.alibabacloud.com/api-tools/demo/Slb/8fc8c098-f76e-4fec-a233-e64822f5e70a)
> 根据上面两个地址可以写个简单的工具，每次滚动更新时，使用脚本执行操控 SLB 权重


## 一. 脚本代码
> 需要指定配置文件，填入
```go
package main
import (  
    "encoding/json"  
    "flag"
    "fmt"
    "gameale-ops/config"
    openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"  
    slb "github.com/alibabacloud-go/slb-20140515/v4/client"  
    util "github.com/alibabacloud-go/tea-utils/v2/service"  
    "github.com/alibabacloud-go/tea/tea"    
    "strings"
    )
/**  
* 使用 AK&SK 初始化账号 Client 
* @param accessKeyId 
* @param accessKeySecret 
* @return Client 
* @throws Exception
* */  
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *slb.Client, err error) {  
    config := &openapi.Config{  
       AccessKeyId:     accessKeyId,  
       AccessKeySecret: accessKeySecret,  
    }  
    config.Endpoint = tea.String("slb.aliyuncs.com")  
    _result = &slb.Client{}  
    _result, err = slb.NewClient(config)  
    return _result, err  
}  
  
func main() {  
    weight0 := flag.Int("weight0", -1, "Weight for the first server (0-100)")  
    weight1 := flag.Int("weight1", -1, "Weight for the second server (0-100)")  
    configPath := flag.String("config", "./etc/slb_config.yaml", "Path to the SLB config file")  
    // 解析命令行参数  
    flag.Parse()  
    // 加载配置文件  
    var conf config.SLBConfig  
    err := conf.LoadFromYAML(*configPath)  
    if err != nil {  
       panic(err)  
    }  
    // 声明实际传入参数  
    var weight0Str string  
    var weight1Str string  
    // 如果命令行传入参数，使用命令行参数，否则使用配置文件参数  
    if *weight0 != -1 && *weight1 != -1 {  
       weight0Str = fmt.Sprintf("%d", *weight0)  
       weight1Str = fmt.Sprintf("%d", *weight1)  
    } else {  
       // 使用配置文件参数  
       weight0Str = fmt.Sprintf("%v", conf.BackendServers[0].Weight)  
       weight1Str = fmt.Sprintf("%v", conf.BackendServers[1].Weight)  
    }  
  
    backendServers := `[  
{ "ServerId": "` + conf.BackendServers[0].ServerId + `", "Weight": "` + weight0Str + `",  
"Type": "ecs",  
"Port":"` + conf.BackendServers[0].Port + `",  
"Description":"` + conf.BackendServers[0].Description + `" },  
  
{ "ServerId": "` + conf.BackendServers[1].ServerId + `", "Weight": "` + weight1Str + `",  
"Type": "ecs",  
"Port":"` + conf.BackendServers[1].Port + `",  
"Description":"` + conf.BackendServers[1].Description + `" },  
]`  
    fmt.Println(backendServers)  
    client, err := CreateClient(tea.String(conf.AccessKeyId), tea.String(conf.AccessKeySecret))  
    if err != nil {  
       panic(err)  
    }  
  
    setVServerGroupAttributeRequest := &slb.SetVServerGroupAttributeRequest{  
       RegionId:       tea.String(conf.RegionId),  
       VServerGroupId: tea.String(conf.VServerGroupId),  
       BackendServers: tea.String(backendServers),  
    }  
  
    runtime := &util.RuntimeOptions{}  
    tryErr := func() (_e error) {  
       defer func() {  
          if r := tea.Recover(recover()); r != nil {  
             _e = r  
          }  
       }()  
       // 复制代码运行请自行打印 API 的返回值  
       res, err := client.SetVServerGroupAttributeWithOptions(setVServerGroupAttributeRequest, runtime)  
       if err != nil {  
          return err  
       }  
       fmt.Println(res)  
       return nil  
    }()  
  
    if tryErr != nil {  
       var error = &tea.SDKError{}  
       if _t, ok := tryErr.(*tea.SDKError); ok {  
          error = _t  
       } else {  
          error.Message = tea.String(tryErr.Error())  
       }  
       // 错误 message       fmt.Println(tea.StringValue(error.Message))  
       // 诊断地址  
       var data interface{}  
       d := json.NewDecoder(strings.NewReader(tea.StringValue(error.Data)))  
       d.Decode(&data)  
       if m, ok := data.(map[string]interface{}); ok {  
          recommend, _ := m["Recommend"]  
          fmt.Println(recommend)  
       }  
       _, err = util.AssertAsString(error.Message)  
       if err != nil {  
          panic(err)  
       }  
    }  
}
```

## 二. 配置结构体
```go
package config  
  
import (  
    "gopkg.in/yaml.v2"  
    "io/ioutil")  
  
type SLBConfig struct {  
    AccessKeyId     string `yaml:"accessKeyId"`  
    AccessKeySecret string `yaml:"accessKeySecret"`  
    RegionId        string `yaml:"regionId"`  
    VServerGroupId  string `yaml:"vServerGroupId"`  
    BackendServers  []struct {  
       ServerId    string `yaml:"serverId"`  
       Weight      string `yaml:"weight"`  
       Port        string `yaml:"port"`  
       Description string `yaml:"description"`  
    } `yaml:"backendServers"`  
}  
  
func (c *SLBConfig) LoadFromYAML(path string) error {  
    yamlFile, err := ioutil.ReadFile(path)  
    if err != nil {  
       return err  
    }  
    err = yaml.Unmarshal(yamlFile, c)  
    if err != nil {  
       return err  
    }  
    return nil  
}
```

## 三. 配置文件
> 需要填写如下的云资源信息，当前负载均衡的虚拟服务器组只有两个后端 Server，如果有多个，需要调用创建服务器后端的方法，`client.SetVServerGroupAttributeWithOptions`  是无法控制新增虚拟服务器到虚拟服务器组的
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