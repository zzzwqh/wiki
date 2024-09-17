
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

需求，临时设置某个路由不可访问

```bash
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
#    alb.ingress.kubernetes.io/cors-allow-origin: "*"
#    alb.ingress.kubernetes.io/ssl-redirect: "true"
#    alb.ingress.kubernetes.io/enable-cors: "true"
#    alb.ingress.kubernetes.io/cors-expose-headers: "*"
#    alb.ingress.kubernetes.io/cors-allow-methods: "GET,POST,PUT,HEAD"
#    alb.ingress.kubernetes.io/cors-allow-credentials: "true"
#    alb.ingress.kubernetes.io/cors-max-age: "600"
    alb.ingress.kubernetes.io/actions.response-403: |
      [{
          "type": "FixedResponse",
          "FixedResponseConfig": {
              "contentType": "text/plain",
              "httpCode": "403",
              "content": "403 forbiden"
          }
      }]
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
    alb.ingress.kubernetes.io/ssl-redirect: 'true'
  name: ror-pre.xxx.com
  namespace: ugsdk-pre
  labels:
    ingress-controller: alb

spec:
  ingressClassName: alb-ugsdk-public-pre-gv
  rules:
    - host: ror-pre.xxx.com
      http:
        paths:
          - backend:
              service:
                name: response-403
                port:
                  name: use-annotation
            path: /pre-register
            pathType: Exact
    - host: ror-pre.xxx.com
      http:
        paths:
          - backend:
              service:
                name: redirect
                port:
                  name: use-annotation
            path: /
            pathType: Exact
    - host: ror-pre.xxx.com
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
        - ror-pre.xxx.com
      secretName: secret-xxx
```