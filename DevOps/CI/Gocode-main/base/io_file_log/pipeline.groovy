pipeline {

parameters {
    //代码分支
    gitParameter branchFilter: 'origin/(.*)', defaultValue: 'master', name: 'BRANCH', type: 'PT_BRANCH'
    //部署环境
    choice choices: ['dev', 'test','prod'], name: 'NAMESPACE'
    //code git地址
    string description: '代码 git地址', name: 'coderegistry', trim: true

}

options {
    //最大保留数
    buildDiscarder(logRotator(numToKeepStr: '15'))
    //时间模块
    timestamps()
    //超时时间
    timeout(time:1, unit:'HOURS')
}

    agent
       {label  'jnlp-slave'}

    stages {
        // build message
        stage('set buildDescription') {
            steps {
	//构建项目名 构建环境 代码分支
                script {
                        buildName ("${JOB_NAME}-${env.BUILD_NUMBER}")
                        buildDescription ("部署环境: ${params.NAMESPACE} \n 代码分支：${params.BRANCH}")
                    }
                }
            }
        stage('Source') {
            steps {
                // git clone code
                checkout([$class: 'GitSCM', branches: [[name: "${params.BRANCH}"]], extensions: [], userRemoteConfigs: [[credentialsId: 'gitlub-pass', url: "${params.coderegistry}"]]])
            }
        }
        stage('Build package') {
            steps {
                container('maven') {
                    sh ' mvn clean  package  -Dmaven.test.skip=true'
                    //打包跳过测试
                }
            }
        }
        stage('build > push  image') {
            steps {
                container('jnlp-slave') {
                    script {

                        //git commit的tag前七位
                        git_commit = sh(script: "git rev-parse --short HEAD", returnStdout: true).trim()

                        //第一个是hub地址,aliyun-docker-registry是 hub 的账号密码
                        docker.withRegistry('https://registry.cn-hangzhou.aliyuncs.com','aliyun-docker-registry') {

                        //build 镜像  tag 镜像
                        def customImage = docker.build("${params.NAMESPACE}/${JOB_NAME}:${env.BUILD_NUMBER}-${git_commit}")
                        //push镜像
                        customImage.push()

                        //删除宿主机的镜像
                        sh "docker rmi ${params.NAMESPACE}/${JOB_NAME}:${env.BUILD_NUMBER}-${git_commit}"
                        sh "docker rmi registry.cn-hangzhou.aliyuncs.com/${params.NAMESPACE}/${JOB_NAME}:${env.BUILD_NUMBER}-${git_commit}"
                        }
                    }
                }
            }
        }
    }
}
