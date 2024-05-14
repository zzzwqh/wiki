### GOST 服务端（HK Server）
```bash
# 下载
wget https://github.com/ginuerzh/gost/releases/download/v2.11.5/gost-linux-amd64-2.11.5.gz 
unzip gost-linux-amd64-2.11.5.gz
chmod +x gost-linux-amd64-2.11.5
mv ./gost-linux-amd64-2.11.5 gost

# 启动服务端
./gost -L=http2://:443
```


### GOST 客户端（ 类似 Clash ）
```bash
# 下载
wget https://github.com/ginuerzh/gost/releases/download/v2.11.5/gost-linux-amd64-2.11.5.gz 
unzip gost-linux-amd64-2.11.5.gz
chmod +x gost-linux-amd64-2.11.5
mv ./gost-linux-amd64-2.11.5 gost

# 构建客户端镜像的 Dockerfile，我把客户端也放在 linux 上了，客户端用了下 docker
$ cat Dockerfile        
FROM scratch 
WORKDIR /root
# 从构建阶段拷贝构建好的应用到最终镜像
COPY ./gost /root/gost
# 设置容器启动时执行的命令 , ${SERVER_IP} 是你香港服务端的地址
CMD ["/root/gost-linux", "-L=:3128", "-F=http2://${SERVER_IP}:443"]



# 部署 GOST 客户端
$ cat docker-compose.yml
version: '3'
services:
  centos1:
    restart: always
    image: xxx.xxx.com/abc/gost:v1
    ports:
      - "3128:3128"   # 通过宿主机的 IP:Port ，访问 Google
    deploy:
      resources:
        limits:
          cpus: '1'
          memory: 1.5G
```

### 测试命令

```bash
# 测试 Google 通不通的命令，${ClIENT_IP} 是你部署客户端（Docker）的机器内网地址
 curl  -xhttp://${CLIENT_IP}:3128 https://www.google.com -I

```