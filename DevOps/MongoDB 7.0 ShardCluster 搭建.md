>参考 https://blog.csdn.net/sweet6hero/article/details/132861202
  当前需求搭建 2 个 mongos 节点 / 2 个 shard 节点，需要改动下 docker-compose.yaml
```yaml
version: '3'
services:
  mongo_shard1:
    image: mongo:7.0.7
    container_name: mongo_shard1
    # --shardsvr: 这个参数仅仅只是将默认的27017端口改为27018,如果指定--port参数，可用不需要这个参数
    # --directoryperdb：每个数据库使用单独的文件夹
    command: mongod --shardsvr --directoryperdb --replSet mongo_shard1 --keyFile /data/mongo-keyfile
    ports:
      - 27021:27018
    volumes:
      - /etc/localtime:/etc/localtime
      - /disk1/shard1:/data/db
      - ./mongo-keyfile:/data/mongo-keyfile
    privileged: true

  mongo_shard2:
    image: mongo:7.0.7
    container_name: mongo_shard2
    command: mongod --shardsvr --directoryperdb --replSet mongo_shard2 --keyFile /data/mongo-keyfile
    ports:
      - 27022:27018
    volumes:
      - /etc/localtime:/etc/localtime
      - /disk2/shard2:/data/db
      - ./mongo-keyfile:/data/mongo-keyfile
    privileged: true

#  mongo_shard3:
#    image: mongo:7.0.7
#    container_name: mongo_shard3
#    command: mongod --shardsvr --directoryperdb --replSet mongo_shard3 --keyFile /data/mongo-keyfile
#    ports:
#      - 27023:27018
#    volumes:
#      - /etc/localtime:/etc/localtime
#      - ./data/shard3:/data/db
#      - ./mongo-keyfile:/data/mongo-keyfile
#    privileged: true

  mongo_config1:
    image: mongo:7.0.7
    container_name: mongo_config1
    # --configsvr: 这个参数仅仅是将默认端口由27017改为27019, 如果指定--port可不添加该参数
    command: mongod --configsvr --directoryperdb --replSet fates-mongo-config  --keyFile /data/mongo-keyfile
    ports:
      - 27031:27019
    volumes:
      - /etc/localtime:/etc/localtime
      - /disk3/mongo-config1:/data/configdb
      - ./mongo-keyfile:/data/mongo-keyfile

  mongo_config2:
    image: mongo:7.0.7
    container_name: mongo_config2
    command: mongod --configsvr --directoryperdb --replSet fates-mongo-config  --keyFile /data/mongo-keyfile
    ports:
      - 27032:27019
    volumes:
      - /etc/localtime:/etc/localtime
      - /disk3/mongo-config2:/data/configdb
      - ./mongo-keyfile:/data/mongo-keyfile

  mongo_config3:
    image: mongo:7.0.7
    container_name: mongo_config3
    command: mongod --configsvr --directoryperdb --replSet fates-mongo-config  --keyFile /data/mongo-keyfile
    ports:
      - 27033:27019
    volumes:
      - /etc/localtime:/etc/localtime
      - /disk3/mongo-config3:/data/configdb
      - ./mongo-keyfile:/data/mongo-keyfile

  mongo-mongos1:
    image: mongo:7.0.7
    container_name: mongo_mongos1
    command: /bin/sh -c 'mongos --configdb fates-mongo-config/mongo_config1:27019,mongo_config2:27019,mongo_config3:27019 --bind_ip 0.0.0.0 --port 27017 --keyFile /data/mongo-keyfile'
    ports:
      - 27017:27017
    volumes:
      - /etc/localtime:/etc/localtime
      - ./mongo-keyfile:/data/mongo-keyfile
    depends_on:
      - mongo_config1
      - mongo_config2
      - mongo_config3

  mongo-mongos2:
    image: mongo:7.0.7
    container_name: mongo_mongos2
    command: /bin/sh -c 'mongos --configdb fates-mongo-config/mongo_config1:27019,mongo_config2:27019,mongo_config3:27019 --bind_ip 0.0.0.0 --port 27017 --keyFile /data/mongo-keyfile'
    ports:
      - 27018:27017
    volumes:
      - /etc/localtime:/etc/localtime
      - ./mongo-keyfile:/data/mongo-keyfile
    depends_on:
      - mongo_config1
      - mongo_config2
      - mongo_config3

networks:
  dev_network:
    external: true
    name: dev_network
```

先不执行启动，先创建网络 / 证书，命令如下

```bash
#创建网络
docker network create --driver bridge dev_network

#创建key：文件放入yml文件目录
openssl rand -base64 745 > mongo-keyfile
chmod 400 ./mongo-keyfile
chown 999:999 ./mongo-keyfile
```

```bash
# 初始化 mongodb config 节点
docker-compose exec mongo_config1 bash -c "echo 'rs.initiate({_id: \"fates-mongo-config\",configsvr: true, members: [{ _id : 0, host : \"mongo_config1:27019\" },{ _id : 1, host : \"mongo_config2:27019\" }, { _id : 2, host : \"mongo_config3:27019\" }]})' | mongosh --port 27019"

  
# 初始化 mongodb shard 节点
docker-compose exec mongo_shard1 bash -c "echo 'rs.initiate({_id: \"mongo_shard1\",members: [{ _id : 0, host : \"mongo_shard1:27018\" }]})' | mongosh --port 27018"

docker-compose exec mongo_shard2 bash -c "echo 'rs.initiate({_id: \"mongo_shard2\",members: [{ _id : 0, host : \"mongo_shard2:27018\" }]})' | mongosh --port 27018"


  
  
# 将 shard 节点添加到 mongos-1 中
docker-compose -f docker-compose.yml exec mongo-mongos1 bash -c "echo 'sh.addShard(\"mongo_shard1/mongo_shard1:27018\")' | mongosh"
docker-compose -f docker-compose.yml exec mongo-mongos1 bash -c "echo 'sh.addShard(\"mongo_shard2/mongo_shard2:27018\")' | mongosh"

  
# 将 shard 节点添加到 mongos-2 中
docker-compose -f docker-compose.yml exec mongo-mongos2 bash -c "echo 'sh.addShard(\"mongo_shard1/mongo_shard1:27018\")' | mongosh"
docker-compose -f docker-compose.yml exec mongo-mongos2 bash -c "echo 'sh.addShard(\"mongo_shard2/mongo_shard2:27018\")' | mongosh"
```