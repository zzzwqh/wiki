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
  namespace: middleware
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
  namespace: middleware
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
apiVersion: beat.k8s.elastic.co/v1beta1
kind: Beat
metadata:
  name: roc
  namespace: middleware
spec:
  type: filebeat
  version: 8.13.4
  elasticsearchRef:
    name: roc
  config:
    filebeat.inputs:
    - type: log
      paths:
          - /data/logs/*game*.log
      tail_files: false
      fields:
        logfile_type: game
      close_inactive: 5m
      ignore_older: 10m
      close_timeout: 1h
      symlinks: true
      json.keys_under_root: true
      json.overwrite_keys: true


    # 索引声明周期配置
    setup.ilm.enabled: false
    # 索引模板配置
    setup.template.name: "roc_dev"
    setup.template.pattern: "roc_dev*"
    setup.template.settings:
      index.number_of_shards: 2

    output.elasticsearch:
      # 正常情况下，用户是 xxx-es-beat-user，但是缺少部分权限，可以用最高权限用户
      username: elastic
      password: 'x'
      index: roc_dev-%{+yyyy.MM.dd}

    processors:
     - decode_json_fields:
         fields: ["log"]
         process_array: false
         max_depth: 1
         target: ""
         overwrite_keys: true
         add_error_key: true
     - rename:
         fields:
           - from: "service"
             to: "service_name"

     - drop_fields:
         fields: ['log']


  daemonSet:
    podTemplate:
      spec:
        # 为了 filebeat 获取更多权限，做日志收集
        dnsPolicy: ClusterFirstWithHostNet
        hostNetwork: true
        securityContext:
          runAsUser: 0
        containers:
        - name: filebeat
          volumeMounts:
          - name: logs
            mountPath: /data/logs
        volumes:
        - name: logs
          hostPath:
            path: /data/logs
```