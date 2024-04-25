

游戏服经常要修改时间，因为 QA 需要测试活动 / 未成年登陆等内容是否正常，InitContainer 执行 date -s "xxx" 需要提权，并且提权 SYS_TIME 实际上修改的是容器所在宿主机的节点时间，因为修改时区只能是在 24 小时之内，直接放弃修改时区的想法，想到的方式，是游戏程序启动前置


https://github.com/CloudNativeGame/fake-time-injector?tab=readme-ov-file

