### GOST 服务端（HK Server）
```
wget https://github.com/ginuerzh/gost/releases/download/v2.11.5/gost-linux-amd64-2.11.5.gz 
unzip gost-linux-amd64-2.11.5.gz
chmod +x gost-linux-amd64-2.11.5

  211  mv ./gost-linux-amd64-2.11.5 gost
  212  gost -L=http2://:443
  213  ./gost -L=http2://:443
```