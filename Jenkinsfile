pipeline {
  agent {
    docker {
      image 'golang:1.12'
    }

  }
  stages {
    stage('Build') {
      agent {
        docker {
          image 'golang:1.12'
        }

      }
      steps {
        sh 'go build . -o main'
      }
    }
  }
}