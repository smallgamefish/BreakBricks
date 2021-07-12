pipeline {
    agent any

    stages {
        stage('拉取代码') {
            steps {
               checkout([$class: 'GitSCM', branches: [[name: "*/${branch}"]], extensions: [], userRemoteConfigs: [[credentialsId: 'e8652a60-e738-456c-9117-6c8ce691acff', url: 'git@github.com:smallgamefish/BreakBricks.git']]])
            }
        }
        stage('停止docker容器') {
            steps {
                sh 'sudo docker-compose down'
            }
        }
        stage('docker容器启动') {
            steps {
                sh 'sudo docker-compose up -d'
            }
        }
    }
}