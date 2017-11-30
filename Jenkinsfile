pipeline {
    agent { docker 'golang:1.9.2' }
    stages {
        stage('Dependency') {
            steps {
                sh 'go version'
            }
        }
    }
}

