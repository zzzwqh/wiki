
#  一. 安装部署
## 1. 安装 Terraform

> Terraform ： https://developer.hashicorp.com/terraform/install?product_intent=terraform

## 2. 安装 Provider

> Aliyun Provider ： https://registry.terraform.io/providers/aliyun/alicloud/latest/docs

编辑下面文件，然后运行 `terraform init`

```bash
~/VSCodeDir/terraform » cat base.tf 
terraform {
  required_providers {
    alicloud = {
      source = "aliyun/alicloud"
      version = "1.214.1"
    }
  }
}

provider "alicloud" {
  # Configuration options
}
```

执行后效果如图所示

![](assets/Terraform%20实践/Terraform%20实践_image_1.png)
## 3. 阿里云配置

> 配置阿里云 Provider 需要做认证授权，当前本地有 aliyun CLI 生成的授权文件如下：


![](assets/Terraform%20实践/Terraform%20实践_image_2.png)


> 所以我直接选择了使用阿里云 Credentials File 的方式，需要配置 Credential File 路径


![](assets/Terraform%20实践/Terraform%20实践_image_3.png)

> 这个文件生成方式参考阿里云官方文档 https://www.alibabacloud.com/help/zh/alibaba-cloud-cli/latest/overview ，更改后的代码如下所示：


```bash
~/VSCodeDir/terraform » cat base.tf 
terraform {
  required_providers {
    alicloud = {
      source = "aliyun/alicloud"
      version = "1.214.1"
    }
  }
}

provider "alicloud" {
  region                  = "cn-hangzhou"
  # 使用 aliyun CLI 工具生成的 Credentials File 
  shared_credentials_file = "/Users/wangqihan-020037/.aliyun/config.json"
  # 注意要和 aliyun configure list 中的 Profile 名字一致
  profile                 = "akProfile"
}

}
```

> 我需要测试新建 SLB，则需要添加一个 slb-install.tf 文件

```bash
~/VSCodeDir/terraform » cat slb-install.tf
resource "alicloud_slb" "instance" {
  load_balancer_name   = "terraform-slb-test"
  load_balancer_spec   = "slb.s2.small"
  internet_charge_type = "PayByTraffic"
  address_type         = "internet"
}
```

> 可以先执行 terraform plan ，能看到执行计划，而未执行变更动作，这里忽略该步骤，直接执行了 terraform apply

![](assets/Terraform%20实践/Terraform%20实践_image_4.png)








