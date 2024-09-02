@Library("CustomLibrary@main") _
pipeline {
  agent {
    kubernetes {
      cloud 'kubernetes'
      label "jnlp"
      yamlFile "ci/pod-jnlp-mvn.yaml"
   }
}
  options {
    // 丢弃旧构建 
    buildDiscarder logRotator(artifactDaysToKeepStr: '5', artifactNumToKeepStr: '50', daysToKeepStr: '3', numToKeepStr: '30')
    // 使用 timestamper 插件，启用后，执行 Job 的日志中显示时间戳 
    timestamps()
    // 构建超时时间，超过这个时间会失败 
    timeout(time: 3000, unit: 'SECONDS')
    // 跳过自动检出代码，这样可以节省时间，在 pipeline 中自己控制拉取代码的时刻，或者使用轻量检出，只检出有变动的代码文件 
    skipDefaultCheckout true
    // 使用 AnsiColor 插件，可以做到输出代码有颜色
    ansiColor('xterm')
    // 失败重试次数，如果某个 stage/steps 下执行失败，整个流水线会从【全部】重试执行，而不只是重试失败的 stage
    // retry(3)
}
 parameters {
    // 使用了 git Parameter 插件，可以自动获取远程代码分支
    // 使用了 pipeline script from SCM 配置，每一个业务代码的构建，都会使用 CI 代码仓库的 JenkinsFile，获取 CI 仓库的分支
    // gitParameter branchFilter: 'origin/(.*)', defaultValue: 'main', name: 'CI_BRANCH', type: 'PT_BRANCH'
    // 获取业务仓库的分支
    // gitParameter branch: '', branchFilter: 'origin/(.*)', defaultValue: 'main', description: '业务仓库分支', name: 'BRANCH', quickFilterEnabled: false, selectedValue: 'NONE', sortMode: 'NONE', tagFilter: '*', type: 'GitParameterDefinition', useRepository: 'https://gitlab.itsky.tech/demo/prometheus.git'
    // 布尔类型参数
booleanParam (
      defaultValue: true, // 这里设置默认值为 true
      description: 'Enable startup probe', // 参数描述
      name: 'startupProbe' // 参数名称
)
}


 
  


  environment { // 环境变量配置，【环境变量】可以从【参数变量】中获取赋值
    APP = "prometheus-java-demo"
    STARTUPPROBE = "${startupProbe}"
    REPOSITORY= "registry.cn-zhangjiakou.aliyuncs.com/wqhns"

}
  stages {  // stages：开始执行构建动作
      stage('Pull Code') {  // stage：拉取代码
          steps { // steps：具体执行的一些步骤
              echo "\033[43m Pull Code ...... \033[0m"
              // 使用了 checkout 片段检出代码（ Snippet Generator 可以生成该代码 ）
              checkout scmGit(branches: [[name: '*/${BRANCH}']], extensions: [], userRemoteConfigs: [[credentialsId: '32e9b289-b93d-4ac5-b789-d07805a4c7fa', url: 'https://gitlab.itsky.tech/demo/prometheus.git']])
          }
      }
      stage('Build Message') {  // stage：构建信息
          steps { // steps：具体执行的一些步骤
              echo "\033[43m Build Message ...... \033[0m"
                // 使用了【 Build Name and Description Setter 】插件，启用该插件后，可以自定义构建记录的名字，如下是根据我们的【环境变量/参数命名】重命名该次构建
              buildName ("${env.BUILD_NUMBER}-${JOB_NAME}-${params.BRANCH}-${ENV}")
              script {  //  script 语句块内，可以使用 groovy 代码，而不是拘束于声明式 pipeline 脚本，使得脚本更灵活
                    //  使用了【 build user vars plugin 】插件，启用该插件后，会给出 jenkins 构建时执行的用户及其用户 ID 等信息
                  wrap([$class: 'BuildUser']) {
                    def user = env.BUILD_USER // 从 env 中获取 User 信息，但我尝试在打印 sh "printenv" 时，不会有 BUILD_USER
                    println("${user}")
                    currentBuild.description = "构建人：${user}" // 即使 sh "printenv" 时不会有 BUILD_USER ，Description 中也是可以生效的，我使用 admin 构建时，显示的是构建人: admin
                    log.CustomMessage("使用了 Build user vars Plugin 生成当前构建的 Descripiton ...,当前函数是调用 CustomLibrary 共享库 vars/log.groovy 下的 CustomMessage 方法...")
                  }
                }
          }
      }
      stage('Maven Build') { // Maven 构建
          steps { // steps: 具体执行的一些步骤
              echo "\033[43m Maven Building ...... \033[0m"
              // 所有的容器的执行构建的工作目录是 => $workspace-volume/workspace/$Jobname/ ，我们可以去观察 pod template，该工作目录挂到了 emptydir 类型的 volume【 name:workspace-volume 】，无论切换到哪个容器都会如此
              // pod 无需主动声明， jenkins 会自动在我们提供的模板中注入这个 emptydir 类型的 volume，并且每个 container 的 Workspace 都是这个目录
              container("maven") {  // 选择执行的容器，选择 pod 模板中的 maven 容器，工作目录会自动切换到 $Jenkins_home/workspace/$Jobname/ 下
                 sh 'mvn clean  package  -Dmaven.test.skip=true -U' // 执行 maven 构建命令
              }
          }
      }
      stage('Sonar Scanner') { // Sonar Scanner  
          steps { // steps: 具体执行的一些步骤
              echo "\033[43m Sonar Scanner ...... \033[0m"
              container("sonar-scanner") {   
                 sh "sonar-scanner -Dsonar.projectKey=devops-maven-service -Dsonar.projectName=devops-maven-service -Dsonar.projectVersion=1.1  -Dsonar.host.url=http://sonarqube:9000 -Dsonar.login=19248d71b53d8e9848dbc6b9117b7579de076498 -Dsonar.sources=src -Dsonar.sourceEncoding=UTF-8 -Dsonar.java.binaries=target/classes -Dsonar.java.test.binaries=target/test-classes -Dsonar.language=java"
               //  请注意，-Dsonar.java.binaries 必须指定，该参数需要与 -Dsonar.sources 参数一起使用，以便 SonarQube 能够正确地将 Java 二进制文件与源代码相关联。SonarQube 使用这些二进制文件来执行代码分析，以便评估代码质量、发现潜在的安全漏洞和缺陷等。
               }
          }
      }      
      stage("Docker Build"){
          steps{ // steps: 具体执行的一些步骤
              echo "\033[43m Generate Docker Image And Push ...... \033[0m"
                script { env.COMMITID = sh(script: "git rev-parse --short HEAD", returnStdout: true).trim() }
                container("docker") { 
                  withDockerRegistry(credentialsId: "aliyun-acr	", url: "https://registry.cn-zhangjiakou.aliyuncs.com/") {
                  sh """ 
                    ls -a ; pwd                 
                    """
                  }
                 
                
              }
          }  
      }
  }
}


