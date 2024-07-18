
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

