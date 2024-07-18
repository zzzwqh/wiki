
## SSL 相关报错

> 最开始压测命令是这样写的 ...
```bash
[root@iZt4n1igqdpps6gm3lrxaaZ ~]# jmeter -JThreadCount=30000 -JRunTime=30 -Jresponse_timeout=2000 -Jconnect_timeout=2000 -n -t /root/account-stress.jmx -l testx.jtl -e -o ./report
```

> 报错如下，压测机的 Jmeter 配置（ properties 文件 ）需要调整

![](assets/Jmeter%20压测小记/Jmeter%20压测小记_image_1.jpeg)

> 原因是 jmeter 客户端的锅，调整 jmeter 的配置文件，SSL 相关的保持就会消失了
https://www.cnblogs.com/to-here/p/13890622.html


## 无法分配地址报错

> 在执行压测时，有如下报错，这也是压测机需要调整配置
 
 ```cobol
Non HTTP response code: java.net.NoRouteToHostException/Non HTTP response message: Cannot assign requested address (Address not available)
```

> 压测机系统需要调整下面的配置

https://blog.csdn.net/songyun333/article/details/134413242


## 504 报错
> 上面的都调整好了，但是压测依然会有 504  gateway time out 报错

![](assets/Jmeter%20压测小记/Jmeter%20压测小记_image_2.png)

> 使用的是阿里云 ALB Ingress，去查 ALB Ingress 访问日志，发现 ALB 返回的 504 gateway time out 请求的日志，他们的 upstream response time 都是到了 5s，这个 5s 是后端响应超时时间，ALB 没有在 5s 获取到响应，所以返回给了客户端一个 504 gateway time out

![](assets/Jmeter%20压测小记/Jmeter%20压测小记_image_3.png)

![](assets/Jmeter%20压测小记/Jmeter%20压测小记_image_4.png)

## 服务端迹象

> 查看服务端：netstat -s|egrep -i 'syn|ignore|overflow|reject|becau|error|drop|back'
![](assets/Jmeter%20压测小记/Jmeter%20压测小记_image_5.png)

 >AI 解读下：
 
 ![](assets/Jmeter%20压测小记/Jmeter%20压测小记_image_6.png)



可以看到 tcp accept 队列其实是满了的，暂时不考虑加机器的话可以先调高连接队列，全队列和半队列连接都调高一下，同时开启 tcp_syncookies，然后再压测

> 方法如下，需要在节点上运行

1.调整半队列值，先 cat /proc/sys/net/ipv4/tcp_max_syn_backlog 查看这个值的大小，然后 使用文本编辑器编辑 /etc/sysctl.conf文件： sudo vi /etc/sysctl.conf 在文件末尾添加或修改以下行，将<新值>替换为您想要设置的新值（可以先设置刚查看值的2倍）： net.ipv4.tcp_max_syn_backlog = <新值> 保存并退出编辑器。 使配置生效：sudo sysctl -p

2.调整全队列值 先 cat /proc/sys/net/core/somaxconn 查看这个值大小，编辑 /etc/sysctl.conf文件： sudo vi /etc/sysctl.conf 在文件末尾添加或修改如下行，将<新值>替换为您想要设置的新值（可以先设置刚查看值的2倍）： net.core.somaxconn = <新值> 保存并退出编辑器。 使配置生效： sudo sysctl -p 

【 需在非业务高峰期时执行操作，压测场景无需关注 】


## 调整压测姿势

> 其实上面服务端调整，可不执行，调整 TreadCount ，然后做梯度压测才是最好的方式

```bash
jmeter -JThreadCount=30000 -JRunTime=30 -Jresponse_timeout=2000 -Jconnect_timeout=2000 -n -t /root/account-stress.jmx -l testx.jtl -e -o ./report
```

命令还是这个命令，jmx 是有变化的（ 可以用 jmeter 的 GUI 生成这个 jmx 文件 ）
![](assets/Jmeter%20压测小记/Jmeter%20压测小记_image_7.png)


> 增量压测方式，这个有个 Blog 系列还不错的地址： https://www.cnblogs.com/xiaodi888/p/18152803


![](assets/Jmeter%20压测小记/Jmeter%20压测小记_image_8.png)

