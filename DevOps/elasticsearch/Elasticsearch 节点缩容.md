> 游戏上线前段时间的日志量很大，如果不导量，会日益减少
> 磁盘都是不支持缩容的，只能通过节点轮替的操作，将数据重新分配到低存储节点



> Tips：扩容相对容易，CPU / 内存不够可以扩节点数量，磁盘不够扩磁盘，磁盘扩容方式：
```bash
# 数据节点磁盘扩容，修改 6000Gi 为指定容量即可
kubectl -n middleware patch pvc elasticsearch-data-elasticsearch-es-data-0   -p '{"spec":{"resources":{"requests":{"storage":"6000Gi"}}}}' 
kubectl -n middleware patch pvc elasticsearch-data-elasticsearch-es-data-1   -p '{"spec":{"resources":{"requests":{"storage":"6000Gi"}}}}' 
kubectl -n middleware patch pvc elasticsearch-data-elasticsearch-es-data-2   -p '{"spec":{"resources":{"requests":{"storage":"6000Gi"}}}}' 
kubectl -n middleware patch pvc elasticsearch-data-elasticsearch-es-data-3   -p '{"spec":{"resources":{"requests":{"storage":"6000Gi"}}}}'
```


## 原配置

```bash
# 这里有个案例可以参考下 ： https://github.com/elastic/cloud-on-k8s/blob/2.12/deploy/eck-stack/examples/elasticsearch/hot-warm-cold.yaml
apiVersion: elasticsearch.k8s.elastic.co/v1
kind: Elasticsearch
metadata:
  name: elasticsearch
  namespace: middleware
spec:
  version: 8.13.4
  image: elastic/elasticsearch:8.13.4 
  http:
    service:
      spec:
        type: NodePort # 为了集群外部也可以使用 ( 内网 )，这里使用 NodePort
  # 删除集群时，PVC 不会被删除
  volumeClaimDeletePolicy: DeleteOnScaledownOnly
  # 更新策略 
  updateStrategy:
    changeBudget:
      maxSurge: 3
      maxUnavailable: 1
  # Pod 终断预算
  podDisruptionBudget:
    spec:
      minAvailable: 3
  nodeSets:
  # ===================================== 主节点配置 ======================================== #
  - name: master
    count: 3
    config:
      node.roles: ["master","remote_cluster_client"]
    # pod 模板，除了资源限制以外，还加了 init Container
    # init Container 原因详见 => https://www.elastic.co/guide/en/cloud-on-k8s/2.12/k8s-virtual-memory.html#k8s_using_an_init_container_to_set_virtual_memory
    podTemplate:
      metadata:
        labels:
          # 让 metricbeat 可以抓取
          scrape: es
      spec:
        imagePullSecrets:
        - name: harbor-auth
        affinity:
          nodeAffinity:
            requiredDuringSchedulingIgnoredDuringExecution:  # 硬策略
              nodeSelectorTerms:
              - matchExpressions:
                - key: module
                  operator: In
                  values:
                  - es
        initContainers:
        - name: sysctl
          securityContext:
            privileged: true
            runAsUser: 0
          command: ['sh', '-c', 'sysctl -w vm.max_map_count=262144']
        containers:
        - name: elasticsearch
          resources:
            limits:
              memory: 10Gi
              cpu: 5
    # 存储卷配置
    volumeClaimTemplates:
    - metadata:
        # 不要更改这个名字！ 改了会很麻烦
        name: elasticsearch-data  
      spec:
        accessModes:
        - ReadWriteOnce
        resources:
          requests:
            storage: 200Gi
        storageClassName: cloud-essd-sc 

  # ===================================== 数据节点配置 ======================================== #
  - name: data
    count: 4
    config:
      node.roles: ["data", "transform","remote_cluster_client"]
    # pod 模板，除了资源限制以外，还加了 init Container
    # init Container 原因详见 => https://www.elastic.co/guide/en/cloud-on-k8s/2.12/k8s-virtual-memory.html#k8s_using_an_init_container_to_set_virtual_memory
    podTemplate:
      metadata:
        labels:
          # 让 metricbeat 可以抓取
          scrape: es
      spec:
        imagePullSecrets:
        - name: harbor-auth
        affinity:
          nodeAffinity:
            requiredDuringSchedulingIgnoredDuringExecution:  # 硬策略
              nodeSelectorTerms:
              - matchExpressions:
                - key: module
                  operator: In
                  values:
                  - es
        initContainers:
        - name: sysctl
          securityContext:
            privileged: true
            runAsUser: 0
          command: ['sh', '-c', 'sysctl -w vm.max_map_count=262144']
        containers:
        - name: elasticsearch
          resources:
            limits:
              memory: 64Gi
              cpu: 32
    # 存储卷配置
    volumeClaimTemplates:
    - metadata:
        # 不要更改这个名字！ 
        name: elasticsearch-data  
      spec:
        accessModes:
        - ReadWriteOnce
        resources:
          requests:
            storage: 5000Gi
        storageClassName: cloud-essd-sc

  # ===================================== 协调节点配置 ======================================== #
  # https://www.elastic.co/guide/en/cloud-on-k8s/current/k8s-traffic-splitting.html
  # 关于协调节点，需要创建额外的 Service 做负载 / 或者 ingress 路由到 Coordinating 节点
  - name: coordinating
    count: 4
    config:
      node.roles: ["remote_cluster_client"]
    podTemplate:
      metadata:
        labels:
          scrape: es
      spec:
        imagePullSecrets:
        - name: harbor-auth
        affinity:
          nodeAffinity:
            requiredDuringSchedulingIgnoredDuringExecution:  # 硬策略
              nodeSelectorTerms:
              - matchExpressions:
                - key: module
                  operator: In
                  values:
                  - es
        initContainers:
        - name: sysctl
          securityContext:
            privileged: true
            runAsUser: 0
          command: ['sh', '-c', 'sysctl -w vm.max_map_count=262144']
        containers:
        - name: elasticsearch
          resources:
            limits:
              memory: 32Gi
              cpu: 16
    # 存储卷配置
    volumeClaimTemplates:
    - metadata:
        # 不要更改这个名字！
        name: elasticsearch-data
      spec:
        accessModes:
        - ReadWriteOnce
        resources:
          requests:
            storage: 200Gi
        storageClassName: cloud-essd-sc
```



