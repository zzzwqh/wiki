

[Argo CD](https://argo-cd.readthedocs.io/en/stable/#what-is-argo-cd) 是针对 Kubernetes 的声明式 GitOps 持续交付工具。

---

##### 快速入门安装

```bash
wget -O argocd.yaml https://raw.githubusercontent.com/argoproj/argo-cd/v2.8.4/manifests/install.yaml
sed -i 's|quay.io/argoproj/argocd:v2.8.4|registry.cn-hangzhou.aliyuncs.com/s-ops/argocd:v2.8.4|g' argocd.yaml
sed -i 's|ghcr.io/dexidp/dex:v2.37.0|registry.cn-hangzhou.aliyuncs.com/s-ops/dex:v2.37.0|g' argocd.yaml
sed -i 's|redis:7.0.11-alpine|registry.cn-hangzhou.aliyuncs.com/s-ops/redis:7.0.11-alpine|g' argocd.yaml
sed -i 's/imagePullPolicy: Always/imagePullPolicy: IfNotPresent/g' argocd.yaml
kubectl create ns argocd
kubectl apply -f argocd.yaml -n argocd
```

---

我们可以通过配置 Ingress 的方式来对外暴露服务，其他 Ingress 控制器的配置可以参考官方文档 https://argo-cd.readthedocs.io/en/stable/operator-manual/ingress/ 进行配置,本案例使用非http的方式。

----

需要在禁用 TLS 的情况下运行 APIServer,编辑 YAML中argocd-server 这个 Deployment 以将` --insecure` 标志添加到 argocd-server 命令，或者简单地在 `argocd-cmd-params-cm` ConfigMap中设置 `server.insecure: "true"` 即可:

![image-20250320162325204](https://wiki.itsky.tech/image/image-20250320162325204.png)

```bash
cat > http-argocd-ingress.yaml <<EOF
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: argocd
  namespace: argocd
spec:
  ingressClassName: nginx
  rules:
    - http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: argocd-server
                port:
                  name: http
      host: argocd.tbchip.com
EOF
```

当然我们也可以通过NodePort的方式来访问

```bash
kubectl patch svc argocd-server -n argocd -p '{"spec": {"type": "NodePort"}}'
```

默认情况下 admin 帐号的初始密码是自动生成的，会以明文的形式存储在 Argo CD 安装的命名空间中名为 argocd-initial-admin-secret 的 Secret 对象下的 password 字段下，我们可以用下面的命令来获

```bash
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d && echo
```

Argocd配置认证,我们的资源文件都在ops-repo仓库,我们这边演示http的方式认证,或者打开URL地址`http://argocd.tbchip.com/settings/repos?addRepo=true`

![image-20250314132722525](https://wiki.itsky.tech/image/image-20250314132722525.png)

发布实践

---

##### jenkins新建一个job

![image-20250326174722325](https://wiki.itsky.tech/image/image-20250326174722325.png)

##### 配置JOB

![image-20250326174941429](https://wiki.itsky.tech/image/image-20250326174941429.png)

##### 服务发布,先点击构建,其次点击取消,再次点击构建就可以根据环境构建,具体Jenkinsfile可以把代码仓库拉下来查看

![image-20250326175120892](https://wiki.itsky.tech/image/image-20250326175120892.png)

##### 访问

```bash
生产:
cat >prod-ingress.yaml <<EOF
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: xsk-mall-erp-prod
  namespace: prod
spec:
  ingressClassName: nginx
  rules:
    - http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: xsk-mall-erp
                port:
                  name: http
      host: ui-prod.tbchip.com
EOF

开发:
cat >dev-ingress.yaml <<EOF
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: xsk-mall-erp-dev
  namespace: prod
spec:
  ingressClassName: nginx
  rules:
    - http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: xsk-mall-erp
                port:
                  name: http
      host: ui-dev.tbchip.com
EOF
```

