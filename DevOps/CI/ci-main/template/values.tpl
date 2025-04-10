namespace: ${NAMESPACE}

environment: ${ENV}

appname: ${APP}

replicaCount: ${NUM}

image: 
  repository: registry.cn-hangzhou.aliyuncs.com/gitops-demo/${APP}
  pullPolicy: IfNotPresent
  tag: ${TAG}-${COMMITID}
  prod_tag: ${TAG}
  pre_tag: ${TAG}

imagePullSecrets: ${ENV}-secret


env:
  profile: ${RUN_PROFILE} 

service:
  type: ClusterIP
  port: ${PORT}
  
resources:
  limits:
    memory: ${MAXMEM}
    cpu: ${MAXCPU}
  requests:
    memory: ${MINMEM}
    cpu: ${MINCPU}
  
lifecycle:
  command: 'curl -X POST http://127.0.0.1:${PORT}/actuator/serviceregistry?status=DOWN -H "Content-Type: application/vnd.spring-boot.actuator.v2+json;charset=UTF-8"'
