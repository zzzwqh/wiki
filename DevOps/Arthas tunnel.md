



## docker-compose


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

```bash
docker exec  ${SID}_${SERVICE}_1 /bin/bash -c "java -jar /root/arthas/arthas-boot.jar 1 --tunnel-server 'ws://10.30.122.173:7777/ws' --agent-id ${SERVICE}-${SID} --attach-only"

```
