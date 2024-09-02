// GO

// 使用split函数以斜杠为分隔符拆分字符串，并提取最后一个元素,忽略大小写
def APP =  env.JOB_NAME.split('/').last().toLowerCase()

// 定义全局变量
def errorMessage 

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
        yamlFile "ci/pod-jnlp-go.yaml"
   }
}   

//helm template 服务变量
environment {  
   
  //服务的名字
  APP = "${APP}"

  //服务端口
  PORT= "${PORT}"

  //镜像TAG
  TAG ="${BUILD_ID}-${BRANCH}"
  
  //镜像仓库namespace
  REPO="${REPO}"

  //镜像仓库地址
  REPOSITORY= "registry.cn-hangzhou.aliyuncs.com/${REPO}"
  
  // 服务命名空间
  NAMESPACE = "${ENV}"
  
  // 业务线
  PROJECT = "${PROJECT}"
  
  // 服务yaml仓库
  CD_REPO = "${CD_REPO}"

  //RESOURCES
  MINCPU = "${env.RESOURCES.split(",")[0]}"
  MAXCPU = "${env.RESOURCES.split(",")[1]}"
  MINMEM = "${env.RESOURCES.split(",")[2]}"
  MAXMEM = "${env.RESOURCES.split(",")[3]}"

  //服务副本数
  NUM = "${env.RESOURCES.split(",")[4]}"

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
        stage('pull code') {
            steps {
                echo '--------------------------拉取代码-----------------------'
                //拉取业务代码
                checkout([$class: 'GitSCM', branches: [[name: "${BRANCH}"]], 
                extensions: [], userRemoteConfigs: [[credentialsId: 'gitlab', 
                url: "${env.GIT_URL}"]]])
            }
        }
        stage('go build') {
                steps {
                    echo '--------------------------编译打包-----------------------'
                    container('golang'){
                        script{
                            try {
                            sh """
                            export GOPROXY=https://goproxy.cn
                            GOOS=linux GOARCH=amd64
                            go build  -o ${APP}
                            """
                        } catch (Exception e) {
                            errorMessage = 'go build失败,终止流水线'
                            error(errorMessage)
                            }
                        }                
                    }
                }
            }                            
        stage('docker images') {
            steps {
                echo '-------------------------打包镜像-----------------------'
                container('docker') {
                    script {
                        try{
                        // 获取 Git 提交的 commit ID，并将这些值存储在环境变量中以供后续使用。
                        env.COMMITID = sh(script: "git rev-parse --short HEAD", returnStdout: true).trim()    
                        // 判断代码根目录是否都dokcerfile,不存在则编写dockerfile
                        def dockerfile = 'Dockerfile'
                        if (!fileExists(dockerfile)) {
                    sh """
cat >${dockerfile}<<EOF
FROM dockerproxy.com/library/alpine:3.16.2 as serviceDeploy
WORKDIR /opt
COPY ${APP}  /opt/
ENTRYPOINT ./${APP}
EOF
                    """
                }
                // 推送 Docker 镜像到 Registry
                withDockerRegistry(credentialsId: "aliyun-image-registry", url: "https://registry.cn-hangzhou.aliyuncs.com") {
                        sh"""
                        docker build --build-arg VERSION="${BRANCH}" -f ${dockerfile} -t ${REPOSITORY}/${APP}:${BUILD_ID}-${params.BRANCH}-${COMMITID} .
                        docker push ${REPOSITORY}/${APP}:${BUILD_ID}-${params.BRANCH}-${COMMITID} 
                        docker rmi ${REPOSITORY}/${APP}:${BUILD_ID}-${params.BRANCH}-${COMMITID}
                        """
                    }        
                } catch (Exception e){
                    errorMessage = 'docker images失败,终止流水线,请联系运维'
                    error(errorMessage)
                }
            }
        }
    }
}     
          stage('pull template ') {
            steps {
                script{
                 echo '--------------------------拉取模板-----------------------'
                checkout([$class: 'GitSCM', branches: [[name: "main"]], 
                extensions: [], userRemoteConfigs: [[credentialsId: 'gitlab', 
                url: "https://gitlab.itsky.tech/devops/ci.git"]]])
                    }
                }
            }
        stage ('generate yaml') {
            steps{
                 container('docker'){
                script{
                     echo '--------------------------生成模板-----------------------'
                        try{    
                            sh """
                        envsubst < ./template/values.tpl > ./template/values.yaml && 
                        helm template  --debug  template/${PROJECT}/service --output-dir=example/${ENV} -f template/values.yaml &&
                        cat example/${ENV}/${PROJECT}/templates/api-service.yaml
                            """
                                } catch (Exception e) {
                                    errorMessage('生成yaml模板失败,终止流水线,请联系运维')
                                    error(errorMessage)
                            }
                        }
                    }
                }
            }
        stage('push yaml') {
            steps {
                echo "--------------------------push yaml-----------------------"
                container('docker') {
                    withCredentials([
                        sshUserPrivateKey(
                            credentialsId: 'gitlab-repo',
                            keyFileVariable: 'identity'
                        )
                    ]) {
                        script {
                            try {
                                lock('git-lock') {
                                    sh '' // 占位符步骤
                                    sh """
                                        cd example/
                                        git config --global user.name '${APP}'
                                        git config --global user.email '${APP}@demo.com'
                                        git init
                                        git remote add ${APP} ${CD_REPO}
                                        git config core.sshCommand "ssh -o StrictHostKeyChecking=no -i $identity"
                                        git pull ${APP} main
                                        mkdir -p ${PROJECT}/${ENV}/${APP}
                                        mv ${ENV}/${PROJECT}/templates/api-service.yaml ${PROJECT}/${ENV}/${APP}/${APP}-${ENV}-dep.yml
                                        git add "${PROJECT}/${ENV}/${APP}/${APP}-${ENV}-dep.yml"
                                        git commit -m "${BUILD_NUMBER}-${APP}-${BRANCH}-TO-${ENV}" || true
                                        git push -u ${APP} HEAD:refs/heads/main --force
                                    """
                                }
                            } catch (Exception e) {
                                // 处理异常，可以输出错误信息或执行其他操作
                                errorMessage = 'push yaml失败,终止流水线,请联系运维'
                                // 可以选择失败流程或继续执行其他操作
                                error(errorMessage)
                            }
                        }
                    }
                }
            }
        }
    }   
}
