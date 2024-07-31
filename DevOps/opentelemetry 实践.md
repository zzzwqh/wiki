
# 安装




```bash
helm repo add open-telemetry https://open-telemetry.github.io/opentelemetry-helm-charts
helm fetch open-telemetry/opentelemetry-operator --untar

# 先安装依赖项
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.13.2/cert-manager.yaml


# 我放在了 ~/otel/ 下面
helm upgrade --install opentelemetry-operator  ~/otel/opentelemetry-operator -f ~/otel/opentelemetry-operator/values.yaml --namespace kube-otel --create-namespace
```

