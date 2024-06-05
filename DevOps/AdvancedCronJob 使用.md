
> 业务日志放在宿主机 hostpath 上，需要将 hostpath 路径的日志打包并上传到 OSS，Kubernetes 原生的 Cronjob 没发以 DaemonSet 的形式做部署，使用了 Kruise 的 AdvancedCronJob + BroadcastJob


## 构建镜像

定时任务需要一个镜像，可以采用如下 Dockerfile
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

其中脚本如下：



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