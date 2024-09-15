
官方文档
https://argocd.devops.gold/understand_the_basics/
阿里云遇到的 CiliumIdentity 问题
https://www.alibabacloud.com/help/zh/ack/gitops-faq
## ingress 访问

```bash
# argocd-apisixroute.yaml
apiVersion: apisix.apache.org/v2
kind: ApisixRoute
metadata:
  name: argocd
  namespace: argocd
spec:
  http:
    - name: root
      match:
        hosts:
          - argocd-server.xxx.com
        paths:
          - "/*"
      backends:
        - serviceName: argocd-server
          servicePort: 80
```


```bash
# argocd-apisixtls.yaml
apiVersion: apisix.apache.org/v2
kind: ApisixTls
metadata:
  name: sample-tls
spec:
  hosts:
  - argocd-server.xxx.com
  # 记得先建立好 tls secret 
  secret:
    name: xxx-tls
    namespace: octopus
```

![](assets/ArgoCD%20文档/ArgoCD%20文档_image_1.png)


没啥需要注意的
 
![](assets/ArgoCD%20文档/ArgoCD%20文档_image_2.png)

## 通知配置

在 argocd-notifications-cm 中配置，卡片 Json 太长了，此处不写了，请见当前目录下的 argocd-install.yaml

我使用了飞书通知，通知如图：


![](assets/ArgoCD%20文档/ArgoCD%20文档_image_3.png)


发布 Pipeline 如下：