
#  一. 部署
## 1. 安装 Terraform

> Terraform ： https://developer.hashicorp.com/terraform/install?product_intent=terraform

## 2. 安装 Provider

> Aliyun Provider ： https://registry.terraform.io/providers/aliyun/alicloud/latest/docs

编辑下面文件，然后运行 `terraform init`

```bash
~/VSCodeDir/terraform » cat terraform.tf 
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

> 我选择了使用阿里云 Credentials File 的方式，需要配置 Credential File 路径，
![](assets/Terraform%20实践/Terraform%20实践_image_2.png)

