貌似没啥写的
## 

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

