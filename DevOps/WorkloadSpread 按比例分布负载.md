
因为我们数据库 / NAS 都在可用区 B，所以在 Serverless ACK 中，想按权重分布业务 Pod ，原声 Kubernetes 可以做到均匀打散 Pod 到各个拓扑域，

 OpenKruise 中


```bash
~/ugsdk-devops/ug-ovs-prod/ug-server (master*) » cat templates/workloadspread.yaml                                                                                     wangqihan-020037@Gameale123
apiVersion: apps.kruise.io/v1alpha1
kind: WorkloadSpread
metadata:
  name: {{ .Chart.Name }}-ws
  namespace: {{ .Release.Namespace }}
spec:
  targetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: {{ .Chart.Name }}
  subsets:
  - name: subset-a
    maxReplicas: 20%
    requiredNodeSelectorTerm:
      matchExpressions:
      - key: topology.kubernetes.io/zone
        operator: In
        values:
        - ap-southeast-1a
  - name: subset-b
    maxReplicas: 80%
    requiredNodeSelectorTerm:
      matchExpressions:
      - key: topology.kubernetes.io/zone
        operator: In
        values:
        - ap-southeast-1b
```



效果如图：

![](assets/WorkloadSpread%20按比例分布负载/WorkloadSpread%20按比例分布负载_image_1.png)

