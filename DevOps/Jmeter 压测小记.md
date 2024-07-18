
## SS
> 最开始压测命令是这样写的 ...
```bash
[root@iZt4n1igqdpps6gm3lrxaaZ ~]# jmeter -JThreadCount=30000 -JRunTime=30 -Jresponse_timeout=2000 -Jconnect_timeout=2000 -n -t /root/account-stress.jmx -l testx.jtl -e -o ./report
```

> 直接很多报错，让我费解的是 SSL 相关怎么会有报错，报错如下

![](assets/Jmeter%20压测小记/Jmeter%20压测小记_image_1.jpeg)

原因是 jmeter 客户端的锅，调整 jmeter 的配置文件，SSL 相关的保持就会消失了
https://www.cnblogs.com/to-here/p/13890622.html