## 新增配置 / 扩容新的节点组

> 上述配置不要动，下面配置，直接添加到行尾，此时要保证有 moudule=es 的节点有足够资源

```bash
  # ===================================== 数据节点配置（低配）======================================== #
  # name 一定要不同！（ 代表了一个节点组，和 PVC / PV 挂钩 ）
  - name: data-shrink
    count: 4
    config:
      node.roles: ["data", "transform","remote_cluster_client"]
    # pod 模板，除了资源限制以外，还加了 init Container
    # init Container 原因详见 => https://www.elastic.co/guide/en/cloud-on-k8s/2.12/k8s-virtual-memory.html#k8s_using_an_init_container_to_set_virtual_memory
    podTemplate:
      metadata:
        labels:
          # 让 metricbeat 可以抓取
          scrape: es
      spec:
        imagePullSecrets:
        - name: harbor-auth
        affinity:
          nodeAffinity:
            requiredDuringSchedulingIgnoredDuringExecution:  # 硬策略
              nodeSelectorTerms:
              - matchExpressions:
                - key: module
                  operator: In
                  values:
                  - es
        initContainers:
        - name: sysctl
          securityContext:
            privileged: true
            runAsUser: 0
          command: ['sh', '-c', 'sysctl -w vm.max_map_count=262144']
        containers:
        - name: elasticsearch
          resources:
            limits:
              memory: 32Gi
              cpu: 16
    # 存储卷配置
    volumeClaimTemplates:
    - metadata:
        # 不要更改这个名字！ 
        name: elasticsearch-data  
      spec:
        accessModes:
        - ReadWriteOnce
        resources:
          requests:
            storage: 3000Gi
        storageClassName: cloud-essd-sc
```



> 需要等待一段时间，数据会平衡下到新的节点

## 旧节点下线

> 直接在 yaml 文件中，删掉旧的高配节点，删除这段内容：

```bash
  # ===================================== 数据节点配置 ======================================== #
  - name: data
    count: 4
    config:
      node.roles: ["data", "transform","remote_cluster_client"]
    # pod 模板，除了资源限制以外，还加了 init Container
    # init Container 原因详见 => https://www.elastic.co/guide/en/cloud-on-k8s/2.12/k8s-virtual-memory.html#k8s_using_an_init_container_to_set_virtual_memory
    podTemplate:
      metadata:
        labels:
          # 让 metricbeat 可以抓取
          scrape: es
      spec:
        imagePullSecrets:
        - name: harbor-auth
        affinity:
          nodeAffinity:
            requiredDuringSchedulingIgnoredDuringExecution:  # 硬策略
              nodeSelectorTerms:
              - matchExpressions:
                - key: module
                  operator: In
                  values:
                  - es
        initContainers:
        - name: sysctl
          securityContext:
            privileged: true
            runAsUser: 0
          command: ['sh', '-c', 'sysctl -w vm.max_map_count=262144']
        containers:
        - name: elasticsearch
          resources:
            limits:
              memory: 64Gi
              cpu: 32
    # 存储卷配置
    volumeClaimTemplates:
    - metadata:
        # 不要更改这个名字！ 
        name: elasticsearch-data  
      spec:
        accessModes:
        - ReadWriteOnce
        resources:
          requests:
            storage: 5000Gi
        storageClassName: cloud-essd-sc
```



> Over ~