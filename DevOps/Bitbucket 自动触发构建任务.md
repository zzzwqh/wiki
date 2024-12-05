
> Bitbucket 代码提交 TAG 后自动构建任务

```groovy
 pipeline {
    agent {
      label {
        label 'jenkins-node'
        // 任务失败重试次数
        retries 3
      }
    }
    options {
        //保持构建15天 最大保持构建的30个 ，时间保留15天
        buildDiscarder logRotator(artifactDaysToKeepStr: '15', artifactNumToKeepStr: '', daysToKeepStr: '15', numToKeepStr: '30')
        //时间模块
        timestamps()
        //超时时间
        timeout(time:15, unit:'MINUTES')
        //控制台输出的字符串变成你想要的颜色的显示
        ansiColor('xterm')

    }
    triggers {
        GenericTrigger(
         genericVariables: [
          [key: 'eventKey', value: '$.eventKey'],
          [key: 'changes_ref_type', value: '$.changes[?(@.ref.type != "")].ref.type'],
          [key: 'changes_ref_ids', value: '$.changes[?(@.ref.type == "TAG")].ref.id'],
          [key: 'changes_ref_display_ids', value: '$.changes[?(@.ref.type == "TAG")].ref.displayId']
         ],
         causeString: 'Triggered on $changes_ref_ids_0',
         token: 'AE610D86-1940-47CC-8942-7F88B8E54DC3',
         tokenCredentialId: '',
         regexpFilterText: '$changes_ref_type_0',
         regexpFilterExpression: '^TAG$'
        )
    }
    stages {
        // 生成构建信息
        stage('Build Message Generating') {
            steps {
                script {
                    currentBuild.displayName = "${changes_ref_display_ids_0} 版本镜像构建#${BUILD_NUMBER}"
                    currentBuild.description = "自动触发构建任务，代码版本 ${changes_ref_ids_0}"
                }
            }
        }
        // 拉取服务代码
       stage('Pulling Game Server Code') {
            steps {
                echo "\033[46;35m 拉取 RommoGame 游戏代码仓库 ~ \033[0m"
                echo "分支 ====> $changes_ref_ids_0"
                checkout scmGit(branches: [[name: '$changes_ref_ids_0']], extensions: [], userRemoteConfigs: [[credentialsId: '********', url: 'ssh://git@bitbucket.*******.com:7999/xxxx/server.git']])
                
            }
        }
        // 拉取构建镜像的 shell 脚本
        stage('Pulling Infra Code') {
            steps {
                echo  "\033[46;35m 拉取 devops 仓库的构建工具代码仓库 ~ \033[0m"
                checkout scmGit(branches: [[name: '*/test']], extensions: [[$class: 'RelativeTargetDirectory', relativeTargetDir: './infra']], userRemoteConfigs: [[credentialsId: '**********', url: 'ssh://git@bitbucket.*****.com:7999/xxxx/server_infra.git']])
            }
        }
       stage('Build Images') {
            steps {
                echo  "\033[46;35m 执行镜像构建 ~ \033[0m"
                sh """
                    echo "building images .... ${changes_ref_display_ids_0}"
                    chmod +x infra/docker/compile.sh
                    ./infra/docker/compile.sh
                    chmod +x infra/docker/build_image.sh
                    ./infra/docker/build_image.sh gate ${changes_ref_display_ids_0}
                    ./infra/docker/build_image.sh game ${changes_ref_display_ids_0}
                    ./infra/docker/build_image.sh dbProxy ${changes_ref_display_ids_0}
                    ./infra/docker/build_image.sh friend ${changes_ref_display_ids_0}
                    ./infra/docker/build_image.sh auction ${changes_ref_display_ids_0}
                    ./infra/docker/build_image.sh center ${changes_ref_display_ids_0}
                    ./infra/docker/build_image.sh chat ${changes_ref_display_ids_0}
                    ./infra/docker/build_image.sh mail ${changes_ref_display_ids_0}
                    ./infra/docker/build_image.sh global ${changes_ref_display_ids_0}
                """
            }
        }
    }
}
```


## POSTMAN 测试数据

POST 请求地址
https://roc-dev-jenkins.gastudio.cn/generic-webhook-trigger/invoke\?token\=xxxxxx-xxxxxx-xxxxxxx

```json
{
    "eventKey": "repo:refs_changed",
    "date": "2023-10-12T12:09:02+0800",
    "actor": {
        "name": "zzzwqh",
        "emailAddress": "zzzwqh@xxx.com",
        "id": 6,
        "displayName": "xxxxx",
        "active": true,
        "slug": "zzzwqh",
        "type": "NORMAL",
        "links": {
            "self": [
                {
                    "href": "http://bitbucket.xxx.com:7990/users/zzzwqh"
                }
            ]
        }
    },
    "repository": {
        "slug": "test_project",
        "id": 215,
        "name": "test_project",
        "hierarchyId": "9df3b79b5eee0c091738",
        "scmId": "git",
        "state": "AVAILABLE",
        "statusMessage": "Available",
        "forkable": true,
        "origin": {
            "slug": "test_project",
            "id": 196,
            "name": "test_project",
            "hierarchyId": "9df3b79b5eee0c091738",
            "scmId": "git",
            "state": "AVAILABLE",
            "statusMessage": "Available",
            "forkable": true,
            "project": {
                "key": "TEST",
                "id": 249,
                "name": "test_project",
                "public": false,
                "type": "NORMAL",
                "links": {
                    "self": [
                        {
                            "href": "http://bitbucket.xxx.com:7990/projects/TEST"
                        }
                    ]
                }
            },
            "public": false,
            "links": {
                "clone": [
                    {
                        "href": "ssh://git@bitbucket.xxxx.com:7999/test/test_server.git",
                        "name": "ssh"
                    },
                    {
                        "href": "http://bitbucket.xxxx.com:7990/scm/test/test_server.git",
                        "name": "http"
                    }
                ],
                "self": [
                    {
                        "href": "http://bitbucket.xxxx.com:7990/projects/TEST/repos/test_server/browse"
                    }
                ]
            }
        },
        "project": {
            "key": "TESTCN",
            "id": 263,
            "name": "test-cn",
            "public": false,
            "type": "NORMAL",
            "links": {
                "self": [
                    {
                        "href": "http://bitbucket.xxxx.com:7990/projects/TESTCN"
                    }
                ]
            }
        },
        "public": false,
        "links": {
            "clone": [
                {
                    "href": "http://bitbucket.xxxx.com:7990/scm/testcn/test_server.git",
                    "name": "http"
                },
                {
                    "href": "ssh://git@bitbucket.xxxx.com:7999/testcn/test_server.git",
                    "name": "ssh"
                }
            ],
            "self": [
                {
                    "href": "http://bitbucket.xxxx.com:7990/projects/TESTCN/repos/test_server/browse"
                }
            ]
        }
    },
    "changes": [
        {
            "ref": {
                "id": "refs/tags/1.0.0.241128.1",
                "displayId": "1.0.0.241128.1",
                "type": "TAG"
            },
            "refId": "refs/tags/1.0.0.241128.1",
            "fromHash": "0000000000000000000000000000000000000000",
            "toHash": "76dc7d1647969fee4d020411366c1888150d9a3c",
            "type": "ADD"
        }
    ]
}
```
![](assets/Bitbucket%20自动触发构建任务/Bitbucket%20自动触发构建任务_image_1.png)

