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
  name: project
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
  name: project
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
    name: project
    namespace: middleware
```


## 4. 配置 Kibana Ingress

```yaml
# APISIX Ingress 配置案例 
apiVersion: apisix.apache.org/v2
kind: ApisixRoute
metadata:
  name: kibana
  namespace: middleware
spec:
  http:
    - name: root
      match:
        hosts:
          - project-dev-kibana.xxxx.com
        paths:
          - "/*"
      backends:
        - serviceName: project-kb-http
          servicePort: 5601
```

## 5. 安装 Beat

```yaml
apiVersion: beat.k8s.elastic.co/v1beta1
kind: Beat
metadata:
  name: project
  namespace: middleware
spec:
  type: filebeat
  version: 8.13.4
  elasticsearchRef:
    name: project
  config:
    # 输入插件配置
    filebeat.inputs:
    # 插件类型
    - type: log
      # 日志收集路径
      paths:
        - /data/logs/*game*.log
      # 设置为 false，则不止是通过 tail 获取新增日志，而是从日志文件的开头开始读取，收集到日志文件中的所有数据（ 具体收集哪些文件和 ignore_older 相关 ）
      tail_files: false
      # 自定义字段加入
      fields:
        log: game
      # 在 5 分钟内，仍没有新的日志数据写入情况下关闭文件，避免文件句柄资源的消耗
      close_inactive: 5m
      # 在 Filebeat 启动时决定是否忽略某些文件，文件修改时间超过 10 分钟的不会读取
      ignore_older: 10m
      # Filebeat 会在经过指定的时间后强制关闭文件，即使文件在这段时间内是活跃的。这有助于避免长时间运行的文件句柄潜在的泄露问题，重新打开文件以刷新读取状态或日志文件句柄
      close_timeout: 1h
      # symlinks 允许 Filebeat 跟踪符号链接，即软连接文件也可以抓取到日志内容
      symlinks: true
      # Filebeat 会将解析的 JSON 文档的字段放置在事件的根级别
      json.keys_under_root: true
      # Filebeat 会来自 JSON 文档的字段，覆盖现有的同名字段
      json.overwrite_keys: true
      # 是否将 key 展开存储，比如 {"a.b.c":"text"} 会展开成 {"a":{"b":{"c":"text"}}}
      json.expand_keys: true
      # 若解析错误，会有报错字段和日志输出到 doc 中
      json.add_error_key: true


    # 索引声明周期配置
    setup.ilm.enabled: false
    # 索引模板配置
    setup.template.name: "project_dev"
    setup.template.pattern: "project_dev*"
    setup.template.settings:
      # 主分片数量，官方建议，若一个索引 40Gi 数据量，则设置 1 分片，80Gi 数据量，设置 2 分片
      index.number_of_shards: 3
      # 副本分片未设置，默认为 1
      # 在设置副本数量 1 的情况下，如果主分片有 3 个，那么副本分片一共也是 3 个，一共 6 个
    
    # 输出插件的配置
    output.elasticsearch:
      # 正常情况下，用户是 project-es-beat-user，但是缺少部分权限，可以用最高权限用户 elastic
      username: elastic
      password: 'x'
      # 指定索引名字
      index: project_dev-%{+yyyy.MM.dd}

    # 数据处理器配置
    processors:
    # 删除某些不需要的字段
    - drop_fields:
        fields: ['message']
    # 重写字段名，例如读取到字段 service 会被映射为 service_id
    - rename:
        fields:
        - from: "service"
          to: "service_id"

    # 是否开启 filebeat debug 模式（观察部分日志）
    # logging.level: debug
    # logging.selectors: ["*"]


  # filebeat daemonset 配置
  daemonSet:
    podTemplate:
      spec:
        # 官网建议配置
        dnsPolicy: ClusterFirstWithHostNet
        hostNetwork: true
        # 为了 filebeat 获取更多权限，做日志收集
        securityContext:
          runAsUser: 0
        # 配置容器的宿主机映射目录，官网给的是 /var/lib/container 全收集
        # 当前根据具体需求配置映射
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



## 数据盘扩容【 阿里云 ESSD StorageClass 】

```bash
# 数据节点磁盘扩容，修改 6000Gi 为指定容量即可

kubectl -n middleware patch pvc elasticsearch-data-elasticsearch-es-data-0   -p '{"spec":{"resources":{"requests":{"storage":"6000Gi"}}}}' 
kubectl -n middleware patch pvc elasticsearch-data-elasticsearch-es-data-1   -p '{"spec":{"resources":{"requests":{"storage":"6000Gi"}}}}' 
kubectl -n middleware patch pvc elasticsearch-data-elasticsearch-es-data-2   -p '{"spec":{"resources":{"requests":{"storage":"6000Gi"}}}}' 
kubectl -n middleware patch pvc elasticsearch-data-elasticsearch-es-data-3   -p '{"spec":{"resources":{"requests":{"storage":"6000Gi"}}}}'
```

## 解码证书
    

```Bash
# 解码证书 
kubectl get secret -n middleware elasticsearch-es-http-certs-public -o jsonpath='{.data.tls\.crt}' | base64 --decode > tls.crt 
 
# 解码私钥 
kubectl get secret -n middleware elasticsearch-es-http-certs-public -o jsonpath='{.data.ca\.crt}' | base64 --decode > ca.crt
```

  

## 其他集群创建该证书

```Bash
kubectl create secret tls elasticsearch-es-http-certs-public \ 
  --cert=tls.crt \ 
  --key=tls.key \ 
  -n middleware   
```
