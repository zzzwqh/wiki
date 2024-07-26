
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
          - argocd-server.gameale.com
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
  - argocd-server.gameale.com
  # 记得先建立好 tls secret 
  secret:
    name: gameale-tls
    namespace: octopus
```
