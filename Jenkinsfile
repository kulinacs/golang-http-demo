pipeline {
  agent {
    docker {
      image 'golang:1.12'
    }

  }
  stages {
    stage('Build') {
      steps {
        sh 'go build . -o main'
      }
    }
  }
}
