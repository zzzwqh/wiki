
无他，搞了个能 ssh 的 - -


```bash
FROM docker.gamealecdn.com/python:3.13-slim

WORKDIR /root

COPY app /root/

# 引入代理安装工具
RUN export https_proxy=http://10.30.113.142:7890 http_proxy=http://10.30.113.142:7890 all_proxy=socks5://http://10.30.113.142:7890 && \
    apt-get update && apt-get install -y tzdata vim net-tools curl openssh-server iproute2 && \
    echo 'root:root' | chpasswd && \
    echo 'PermitRootLogin yes' >> /etc/ssh/sshd_config && \
    echo 'PasswordAuthentication yes' >> /etc/ssh/sshd_config && \
    sed -i 's/#UsePAM yes/UsePAM no/' /etc/ssh/sshd_config && \
    mkdir /run/sshd && \
    chmod 755 /run/sshd


# 安装 python 依赖
RUN python -m pip install -i https://mirrors.aliyun.com/pypi/simple/ --trusted-host=mirrors.aliyun.com --upgrade pip && \
    pip install -r requirements.txt  -i https://mirrors.aliyun.com/pypi/simple/ --trusted-host=mirrors.aliyun.com && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo Asia/Shanghai > /etc/timezone


CMD ["bash","-c","/usr/sbin/sshd -D && python3 /app/main.py"]
#CMD ["python3", "/app/main.py"]
```
