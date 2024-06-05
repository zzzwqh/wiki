
> 业务日志放在宿主机 hostpath 上，需要将 hostpath 路径的日志打包并上传到 OSS，Kubernetes 原生的 Cronjob 没发以 DaemonSet 的形式做部署，使用了 Kruise 的 AdvancedCronJob + BroadcastJob


## 构建镜像

> 定时任务需要一个镜像，可以采用如下 Dockerfile
```bash
# 使用 Alpine 作为基础镜像
FROM alpine:latest
# 更新软件包库并安装必要的工具, 配置时区, find 命令需要根据时间打包日志文件
RUN apk update && apk add --no-cache curl tar zip unzip tzdata && mkdir /app/scripts -p && \
    cp /usr/share/zoneinfo/Asia/Bangkok /etc/localtime && \
    echo "Asia/Bangkok" > /etc/timezone
# 设置工作目录
WORKDIR /app
# 拷贝相应工具到镜像中
COPY ossutil64 /app/
COPY scripts /app/scripts

# 默认命令
CMD ["/bin/sh","-c","/app/scripts/main.sh"]

```

> 其中 main.sh 脚本如下
```bash
# 一共写了三个脚本
$ ls scripts 
archive_and_upload_analytic_logs.sh  archive_and_upload_service_logs.sh  main.sh
$ cat scripts/main.sh
#!/bin/sh

# 获取当前时间
current_time=$(date +"%Y-%m-%d %H:%M:%S")

# 检查环境变量并执行相应脚本
if [ "$ARCHIVED_SERVICE" == "true" ]; then
  echo "$current_time Executing archive_and_upload_service_logs.sh"
  /app/scripts/archive_and_upload_service_logs.sh
else
  echo "$current_time Skipping archive_and_upload_service_logs.sh"
fi

if [ "$ARCHIVED_ANALYTIC" == "true" ]; then
  echo "$current_time Executing archive_and_upload_analytic_logs.sh"
  /app/scripts/archive_and_upload_analytic_logs.sh
else
  echo "$current_time Skipping archive_and_upload_analytic_logs.sh"
fi
```
> 打包 service 日志脚本

```bash
$ cat scripts/archive_and_upload_service_logs.sh
#!/bin/sh


# 日志目录
LOG_DIR="/data/logs"
# 存储桶信息
OSS_BUCKET="oss://xxxxxxx/archived-logs/service/"
# 阿里云配置文件，需要以 secret 形式挂载
CONFIG_PATH="/app/.ossutilconfig"
# 获取节点 IP 地址
HOST=$(hostname -i)
# 打包的文件名字
ARCHIVE_NAME="${HOST}-service-archived-logs-$(date +%Y%m%d_%H%M%S).tar.gz"

# 获取当前时间,这个时间给日志输出用
current_time=$(date +"%Y-%m-%d %H:%M:%S")

# 查找并打包 10 天以前的日志文件，传到 /data/temp-storage 目录空间下 , !!! 这里 mv 等同于删除操作 !!!
find $LOG_DIR -type f -name "*.log" -mtime +5 -exec mv {} /data/temp-storage/ \;
# 将传过来的文件打包
find /data/temp-storage/ -type f -name "*.log" | tar -czf /data/temp-storage/$ARCHIVE_NAME -T -

# 检查压缩文件是否成功创建 , 不成功则退出
if [ ! -f /data/temp-storage/$ARCHIVE_NAME ]; then
  echo "$current_time Failed to create archived-logs. Exiting."
  exit 1
fi

# 上传打包文件到OSS ，上传文件到 OSS 路径
/app/ossutil64 -c ${CONFIG_PATH}  cp /data/temp-storage/$ARCHIVE_NAME $OSS_BUCKET

# 检查上传是否成功
if [ $? -ne 0 ]; then
  echo "$current_time Failed to upload archive to Aliyun OSS. Exiting."
  exit 1
else
  echo "$current_time Successfully uploaded archive to Aliyun OSS."
fi


# 删除临时目录中的文件
find /data/temp-storage/ -type f -name "*.log" -exec rm -f {} \;

# 删除压缩包
rm -f /data/temp-storage/$ARCHIVE_NAME

# 打印结果
echo "$current_time Successfully deleted old logs on Kubernetes nodes."
```

> 打包 analytic 日志脚本

