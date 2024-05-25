## 1. 安装 CRD / Operator

```bash
kubectl create -f https://download.elastic.co/downloads/eck/2.12.0/crds.yaml
kubectl create -f https://download.elastic.co/downloads/eck/2.12.0/operator.yaml
```

-----

## 2. ElasticSearch 集群搭建

```yaml
# 这里有个案例可以参考下 ： https://github.com/elastic/cloud-on-k8s/blob/2.10/deploy/eck-stack/examples/elasticsearch/hot-warm-cold.yaml
apiVersion: elasticsearch.k8s.elastic.co/v1
kind: Elasticsearch
metadata:
  name: roc
  namespace: roc
spec:
  # 指定 elasticsearch 镜像和版本
  version: 8.13.4
  image: elastic/elasticsearch:8.13.4 
  # 删除集群时，PVC 不会被删除
  volumeClaimDeletePolicy: DeleteOnScaledownOnly
  nodeSets:
  # ==================== 主节点配置 ==================== #
  - name: master
    count: 3
    # 指定节点角色，即一共 3 台 master 节点
    config:
      node.roles: ["master"]
    # ES 节点 pod 模板
    podTemplate:
      spec:
        # init Container 原因详见 => https://www.elastic.co/guide/en/cloud-on-k8s/2.10/k8s-virtual-memory.html#k8s_using_an_init_container_to_set_virtual_memory
        initContainers:
        - name: sysctl
          securityContext:
            privileged: true
            runAsUser: 0
          command: ['sh', '-c', 'sysctl -w vm.max_map_count=262144']
        containers:
        - name: elasticsearch
          # 限制节点资源，Operator 会根据配置的 resource limit 自动配置 JVM 参数
          resources:
            limits:
              memory: 4Gi
              cpu: 2
    # 存储卷配置
    volumeClaimTemplates:
    - metadata:
        # 不要更改这个名字！ 改了会很麻烦
        name: elasticsearch-data  
      spec:
        accessModes:
        - ReadWriteOnce
        # 阿里云的云盘 StorageClass 申请 PV 最少申请 20Gi，建议大于 20Gi
        resources:
          requests:
            storage: 50Gi
        storageClassName: nfs-client

  # ==================== 数据节点配置 ==================== #
  - name: data
    count: 2
    config:
      node.roles: ["data", "transform"]
    # pod 模板，除了资源限制以外，还加了 init Container
    # init Container 原因详见 => https://www.elastic.co/guide/en/cloud-on-k8s/2.10/k8s-virtual-memory.html#k8s_using_an_init_container_to_set_virtual_memory
    podTemplate:
      spec:
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
              memory: 4Gi
              cpu: 2
    # 存储卷配置
    volumeClaimTemplates:
    - metadata:
        # 不要更改这个名字！ 
        name: elasticsearch-data  
      spec:
        accessModes:
        - ReadWriteOnce
        # 阿里云的云盘 StorageClass 申请 PV 最少申请 20Gi，建议大于 20Gi
        resources:
          requests:
            storage: 50Gi
        storageClassName: nfs-client
```


## 3. 安装 Kibana

```yaml
apiVersion: kibana.k8s.elastic.co/v1
kind: Kibana
metadata:
  name: roc
  namespace: roc
spec:
  # 禁用 Kibana TLS 
  http:
    tls:
      selfSignedCertificate:
        disabled: true
  version: 8.13.4
  count: 1
  elasticsearchRef:
    name: roc
    namespace: roc
```


## 4. 配置 Kibana Ingress

```yaml
# APISIX Ingress 配置案例 
apiVersion: apisix.apache.org/v2
kind: ApisixRoute
metadata:
  name: kibana
  namespace: roc
spec:
  http:
    - name: root
      match:
        hosts:
          - roc-dev-kibana.xxxx.com
        paths:
          - "/*"
      backends:
        - serviceName: roc-kb-http
          servicePort: 5601
```

## 5. 安装 Beat

```yaml

```