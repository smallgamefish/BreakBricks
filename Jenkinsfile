pipeline {
    agent any

    stages {
        stage('拉取代码') {
            steps {
               checkout([$class: 'GitSCM', branches: [[name: "*/${branch}"]], extensions: [], userRemoteConfigs: [[credentialsId: '05bad6fc-2115-4e94-a96d-e7f020f2a7fa', url: 'git@github.com:smallgamefish/BreakBricks.git']]])
            }
        }
        stage('停止应用容器') {
            steps {
                sh 'sudo docker-compose -f run-docker-compose.yml down'
            }
        }
        stage('停止编译容器') {
            steps {
                sh 'sudo docker-compose -f build-docker-compose.yml down'
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