// Tips: 这个版本的 Jenkins Pipeline 是我为了测试 Active Chioce 插件而使用的
// 1. 到页面上新增 Active Choices Parameter 参数 ENV ，Groovy Script 如下
// return ["test","main","dev"]


// 2. 到页面上新增 Active Choices Reactive Parameter 参数 Branch ，
// 这个参数的 Referenced parameters 值要配置填写 ENV,git_url 
// Groovy Script 如下
// def trim_git_url = git_url.replaceFirst("https://", "")
// // println trim_git_url
// // def ENV= "test"
// if (ENV.equals("test")){
// def gettags = ("git ls-remote -h https://root:wqh127.0.0.1@${trim_git_url} release*").execute()
// return gettags.text.readLines().collect { it.split()[1].replaceAll('refs/heads/', '') }.unique()
// }
// else{
// def gettags = ("git ls-remote -h https://root:wqh127.0.0.1@${trim_git_url} develop feature*").execute()
// return  gettags.text.readLines().collect { it.split()[1].replaceAll('refs/heads/', '') }.unique()
// } 

// 3. 上面的 Active Choices Reactive Parameter 参数 Branch 参考两个变量而动态变化，一个是 ENV，还有个是 git_url 
// 上述 Pipeline 中并没有 git_url ,所以要么上面写个，要么自己在 UI 里写一个