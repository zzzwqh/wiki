
游戏服经常要修改时间，因为 QA 需要测试活动 / 未成年登陆等内容是否正常，最开始想到了几个方式，因为修改时区只能是在 24 小时范围之内，直接放弃修改时区的想法
- initContainer 执行 date -s "xxx" 需要提权，并且提权 SYS_TIME 实际上修改的是容器所在宿主机的节点时间，但是 Kubernetes 集群，我们只能同时修改所有节点的时间
- 游戏程序启动前置脚本检查，如果获取到环境变量 TimeController: enabled , TimeSet: "2024-01-01 10:00:30" 这两个字段，就在容器内去执行，但是实际也要提权，修改了节点时间

所以综上所述，并没有好方法，只能开个新集群，给时间服用，后来发现了下面的工具，只需要加 annotation 就可以，是真好用
https://github.com/CloudNativeGame/fake-time-injector?tab=readme-ov-file

