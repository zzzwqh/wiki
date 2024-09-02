// pipeline 
// 引用共享库
@Library("jenkins_shareLibrary") _

// 应用共享库中的方法
def tools = new org.devops.tools()

pipeline {
    agent {
    kubernetes {
        cloud 'kubernetes'
        label "jnlp-slave-${UUID.randomUUID().toString().substring(0, 8)}"
        yamlFile "ci/pod-jnlp-mvn.yaml"
   }
}

parameters {
    //代码分支
   gitParameter branchFilter: 'origin/(.*)', defaultValue: 'main', name: 'BRANCH', type: 'PT_BRANCH'
}

environment {  
  //helm template 服务变量
  APP = "${JOB_NAME}"
  //镜像仓库地址
  REPOSITORY= "xxxx/devops"
  //副本数
  NUM = "${replicas}"
  //开启存活探针
  LIVE_VALUE = "true"
  //开启就绪探针
  READ_VALUE = "true"
  // 命名空见
  NAMESPACE = "hundun-${ENV}"
  // 业务线
  PROJECT = "${project}"
  // yaml仓库
  CD_REPO = "${cd_repo}"
  //jvm
  RUN_PROFILE= "${jvm}"
  TAG ="${BUILD_ID}-${BRANCH}"
}

options {
    //保持构建15天 最大保持构建的30个 发布包保留15天
    buildDiscarder logRotator(artifactDaysToKeepStr: '15', artifactNumToKeepStr: '', daysToKeepStr: '15', numToKeepStr: '30')
    //时间模块
    timestamps()
    //超时时间
    timeout(time:15, unit:'MINUTES')
    //跳过默认设置的代码check out
    skipDefaultCheckout true
    //控制台输出的字符串变成你想要的颜色的显示
    ansiColor('xterm')
}  


    stages {
        stage('pull code') {
            steps {
                echo '\033[34mstart-------------------------\033[0m \033[33mpull------------------------\033[0m \033[35mcode!\033[0m'
                checkout([$class: 'GitSCM', branches: [[name: "${BRANCH}"]], 
                extensions: [], userRemoteConfigs: [[credentialsId: 'gitlab-uat', 
                url: "${git_url}"]]])
                
                echo "${PROJECT}"


            
        }
        stage('Build message') {
            steps {
                script {
                    
                    wrap([$class: 'BuildUser']) {
                        buildName ("${env.BUILD_NUMBER}-${JOB_NAME}-${params.BRANCH}-${ENV}")
                        def user = env.BUILD_USER
                        println("${user}")
                        currentBuild.description = "构建人：${user}"
                        tools.PrintMes("构建人: ${user}","green1")
                    }
                }
            }
        } 
        stage('mvn ') {
            steps {
                container('maven'){
                    //
                    script{
                        tools.PrintMes("编译打包","green1")
                        sh ' mvn clean  package  -Dmaven.test.skip=true -U'
                    }
                }                
            }
        }
        stage('docker images ') {
            steps {
                container('docker'){
                    script{
                        tools.PrintMes("build images","green1")
                        //获取git提交的commit_ID前7位
                        env.COMMITID = sh(script: "git rev-parse --short HEAD", returnStdout: true).trim()
                        withDockerRegistry(credentialsId: "aliyun-image-registry", url: "https://xxxr.aliyuncs.com/") {
                    sh """
cat >Dockerfile<<EOF
FROM registry.cn-hangzhou.aliyuncs.com/s-ops/alpine-jdk8:latest
WORKDIR /opt
COPY  ./target/*.jar /opt/
ENTRYPOINT ["sh","-c","exec java -jar /opt/*.jar"]
EOF
                    docker build -f Dockerfile -t ${REPOSITORY}/${APP}:${BUILD_ID}-${BRANCH}-${COMMITID} .
                    docker push ${REPOSITORY}/${APP}:${BUILD_ID}-${BRANCH}-${COMMITID}
                    docker rmi ${REPOSITORY}/${APP}:${BUILD_ID}-${BRANCH}-${COMMITID}
                    """
                        }
                    }
                }
            }
        }
        stage('pull  template ') {
            steps {
                script{
                echo '\033[34mstart-------------------------\033[0m \033[33mpull------------------------\033[0m \033[35mtemplate!\033[0m'
                checkout([$class: 'GitSCM', branches: [[name: "master"]], 
                extensions: [], userRemoteConfigs: [[credentialsId: 'gitlab-uat', 
                url: "http://172.19.32.226:8900/devops/ci.git"]]])
                    }
                }
            }
        stage ('生成yaml') {
            steps{
                 container('docker'){
                script{
                    echo '\033[34mgenerate-------------------------\033[0m \033[33mtemplate------------------------\033[0m \033[35myaml!\033[0m'
                    dir ("pipeline") {
                        sh """
                        envsubst < ../template/${PROJECT}/values.tpl > ${PROJECT}/service/values.yaml &&\
                       helm template ${PROJECT}/service --output-dir=../workloads/${ENV} -f ${PROJECT}/service/values.yaml &&\
                       mv ../workloads/${ENV}/${PROJECT}/templates/api-service.yaml ../workloads/${ENV}/${PROJECT}/templates/${APP}-${ENV}-dep.yaml && ls
                       cd  ../workloads/
                        """
                        }
                    }
                }
            }
        }   
        stage('push yaml ') {
            steps { 
                container('docker'){
                withCredentials([
                    sshUserPrivateKey(
                    credentialsId: 'gitlab-repo', 
                    keyFileVariable: 'identity')
                ]) {
            sh """
            pwd
            cd  workloads/
            pwd
            rm -fr dev/README.md test/README.md prod/README.md hpa/README.md
            git config --global user.name 'pumingang'
            git config --global user.email 'pumingang@hundunyun.com.cn'
            git init
            git remote add origin_new ${CD_REPO}
            git config core.sshCommand "ssh -o StrictHostKeyChecking=no -i $identity"
            git pull origin_new ${ENV}
            mv ${ENV}/${PROJECT}/templates/${APP}-${ENV}-dep.yaml ${PROJECT}/${ENV}/${APP}-${ENV}-dep.yml
            git add ${PROJECT}/${ENV}/${APP}-${ENV}-dep.yml                     
            git commit -m "${BUILD_NUMBER}-${APP}-${BRANCH}-TO-${ENV}"|| true
            git push -u origin_new HEAD:refs/heads/${ENV} --force
            """      
                }     
            }  
          }
        }
    }
}
