
因为我们数据库 / NAS 都在可用区 B（ 跨可用区会有 0.5-0.8 ms 的延时 ），为了更好的读写性能，所以在 Serverless ACK 中，想按权重分布业务 Pod ，原生 Kubernetes 可以做到均匀打散 Pod 到各个拓扑域，但是无法按权重比例做分布

OpenKruise 中提供了一种 CRD，可以作用于 Deployment ，将 Deployment Pod 按照比例分布在制定的不同区域，叫做 WorkloadSpread，我们有需求如下：
- 业务 Deployment 中的 80% 的 Pod 会分布在可用区 B， 20% 的 Pod 会分布在可用区 A
- 即使触发 HPA 弹性扩容，也遵守该
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

>更多使用细节可以参考： https://openkruise.io/zh/blog/workloadspread