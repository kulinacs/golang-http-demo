pipeline {
    agent {
        docker {
            image 'golang:1.12'
        }

    }
    stages {
        stage('Build') {
            steps {
                sh 'go build -o main .'
            }
        }
        stage('Test') {
            steps {
                sh 'go test'
            }
        }
        stage('Deliver for development') {
            when {
                branch 'development'
            }
            steps {
                input message: 'Deliver to develop? (Click "Proceed" to continue)'
            }
        }
        stage('Deliver for production') {
            when {
                branch 'master'
            }
            steps {
                input message: 'Deliver to develop? (Click "Proceed" to continue)'
            }
        }
    }
}
