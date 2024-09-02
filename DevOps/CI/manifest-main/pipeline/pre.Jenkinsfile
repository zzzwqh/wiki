//使用列表切片语法[0..-2]选择列表中除最后一个元素之外的所有元素，即去掉了最后一个"-"字符以及其后面的部分。最后，它使用join函数将剩余的字符串列表连接成一个新的字符串，并将其赋值给appName变量
def appname = env.JOB_NAME.split('-')[0..-2].join('-')

// 引用共享库
@Library("jenkins_shareLibrary") _

// 应用共享库中的方法
def tools = new org.devops.tools()

pipeline {
    agent {
    kubernetes {
        //Jenkins连接的k8s集群设置的别名
        cloud 'kubernetes'

        //jenkins动态slave的命名
        label "jnlp-slave-${UUID.randomUUID().toString().substring(0, 8)}"

        //设置为2 意味着当一个构建执行器（worker node）在空闲状态超过 2 分钟时，Jenkins 将会终止与该节点关联的执行器
        idleMinutes 2 

        //调用mvn的pod
        yamlFile "ci/pod-jnlp.yaml"
   }
}

environment {  
  //helm template 服务变量
  //APP = "${JOB_NAME.split("-")[0]}"
  APP = "${appname}"
  TAG ="${params.IMAGE_VERSION}"
  
  //镜像仓库地址
  REPOSITORY= "hundun-registry-registry.cn-shanghai.cr.aliyuncs.com/hundun_registry"

  //开启存活探针
  LIVE_VALUE = "true"
  //开启就绪探针
  READ_VALUE = "true"
  ENV ="${JOB_NAME.split("-")[-1]}"
  // 命名空间
  NAMESPACE = "hundun-${ENV}"
  
  // 业务线
  PROJECT = "${project}"
  
  // 服务yaml仓库
  CD_REPO = "${cd_repo}"
  
  //jvm
  RUN_PROFILE= "${jvm}"

  //RESOURCES
  MINCPU = "${env.Resources.split(",")[0]}"
  MAXCPU = "${env.Resources.split(",")[1]}"
  MINMEM = "${env.Resources.split(",")[2]}"
  MAXMEM = "${env.Resources.split(",")[3]}"
  //服务副本数
  NUM = "${env.Resources.split(",")[4]}"
}

   parameters {
    RESTList(
      name: 'IMAGE_VERSION',
      description: '',
      restEndpoint: "http://172.19.100.13:30794/repo/hundun_registry/${appname}/tags",
      credentialId: '',
      mimeType: 'APPLICATION_JSON',
      valueExpression: '$.Images[*].Tag',
      cacheTime: 10,    // optional
      defaultValue: '', // optional
      filter: '.*',     // optional
      valueOrder: 'NONE' // optional
    )
}
options {
    //保持构建15天 最大保持构建的30个 发布包保留15天
    buildDiscarder logRotator(artifactDaysToKeepStr: '15', artifactNumToKeepStr: '', daysToKeepStr: '15', numToKeepStr: '30')
    //时间模块
    timestamps()
    //超时时间
    timeout(time:20, unit:'MINUTES')
    //跳过默认设置的代码check out
    skipDefaultCheckout true
    //控制台输出的字符串变成你想要的颜色的显示
    ansiColor('xterm')
}
 
    stages {
        stage('Build message') {
            steps {
                script {
                    //Jenkins页面显示构建信息
                    buildName ("${APP}-${params.IMAGE_VERSION}-${ENV}")
                    wrap([$class: 'BuildUser']) {
                        def user = env.BUILD_USER
                        println("${user}")
                        currentBuild.description = "构建人：${user}"
                        tools.PrintMes("构建人: ${user}","green1")
                    }
                }
            }
        } 
        stage('pull  template ') {
            steps {
                script{
                echo '\033[34mstart-------------------------\033[0m \033[33mpull------------------------\033[0m \033[35mtemplate!\033[0m'
                //拉取helm模板
                checkout([$class: 'GitSCM', branches: [[name: "master"]], 
                extensions: [], userRemoteConfigs: [[credentialsId: 'gitlab-uat', 
                url: "http://172.19.32.226:8900/devops/ci.git"]]])
                echo "${APP}"
                    }
                }
            }
        stage ('生成yaml') {
            steps{
                 container('docker'){
                script{
                    echo '\033[34mgenerate-------------------------\033[0m \033[33mtemplate------------------------\033[0m \033[35myaml!\033[0m'
                    //渲染服务helm模板
                    dir ("pipeline") {
                        sh """
                       envsubst < ../template/${PROJECT}/values.tpl > ${PROJECT}/service/values.yaml &&\
                       helm template ${PROJECT}/service --output-dir=../workloads/${ENV} -f ${PROJECT}/service/values.yaml &&\
                       mv ../workloads/${ENV}/${PROJECT}/templates/api-service.yaml	../workloads/${ENV}/${PROJECT}/templates/${APP}-${ENV}-dep.yaml && ls &&\
                       cat ../workloads/${ENV}/${PROJECT}/templates/${APP}-${ENV}-dep.yaml
                        """
                        }
                    }
                }
            }
        }   
        stage('push yaml ') {
            //上传业务yaml到对应的业务线
            steps {	
                container('docker'){
                withCredentials([
                    sshUserPrivateKey(
                    credentialsId: 'gitlab-repo', 
                    keyFileVariable: 'identity')
                ]) {
             echo "\033[43m Push template yaml!  \033[0m"
            sh """
            pwd
			cd  workloads/
            pwd
            rm -fr dev/README.md test/README.md prod/README.md hpa/README.md
			git config --global user.name 'ci'
            git config --global user.email 'ci@hundunyun.com.cn'
			git init
            git remote add origin_new ${CD_REPO}
			git config core.sshCommand "ssh -o StrictHostKeyChecking=no -i $identity"
			git pull origin_new master
            [[ -f ${ENV}/${PROJECT}/templates/${APP}-${ENV}-dep.yaml ]] && mv ${ENV}/${PROJECT}/templates/${APP}-${ENV}-dep.yaml ${PROJECT}/${ENV}/${APP}-${ENV}-dep.yml
			[[ -f ${PROJECT}/${ENV}/${APP}-${ENV}-dep.yml ]] && git add ${PROJECT}/${ENV}/${APP}-${ENV}-dep.yml				
            git commit -m "${APP}-${params.IMAGE_VERSION}-TO-${ENV}"|| true
            git push -u origin_new HEAD:refs/heads/master --force
			"""      
               }     
			}  
          }
        }
    }
}
