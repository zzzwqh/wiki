
>  AdvanceCronJob 生产环境使用，打包一些埋点 / 业务日志上传 OSS
```bash
# cat acj.yaml
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
            affinity:
              nodeAffinity:
                requiredDuringSchedulingIgnoredDuringExecution:  # 硬策略
                  nodeSelectorTerms:
                  - matchExpressions:
                    - key: module
                      operator: In
                      values:
                      - game
            hostNetwork: true
            imagePullSecrets:
              - name: harbor-auth
            containers:
              - name: nodes-cronjob
                image: xxxx.xxxx.com/xxx/nodes-cronjob:1.2
                env:
                  - name: ARCHIVED_SERVICE
                    value: "true"  # 设置为 "true" 时执行打包服务日志
                  - name: ARCHIVED_ANALYTIC
                    value: "true"  # 设置为 "true" 时执行打包埋点日志
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

 执行查看 ACJ 相关资源 ： kubectl get acj,bcj,po,secret -n cronjob -owide
 
![](assets/openkruise%20CRD%20实践/openkruise%20CRD%20实践_image_1.png)

