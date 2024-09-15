
官方文档
https://argocd.devops.gold/understand_the_basics/
阿里云遇到的 CiliumIdentity 问题
https://www.alibabacloud.com/help/zh/ack/gitops-faq
## ingress 访问

```bash
# argocd-apisixroute.yaml
apiVersion: apisix.apache.org/v2
kind: ApisixRoute
metadata:
  name: argocd
  namespace: argocd
spec:
  http:
    - name: root
      match:
        hosts:
          - argocd-server.xxx.com
        paths:
          - "/*"
      backends:
        - serviceName: argocd-server
          servicePort: 80
```


```bash
# argocd-apisixtls.yaml
apiVersion: apisix.apache.org/v2
kind: ApisixTls
metadata:
  name: sample-tls
spec:
  hosts:
  - argocd-server.xxx.com
  # 记得先建立好 tls secret 
  secret:
    name: xxx-tls
    namespace: octopus
```

![](assets/ArgoCD%20文档/ArgoCD%20文档_image_1.png)


没啥需要注意的
 
![](assets/ArgoCD%20文档/ArgoCD%20文档_image_2.png)

## 通知配置

在 argocd-notifications-cm 中配置，卡片 Json 太长了，此处不写了，请见当前目录下的 argocd-install.yaml

我使用了飞书通知，通知如图：


![](assets/ArgoCD%20文档/ArgoCD%20文档_image_3.png)


发布流水线脚本如下：

```groovy
pipeline {
    agent  {
        label "tc-jenkins"
    }
    parameters {
        extendedChoice(
            description: '需要发布的服务',
            multiSelectDelimiter: ',',
            name: 'SERVICES',
            quoteValue: false,
            saveJSONParameterToFile: false,
            type: 'PT_MULTI_SELECT',
            value: 'ug-server,ug-user-center,ug-quartz,ug-quartz-client,ug-manager,ug-manager-client',
            visibleItemCount: 8
        )
         string name: 'TAG', trim: true
    }

    stages {
        stage('Build message') {
            steps {
                script {
                    //Jenkins页面显示构建信息
                    buildName ("#${env.BUILD_NUMBER}-ugsdk-os-pre 发布版本: ${TAG}")
                    wrap([$class: 'BuildUser']) {
                        def user = env.BUILD_USER
                        currentBuild.description = "构建人：${user}\n发布服务: ${SERVICES}"
                    }
                }
            }
        }
        stage('Update Serices Tag / Upload static resources') {
            steps {
                checkout scmGit(branches: [[name: '*/master']], extensions: [lfs(), [$class: 'RelativeTargetDirectory', relativeTargetDir: 'infra']], userRemoteConfigs: [[credentialsId: 'node-root', url: 'ssh://git@bitbucket.x.gameale.com:7999/rov/rov_server_infra.git']])
                script {
                    def services = params.SERVICES.split(',')
                    services.each { service ->
			if (service == 'ug-user-center') {
			    sh """
                                docker create --name ug-usercenter-static-tmp registry-dev.gameale.com/tc/ug-user-center-os:${TAG}
                                docker cp ug-usercenter-static-tmp:/usr/share/nginx/html ./
                                docker rm ug-usercenter-static-tmp
                                mv html ${TAG}
                                ${WORKSPACE}/infra/ali_oss/ossutil64 -c ${WORKSPACE}/infra/ali_oss/alig_oss_rov_hmt_obt_accelerate --jobs 200 --loglevel info --force cp -r ./${TAG}/ oss://rov-hmt-obt-alig/static/usercenter/${TAG}
                                sed -i "s#tag: .*#tag: ${TAG}#g" ${WORKSPACE}/ug-ovs-pre/${service}/values.yaml
			    """
			} else {
			    sh """
                                sed -i "s#tag: .*#tag: ${TAG}#g" ${WORKSPACE}/ug-ovs-pre/${service}/values.yaml
			    """
			}
                    }
                }
            }
        }
        stage('Git Push') {
            steps {
                script {
                    wrap([$class: 'BuildUser']) {
                        def user = env.BUILD_USER
                        def user_id = env.BUILD_USER_ID
                        sh """
                        git config --global user.email "${user_id}@gameale.com"
                        git config --global user.name "${user}"
                        """
                    }
                }
                withCredentials([sshUserPrivateKey(credentialsId: 'node-root', keyFileVariable: 'IDENTITY')]) {
            sh(script: """
                git config core.sshCommand 'ssh -o StrictHostKeyChecking=no -i $IDENTITY'
                git checkout master
                git pull origin master
                git add ${WORKSPACE}/ug-ovs-pre/
                git commit -m "Update ${params.SERVICES} ${TAG}" || true
                git push origin master
            """, label: "Push changes to repository")
                }
           }
       }
   }
    // 后置操作
    post {
        always {
            echo "Post Cleanup: Deleting workspace."
            cleanWs()
            }
    }
}
```

