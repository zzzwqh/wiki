package org.devops

// 上传制品
def PushNexusArtifact(repoId, targetDir, pkgPath, pkgName){
    withCredentials([usernamePassword(credentialsId: '96a795bb-f206-475f-b0f4-a8ccab185a04', \
                                    passwordVariable: 'PASSWD', 
                                    usernameVariable: 'USERNAME')]) {
        sh """
            curl -X 'POST' \
              "http://192.168.1.200:8081/service/rest/v1/components?repository=${repoId}" \
              -H 'accept: application/json' \
              -H 'Content-Type: multipart/form-data' \
              -F "raw.directory=${targetDir}" \
              -F "raw.asset1=@${pkgPath}/${pkgName};type=application/java-archive" \
              -F "raw.asset1.filename=${pkgName}" \
              -u ${USERNAME}:${PASSWD}
        """
    }
}

//发布制品
def AnsibleDeploy(){
    //将主机写入清单文件
    sh "rm -fr hosts "
    for (host in "${env.deployHosts}".split(',')){
        sh " echo ${host} >> hosts"
    }


    // ansible 发布jar
    sh """
        # 主机连通性检测
        ansible "${env.deployHosts}" -m ping -i hosts 
        
        # 创建备份目录
        ansible "${env.deployHosts}" -m shell -a "mkdir -p ${env.targetDir}/${env.projectName}.bak || echo file is exists" 
        # 备份上次构建
        ansible "${env.deployHosts}" -m shell -a " mv ${env.targetDir}/${env.projectName}/* ${env.targetDir}/${env.projectName}.bak/ || echo file not exists"

        # 清理和创建发布目录
        ansible "${env.deployHosts}" -m shell -a "rm -fr ${env.targetDir}/${env.projectName}/* &&  mkdir -p ${env.targetDir}/${env.projectName} || echo file is exists" 
        # 复制app
        ansible "${env.deployHosts}" -m copy -a "src=${env.pkgName}  dest=${env.targetDir}/${env.projectName}/${env.pkgName}" 
    """
    // 发布脚本
    fileData = libraryResource 'scripts/service.sh'
    //println(fileData)
    writeFile file: 'service.sh', text: fileData
    //sh "ls -a ; cat service.sh "

    sh """
        # 修改变量
        sed -i 's#APPNAME=NULL#APPNAME=${env.projectName}#g' service.sh
        sed -i 's#VERSION=NULL#VERSION=${env.releaseVersion}#g' service.sh
        sed -i 's#PORT=NULL#PORT=${env.port}#g' service.sh

        # 复制脚本
        ansible "${env.deployHosts}" -m copy -a "src=service.sh  dest=${env.targetDir}/${env.projectName}/service.sh" 
        # 启动服务
        ansible "${env.deployHosts}" -m shell -a "cd ${env.targetDir}/${env.projectName} ;source /etc/profile && sh service.sh start" -u root

        # 检查服务 
        sleep 10
        ansible "${env.deployHosts}" -m shell -a "cd ${env.targetDir}/${env.projectName} ;source /etc/profile && sh service.sh  check" -u root
    """
}

//RollBack
def AnsibleRollBack(){

    sh """
        # 清理和创建发布目录
        ansible "${env.deployHosts}" -m shell -a "rm -fr ${env.targetDir}/${env.projectName}/* &&  mkdir -p ${env.targetDir}/${env.projectName} || echo file is exists" 
        
        # 将备份目录内容复制到发布目录
        ansible "${env.deployHosts}" -m shell -a " mv ${env.targetDir}/${env.projectName}.bak/* ${env.targetDir}/${env.projectName}/ || echo file not exists"
        
        # 停止服务
        ansible "${env.deployHosts}" -m shell -a "cd ${env.targetDir}/${env.projectName} ;source /etc/profile  && sh service.sh stop" -u root


        # 启动服务
        ansible "${env.deployHosts}" -m shell -a "cd ${env.targetDir}/${env.projectName} ;source /etc/profile  && sh service.sh  start" -u root

        # 检查服务 
        sleep 10
        ansible "${env.deployHosts}" -m shell -a "cd ${env.targetDir}/${env.projectName} ;source /etc/profile  && sh service.sh check" -u root

        """
}
