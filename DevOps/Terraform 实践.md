
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

> 我需要测试新建 SLB / SLB 监听，添加一个loadbalancer-install.tf 文件，声明 slb 资源

```bash
~/VSCodeDir/terraform » cat loadbalancer-install.tf
resource "alicloud_slb" "instance-1" {
  load_balancer_name   = "terraform-slb-test-1"
  load_balancer_spec   = "slb.s1.small"
  internet_charge_type = "PayByTraffic"
  address_type         = "internet"
}

resource "alicloud_slb" "instance-2" {
  load_balancer_name   = "terraform-slb-test-2"
  load_balancer_spec   = "slb.s1.small"
  internet_charge_type = "PayByTraffic"
  address_type         = "internet"
}%
```

> 再添加一个 listener-install.tf 文件，声明监听资源

```bash
~/VSCodeDir/terraform » cat listener-install.tf
resource "alicloud_slb_listener" "tcp-1" {
  load_balancer_id          = alicloud_slb.instance-1.id
  backend_port              = "22"
  frontend_port             = "22"
  protocol                  = "tcp"
  bandwidth                 = "10"
  health_check_type         = "tcp"
  persistence_timeout       = 3600
  healthy_threshold         = 8
  unhealthy_threshold       = 8
  health_check_timeout      = 8
  health_check_interval     = 5
  health_check_http_code    = "http_2xx"
  health_check_connect_port = 20
  health_check_uri          = "/console"
  established_timeout       = 600
}

resource "alicloud_slb_listener" "tcp-2" {
  load_balancer_id          = alicloud_slb.instance-1.id
  backend_port              = "23"
  frontend_port             = "24"
  protocol                  = "tcp"
  bandwidth                 = "10"
  health_check_type         = "tcp"
  persistence_timeout       = 3600
  healthy_threshold         = 8
  unhealthy_threshold       = 8
  health_check_timeout      = 8
  health_check_interval     = 5
  health_check_http_code    = "http_2xx"
  health_check_connect_port = 20
  health_check_uri          = "/console"
  established_timeout       = 600
}
resource "alicloud_slb_listener" "udp" {
  load_balancer_id          = alicloud_slb.instance-2.id
  backend_port              = 2001
  frontend_port             = 2001
  protocol                  = "udp"
  bandwidth                 = 10
  persistence_timeout       = 3600
  healthy_threshold         = 8
  unhealthy_threshold       = 8
  health_check_timeout      = 8
  health_check_interval     = 4
  health_check_connect_port = 20
}

resource "alicloud_slb_listener" "http" {
  load_balancer_id          = alicloud_slb.instance-2.id
  backend_port              = 80
  frontend_port             = 80
  protocol                  = "http"
  sticky_session            = "on"
  sticky_session_type       = "insert"
  cookie                    = "testslblistenercookie"
  cookie_timeout            = 86400
  health_check              = "on"
  health_check_uri          = "/cons"
  health_check_connect_port = 20
  healthy_threshold         = 8
  unhealthy_threshold       = 8
  health_check_timeout      = 8
  health_check_interval     = 5
  health_check_http_code    = "http_2xx,http_3xx"
  bandwidth                 = 10
  request_timeout           = 80
  idle_timeout              = 30
}
```


> 可以先执行 terraform plan ，能看到执行计划，而未执行变更动作，这里忽略该步骤，直接执行了 terraform apply，执行结果如图

![](assets/Terraform%20实践/Terraform%20实践_image_4.png)

> 可以看到框选的几个部分，终端是创建成功的输出结果返回，另外一个是 terraform.tfstate 文件，这个文件是 terraform 的状态文件，是 terraform 本地自动生成的，它记录了 terraform 所管理的所有资源的当前状态，这个文件很有重要也很有用，在接下来的所有Terraform 运行中，Terraform 会使用这个状态文件来确定需要对基础设施做哪些更改，Terraform 只管理它 "认为" 自己负责的资源，即在 Terraform 的状态文件中有记录的。如果一个资源不是由 Terraform 创建的，或者它的信息没被记录在 Terraform 的状态文件中，Terraform 就不会尝试去管理（即修改或销毁）那个资源。