```bash
$ cat scripts/archive_and_upload_analytic_logs.sh
#!/bin/sh

# 日志目录
LOG_DIR="/data/Analytic"
# 存储桶信息
OSS_BUCKET="oss://xxxxxxx/archived-logs/analytic/"
# 阿里云配置文件，需要以 secret 形式挂载
CONFIG_PATH="/app/.ossutilconfig"
# 获取节点 IP 地址
HOST="$(hostname -i)"
# 打包的文件名字
ARCHIVE_NAME="${HOST}-analytic-archived-logs-$(date +%Y%m%d_%H%M%S).tar.gz"

# 获取当前时间,这个时间给日志输出用
current_time=$(date +"%Y-%m-%d %H:%M:%S")

# 查找并打包 5 天以前的日志文件 , 其中 /data/temp-storage 也是映射的 hostpath 目录 【 映射原因：避免容器出问题而丢失日志数据 】
find $LOG_DIR -type f -name "*.log.*" -mtime +5 | tar -czf /data/temp-storage/$ARCHIVE_NAME -T -

# 检查压缩文件是否成功创建 , 不成功则退出
if [ ! -f /data/temp-storage/$ARCHIVE_NAME ]; then
  echo "$current_time Failed to create archived-logs. Exiting."
  exit 1
fi

# 上传到OSS ，上传文件到 OSS 路径
/app/ossutil64 -c ${CONFIG_PATH}  cp /data/temp-storage/$ARCHIVE_NAME $OSS_BUCKET

# 检查上传是否成功
if [ $? -ne 0 ]; then
  echo "$current_time Failed to upload archive to Aliyun OSS. Exiting."
  exit 1
else
  echo "$current_time Successfully uploaded archive to Aliyun OSS."
fi

# !!! 这里删除操作 !!!
# 删除 5 天以前的日志文件和压缩文件 , Analytic 日志按天分割，除非 Cronjob 执行一天没执行完，否则不会漏掉日志没有打包而被删除, 但是 service 业务日志按小时分割, 采取了另外的逻辑
find $LOG_DIR -type f -name "*.log.*" -mtime +5 -exec rm -f {} \;

# 删除压缩包
rm -f /data/temp-storage/$ARCHIVE_NAME

# 打印结果
echo "$current_time Successfully deleted old logs on Kubernetes nodes."

```

## AdvancedCronJob 资源声明

```yaml
apiVersion: apps.kruise.io/v1alpha1
kind: AdvancedCronJob
metadata:
  name: nodes-cronjob
  namespace: cronjob
spec:
  schedule: "10 2 * * *"
  timeZone: "Asia/Bangkok"
  # 保留成功的历史的 Cronjob Pod 数量，可以用于观察日志
  successfulJobsHistoryLimit: 3
  template:
    broadcastJobTemplate:
      spec:
        template:
          spec:
            hostNetwork: true
            imagePullSecrets:
              - name: harbor-auth
            containers:
              - name: nodes-cronjob
                image: xxx.xxx.xxx/nodes-cronjob:1.0
                env:
                  - name: ARCHIVED_SERVICE
                    value: "true"  # 设置为 "true" 时执行 archived_and_upload_service_logs.sh
                  - name: ARCHIVED_ANALYTIC
                    value: "true"  # 设置为 "true" 时执行 archived_and_upload_analytic_logs.sh
                volumeMounts:
                  - name: services-logdir
                    mountPath: "/data/logs"
                  - name: analytic-logdir
                    mountPath: "/data/Analytic"
                  - name: temp-storage  # 为什么需要这个临时目录，保证 find 命令查找到的文件先移动到临时目录，再打包 【 避免两个 find 命令导致部分文件漏掉 】
                    mountPath: "/data/temp-storage"
                  - name: secret-volume
                    mountPath: "/app/.ossutilconfig"
                    subPath: .ossutilconfig
                    readOnly: true
            volumes:
              - name: services-logdir
                hostPath:
                  path: "/data/logs"
              - name: analytic-logdir
                hostPath:
                  path: "/data/Analytic"
              - name: temp-storage
                hostPath:
                  path: "/data/temp-storage"
              - name: secret-volume
                secret:
                  secretName: ossutilconfig-secret
            restartPolicy: Never
            failurePolicy:
              type: FailFast
              restartLimit: 1
```

## 其他资源创建

> 上面的  AdvancedCronJob 需要两个 Secret 创建，可以参考如下命令创建出来

```bash
# 镜像拉取 Secret
kubectl create secret docker-registry harbor-auth --docker-username=xxxxx --docker-password=xxxxx --docker-server=xxx.xxx.xxx/rov -n cronjob

# ossutil 配置文件 Secret，需要指定文件
kubectl create secret generic ossutilconfig-secret --from-file=.ossutilconfig -n cronjob

```