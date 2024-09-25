需求：绑定控制台创建的 alb
https://www.alibabacloud.com/help/zh/ack/serverless-kubernetes/user-guide/use-an-alb-ingress-to-configure-certificates-for-an-https-listener-1#section-f32-v8m-rfs
```bash
apiVersion: alibabacloud.com/v1
kind: AlbConfig
metadata:
  name: alb-ugsdk-pre-public
spec:
  config:
    # 如果重用已经存在的 ALB，可以直接已存在的 ALB id
    id: alb-xxxxxxxx
    # 如果新创建 ALB 的话，可以打开下面的注释，并注释掉上一行 id
    #name: alb-project-env-xxx
    # 资源组
    #resourceGroupId: rg-xxxxxx
    #addressType: Internet
    # 绑定共享带宽
    #billingConfig:
    #  bandWidthPackageId: "cbwp-xxxxxx"
    #zoneMappings:
    #- vSwitchId: vsw-xxxxx
    #- vSwitchId: vsw-xxxxx
  listeners:
    - port: 80
      protocol: HTTP
    - port: 443
      protocol: HTTPS
      caEnabled: false
      certificates:
      - CertificateId: 123456
        IsDefault: true
```


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

需求，流量按比例分批打到后端，37开

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
    alb.ingress.kubernetes.io/actions.forward: |
       [{
           "type": "ForwardGroup",
           "ForwardConfig": {
             "ServerGroups" : [{
               "ServiceName": "ug-user-center-svc",
               "Weight": 30,
               "ServicePort": 80
             },
             {
               "ServiceName": "ug-user-center-tea-svc",
               "Weight": 70,
               "ServicePort": 80
             }]
           }
       }]
    alb.ingress.kubernetes.io/actions.redirect: |
      [{
          "type": "Redirect",
          "RedirectConfig": {
              "host": "${host}",
              "path": "/pre-register",
              "port": "443",
              "protocol": "https",
              "query": "${query}",
              "httpCode": "301"
          }
      }]
  name: ror.xxx.com
  namespace: ugsdk-pre
  labels:
    ingress-controller: alb

spec:
  ingressClassName: alb-ugsdk-public-pre-gv
  rules:
    - host: ror.xxx.com
      http:
        paths:
          - backend:
              service:
                name: redirect
                port:
                  name: use-annotation
            path: /reservation
            pathType: Exact
    - host: ror.xxx.com
      http:
        paths:
          - backend:
              service:
                name: redirect
                port:
                  name: use-annotation
            path: /
            pathType: Exact
    - host: ror.xxx.com
      http:
        paths:
          - backend:
              service:
                name: forward
                port:
                  name: use-annotation
            path: /
            pathType: Prefix
  tls:
    - hosts:
        - ror.xxx.com
      secretName: secret-gameale
```