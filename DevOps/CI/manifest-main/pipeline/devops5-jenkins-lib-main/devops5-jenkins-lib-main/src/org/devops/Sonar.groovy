package org.devops

//代码扫描
def CodeScan(){
    withSonarQubeEnv(credentialsId: 'f775b64a-20da-4cb8-9e6e-12ea13777818') {
        withCredentials([string(credentialsId: '303ac14d-60da-4b19-9d0b-c8f7a43a79a5', 
                                variable: 'GITLABTOKEN')]) {
            sh """sonar-scanner \
                -Dsonar.login=${SONAR_AUTH_TOKEN} \
                -Dsonar.projectVersion=${env.branchName} \
                -Dsonar.branch.name=${env.branchName} \
                -Dsonar.gitlab.user_token=${GITLABTOKEN} \
                -Dsonar.gitlab.ref_name=${env.branchName} \
                -Dsonar.gitlab.commit_sha=${env.commitID}
            """
        }                    
    }
}
