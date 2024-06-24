
> 使用容器部署 cadvisor，会缺一些监控数据，所以部署在虚拟机上

```bash
 wget https://github.com/google/cadvisor/releases/download/v0.49.1/cadvisor-v0.49.1-linux-amd64
 mkdir /data/monitor
 mv cadvisor-v0.49.1-linux-amd64 /data/monitor/
```

> Systemd 管理文件

```bash
cat /lib/systemd/system/cadvisor.service
[Unit]
Description=cAdvisor Service
Documentation=https://github.com/google/cadvisor
After=network.target

[Service]
ExecStart=/data/monitor/cadvisor-v0.49.1-linux-amd64 -port 18080
Restart=always
RestartSec=5s
User=root
Group=root
Environment="HOME=/root"
WorkingDirectory=/data/monitor

[Install]
WantedBy=multi-user.target
```

