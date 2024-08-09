

```bash
docker exec  ${SID}_${SERVICE}_1 /bin/bash -c "java -jar /root/arthas/arthas-boot.jar 1 --tunnel-server 'ws://10.30.122.173:7777/ws' --agent-id ${SERVICE}-${SID} --attach-only"

```
