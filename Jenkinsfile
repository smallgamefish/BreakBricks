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
                sh 'sudo docker-compose -f run-docker-compose.yml down'
            }
        }
        stage('docker里面编译golang应用') {
            steps {
                sh 'sudo docker-compose -f build-docker-compose.yml up'
            }
        }
        stage('docker里面运行golang应用') {
            steps {
                sh 'sudo docker-compose -f run-docker-compose.yml up -d'
            }
        }
    }
}