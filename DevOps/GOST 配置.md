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

# 启动客户端
# root @ company-public in /docker/gost [12:02:25] C:130
$ cat Dockerfile        
FROM scratch 
WORKDIR /root
# 从构建阶段拷贝构建好的应用到最终镜像
COPY ./gost-linux /root/gost-linux
# 设置容器启动时执行的命令
#ENTRYPOINT ["/root/gost-linux -L=:3128 -F=http2://8.217.172.184:443"]
CMD ["/root/gost-linux", "-L=:3128", "-F=http2://8.217.172.184:443"]
```