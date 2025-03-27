##  0. 安装
```bash
# 安装环境
apt install openjdk-11-jdk
wget https://dlcdn.apache.org/jmeter/binaries/apache-jmeter-5.6.3.tgz
tar xf apache-jmeter-5.6.3.tgz

# 将下面环境变量，追加到 .zshrc ( 如果没有 zsh ，加入到 /etc/profile )
export HISTTIMEFORMAT='%F %T ' 
export JMETER_HOME=/root/apache-jmeter-5.6.3 
export CLASSPATH=${JMETER_HOME}/lib/ext/ApacheJMeter_core.jar:${JMETER_HOME}/lib/jorphan.jar:${CLASSPATH} 
export PATH=${JMETER_HOME}/bin:$PATH
```
## 1. SSL 相关报错

> 最开始压测命令是这样写的 ...
```bash
[root@iZt4n1igqdpps6gm3lrxaaZ ~]# jmeter -JThreadCount=50000 -JRunTime=30 -Jresponse_timeout=2000 -Jconnect_timeout=2000 -n -t /root/account-stress.jmx -l testx.jtl -e -o ./report
```

> 报错如下，压测机的 Jmeter 配置（ properties 文件 ）需要调整

![](assets/Jmeter%20压测小记/Jmeter%20压测小记_image_1.jpeg)

> 原因是 jmeter 客户端的锅，调整 jmeter 的配置文件，SSL 相关的保持就会消失了
https://www.cnblogs.com/to-here/p/13890622.html


## 2. 无法分配地址报错

> 在执行压测时，有如下报错，这也是压测机需要调整配置
 
 ```cobol
Non HTTP response code: java.net.NoRouteToHostException/Non HTTP response message: Cannot assign requested address (Address not available)
```

> 压测机系统需要调整下面的配置

https://blog.csdn.net/songyun333/article/details/134413242


## 3. 压测报告很多 504 报错
> 上面的都调整好了，但是压测依然会有 504  gateway time out 报错

![](assets/Jmeter%20压测小记/Jmeter%20压测小记_image_2.png)

> 使用的是阿里云 ALB Ingress，去查 ALB Ingress 访问日志，发现 ALB 返回的 504 gateway time out 请求的日志，他们的 upstream response time 都是到了 5s，这个 5s 是后端响应超时时间，ALB 没有在 5s 获取到响应，所以返回给了客户端一个 504 gateway time out

![](assets/Jmeter%20压测小记/Jmeter%20压测小记_image_3.png)

![](assets/Jmeter%20压测小记/Jmeter%20压测小记_image_4.png)

## 4. 服务端丢包迹象

> 查看服务端：netstat -s|egrep -i 'syn|ignore|overflow|reject|becau|error|drop|back'
![](assets/Jmeter%20压测小记/Jmeter%20压测小记_image_5.png)

 >AI 解读下：
 
 ![](assets/Jmeter%20压测小记/Jmeter%20压测小记_image_6.png)



可以看到 tcp accept 队列其实是满了的，暂时不考虑加机器的话可以先调高连接队列，全队列和半队列连接都调高一下，同时开启 tcp_syncookies，然后再压测

> 方法如下，需要在节点上运行

1. 调整半队列值，先 cat /proc/sys/net/ipv4/tcp_max_syn_backlog 查看这个值的大小，然后 使用文本编辑器编辑 /etc/sysctl.conf文件： sudo vi /etc/sysctl.conf 在文件末尾添加或修改以下行，将<新值>替换为您想要设置的新值（可以先设置刚查看值的2倍）： net.ipv4.tcp_max_syn_backlog = <新值> 保存并退出编辑器。 使配置生效：sudo sysctl -p

2. 调整全队列值 先 cat /proc/sys/net/core/somaxconn 查看这个值大小，编辑 /etc/sysctl.conf文件： sudo vi /etc/sysctl.conf 在文件末尾添加或修改如下行，将<新值>替换为您想要设置的新值（可以先设置刚查看值的2倍）： net.core.somaxconn = <新值> 保存并退出编辑器。 使配置生效： sudo sysctl -p 

【 需在非业务高峰期时执行操作，压测场景无需关注 】


## 6. 调整压测姿势

> 其实上面服务端调整，可不执行，调整 TreadCount ，然后做梯度压测才是最好的方式
> 如下命令 + 图中的 iniProp 意思是，10000个线程，分 10s 逐步启动，也就是 1s 起 1000 个，第 10s 起完 10000 个，这就是梯度压测

```bash
jmeter -JThreadCount=10000 -JRunTime=30 -Jresponse_timeout=2000 -Jconnect_timeout=2000 -n -t /root/account-stress.jmx -l testx.jtl -e -o ./report
```

命令还是这个命令，并发数适当减少了下，jmx 是有变化的（ 可以用 jmeter 的 GUI 生成这个 jmx 文件 ）
![](assets/Jmeter%20压测小记/Jmeter%20压测小记_image_7.png)


> 增量压测方式，这个有个 Blog 系列还不错的地址： https://www.cnblogs.com/xiaodi888/p/18152803

![](assets/Jmeter%20压测小记/Jmeter%20压测小记_image_8.png)



## 7. 分布式压测

为什么要分布式压测 ： https://baijiahao.baidu.com/s?id=1784588482195041235&wfr=spider&for=pc
有些分布式压测，使用了 master + slave pod 的方式，放到了 Kubernetes ： https://cloud.tencent.com/developer/article/1808175 

https://zuozewei.blog.csdn.net/article/details/115299107
- 分布式压测，就是指定多个发压 slave 节点，执行压测
- 使用 -R 指定 slave 节点，要注意用 -GThreadCount 而不是 -JThreadCount（其他命令行选项也是），-J 无法将数值下传到各个 slave 节点
- 每个要执行压测的节点，使用 jmeter-server 命令启动（有些 jmeter.properties 配置需要调整，可以百度下）然后用 jmeter 命令去调用 slave 执行压测

```bash
# 压测命令，jmeter 命令去调用 slave 执行压测
# 此处注意
 jmeter -n -t gasdk-pressure-test.jmx -R 10.66.2.46,10.66.2.82 -GThreadCount=4000 -GRampUpTime=10 -GRunTime=300 -GHttp=https -GHost=www.test.com -GPort=443 -l test.jtl -e -o  /data/intranet_report_thread_4000_replicaCount_8_distribute-$(date +%Y%m%d_%H%M%S)
# nginx 配置
server {
    listen 80;

    location /data/ {
        autoindex on;                             # 显示目录列表
        autoindex_exact_size off;                 # 文件大小以KB, MB显示
        autoindex_localtime on;                   # 显示本地时间
        alias /data/;                     # 指定实际的文件目录

    }
}
```

![](assets/Jmeter%20压测小记/Jmeter%20压测小记_image_9.png)

![](assets/Jmeter%20压测小记/Jmeter%20压测小记_image_10.png)







其他：
https://www.cnblogs.com/dreamanddead/p/why-should-you-set-hostname-in-jmeter-distribute-test.html


## 8. 压测标准
1. 一直增加 threadCount 直到报错
2. 观察 请求平均耗时，直到不可承受的情况