
## Arthas tunnel 
### Dockerfile 文件


```bash
# root @ test in ~/arthas-tunnel [10:40:11] 
$ ls
arthas-tunnel-server-3.7.2-fatjar.jar  Dockerfile

# root @ test in ~/arthas-tunnel [10:40:12] 
$ cat Dockerfile        
FROM registry.xxxx.com/xxxx/jdk21:latest
WORKDIR /root
COPY arthas-tunnel-server-3.7.2-fatjar.jar .
CMD ["java", "-jar", "arthas-tunnel-server-3.7.2-fatjar.jar"]

```

### docker-compose.yaml 文件


```bash
version: '3'
services:
  tunnel:
    image: xxx.zzzz.com/xxx/jdk21:arthas_tunnel
    command: ["java", "-Dserver.port=8080", "-Dspring.security.user.name=arthas", "-Dspring.security.user.password=admin", "-jar", "/root/arthas-tunnel-server-3.7.2-fatjar.jar"]
    ports:
      - "9090:8080"
      - "7777:7777"
```



## 业务连接 Arthas tunnel

### 将 arthas agent 打进业务镜像

> 下载全量包 （ Maven 仓库下载 ） https://arthas.aliyun.com/doc/download.html

```bash
# 将包中的文件，放入 ~/arthas 路径下
$ ll ~/arthas
total 32M
-rw-r--r-- 1 root root 7.9K Sep 27  2020 arthas-agent.jar
-rw-r--r-- 1 root root 139K Sep 27  2020 arthas-boot.jar
-rw-r--r-- 1 root root 421K Sep 27  2020 arthas-client.jar
-rw-r--r-- 1 root root  13M Sep 27  2020 arthas-core.jar
-rw-r--r-- 1 root root  18M Jun 17 18:36 arthas-packaging-3.7.2-bin.zip
-rw-r--r-- 1 root root  531 Sep 27  2020 arthas.properties
-rw-r--r-- 1 root root 5.0K Sep 27  2020 arthas-spy.jar
-rwxr-xr-x 1 root root 3.1K Sep 27  2020 as.bat
-rwxr-xr-x 1 root root 7.6K Sep 27  2020 as-service.bat
-rwxr-xr-x 1 root root  34K Sep 27  2020 as.sh
drwxr-xr-x 2 root root 4.0K Sep 27  2020 async-profiler
-rwxr-xr-x 1 root root  635 Sep 27  2020 install-local.sh
drwxr-xr-x 2 root root 4.0K Sep 27  2020 lib
-rw-r--r-- 1 root root 2.0K Sep 27  2020 logback.xml
-rw-r--r-- 1 root root 4.1K Sep 27  2020 math-game.jar

```

> Dockerfile 文件如下

```bash
FROM xxx.zzz.com/xxx/jdk21:latest
WORKDIR /root
COPY jar/ .
COPY arthas/ ./arthas
```


### 业务启动 agent 端

 > 此处使用 arthas-boot.jar 启动，可参考 https://arthas.aliyun.com/doc/download.html

```bash
docker exec  ${SID}_${SERVICE}_1 /bin/bash -c "java -jar /root/arthas/arthas-boot.jar 1 --tunnel-server 'ws://10.30.122.173:7777/ws' --agent-id ${SERVICE}-${SID} --attach-only"

```


## 访问 WebUI

> 需要配置 AgentID，连接到业务 JVM 排查问题

```bash
# 查看系统变量
[arthas@1]$  ognl '@java.lang.System@getenv("app_env")'
@String[-Dnacos.config.data.private=game-10101 -Dnacos.config.namespace=dev -Dnacos.config.host=10.30.122.173:8848 -Dnacos.config.username=nacos -Dnacos.config.password=nacos]
```

![](assets/Arthas%20tunnel/Arthas%20tunnel_image_1.png)

