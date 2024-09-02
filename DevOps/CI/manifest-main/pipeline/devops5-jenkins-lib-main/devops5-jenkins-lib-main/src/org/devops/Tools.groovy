package org.devops

def GetCommitID(){
    ID = sh returnStdout: true, script: "git rev-parse HEAD"
    return ID - "\n"
}

//下载代码
def GetCode(branchName, srcUrl){
    checkout([$class: 'GitSCM', 
            branches: [[name: branchName]], 
            extensions: [], 
            userRemoteConfigs: [[
                credentialsId: 'd7e4e500-e5c6-4673-ae4b-d43bf4ff5d19', 
                url: srcUrl]]])
}
