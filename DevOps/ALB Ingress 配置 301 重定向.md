
需求： 访问 / 直接跳转到 /reservation

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
    alb.ingress.kubernetes.io/actions.redirect: |
      [{
          "type": "Redirect",
          "RedirectConfig": {
              "host": "${host}",
              "path": "/reservation",
              "port": "443",
              "protocol": "https",
              "query": "${query}",
              "httpCode": "301"
          }
      }]
    alb.ingress.kubernetes.io/ssl-redirect: "true"
  labels:
    ingress-controller: alb
  name: xxx.xxx.com
  namespace: ugsdk
spec:
  ingressClassName: alb-ugsdk-public-gv
  rules:
  - host: xxx.xxx.com
    http:
      paths:
      - backend:
          service:
            name: redirect
            port:
              name: use-annotation
        path: /
        pathType: Exact
  - host: xxx.xxx.com
    http:
      paths:
      - backend:
          service:
            name: ug-user-center-svc
            port:
              number: 80
        path: /
        pathType: Prefix
  tls:
  - hosts:
    - xxx.xxx.com
    secretName: secret-xxx
```