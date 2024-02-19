# 利用宿主机命令到容器执行
```bash
# root @ journey-alpha-online-game in /data/filebeat [17:17:36] 
$ docker ps | grep  journey_node-gate-1_1   
fae0bfe74c69   harbor-online.gameale.com/pixel/server:cn_alpha_240219.1   "/bin/sh -c ./bin/ma…"   38 minutes ago      Up 38 minutes      0.0.0.0:8090-8091->8090-8091/tcp, :::8090-8091->8090-8091/tcp, 8092-8093/tcp, 0.0.0.0:8092->8081/tcp, :::8092->8081/tcp   journey_node-gate-1_1

# root @ journey-alpha-online-game in /data/filebeat [17:17:51] 
$ docker inspect --format '{{.State.Pid}}' fae0bfe74c69 
57435

# root @ journey-alpha-online-game in /data/filebeat [17:18:04] 
$  nsenter --target 57435  -n

# root @ journey-alpha-online-game in /data/filebeat [17:18:08] 
$ netstat -lnpt  
Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address           Foreign Address         State       PID/Program name    
tcp        0      0 127.0.0.11:40893        0.0.0.0:*               LISTEN      9138/dockerd        
tcp        0      0 0.0.0.0:8081            0.0.0.0:*               LISTEN      57742/java          
tcp        0      0 0.0.0.0:8011            0.0.0.0:*               LISTEN      57742/java          

```