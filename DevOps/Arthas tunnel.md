

## Dockerfile 文件


```bash
# root @ test in ~/arthas-tunnel [10:40:11] 
$ ls
arthas-tunnel-server-3.7.2-fatjar.jar  Dockerfile

# root @ test in ~/arthas-tunnel [10:40:12] 
$ cat Dockerfile        
FROM registry-dev.gameale.com/roc/jdk21:latest
WORKDIR /root
COPY arthas-tunnel-server-3.7.2-fatjar.jar .
CMD ["java", "-jar", "arthas-tunnel-server-3.7.2-fatjar.jar"]

```

## docker-compose.yaml 文件


```bash
version: '3'
services:
  tunnel:
    image: registry-dev.gameale.com/roc/jdk21:arthas_tunnel
    command: ["java", "-Dserver.port=8080", "-Dspring.security.user.name=arthas", "-Dspring.security.user.password=admin", "-jar", "/root/arthas-tunnel-server-3.7.2-fatjar.jar"]
    ports:
      - "9090:8080"
      - "7777:7777"
```



## 业务连接 Arthas tunnel

### 将 arthas agent 打进业务镜像



> Dockerfile

```bash
FROM registry-dev.gameale.com/roc/jdk21:latest
WORKDIR /root
COPY jar/ .
COPY arthas/ ./arthas
```


### 业务启动 agent 端

 > 此处使用 arthas-boot 启动，可参考 https://arthas.aliyun.com/doc/download.html

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

