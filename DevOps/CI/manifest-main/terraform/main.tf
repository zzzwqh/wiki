# 要写 AK
provider "alicloud" {
}


variable "transit_router_attachment_name" {
  default = "sdk_rebot"
}

variable "transit_router_attachment_description" {
  default = "sdk_rebot"
}

data "alicloud_cen_transit_router_available_resources" "default" {

}

// VPC
resource "alicloud_vpc" "default-vpc" {
  vpc_name   = "sdk_rebot_cen_tr_yaochi"
  cidr_block = "192.168.0.0/16"
}

// 交换机
resource "alicloud_vswitch" "default-vs" {
  vswitch_name = "sdk_rebot_cen_tr_yaochi"
  vpc_id       = alicloud_vpc.default-vpc.id
  cidr_block   = "192.168.1.0/24"
  zone_id      = data.alicloud_cen_transit_router_available_resources.default.resources[0].master_zones[0]
}

// 安全组
resource "alicloud_security_group" "ethan_group" {
  name        = "terraform-group"
  description = "New security group"
  vpc_id      = alicloud_vpc.default-vpc.id
}
// 安全组规则
resource "alicloud_security_group_rule" "allow_all_tcp_6666" {
  type              = "ingress"
  ip_protocol       = "tcp"
  nic_type          = "intranet"
  policy            = "accept"
  port_range        = "6666/6666"
  priority          = 1
  security_group_id = alicloud_security_group.ethan_group.id
  cidr_ip           = "0.0.0.0/0"
}
