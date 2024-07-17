def COLOR_MAP = [
    'SUCCESS': 'good',
    'FAILURE': 'danger',
]
pipeline {
    agent any

    tools {
        go 'go22.0'
    }

    stages {
        stage('Build') {
            steps {
                sh '''go mod init xling.online
                go mod tidy
                go build main.go
                '''
            }
            post{
                success{
                    archiveArtifacts artifacts: 'main' 
                }
            }
        }
    }
    
    post {
        always {
            slackSend channel: '#jenkins-ci', color: COLOR_MAP[currentBuild.currentResult],
            message: "Find Status of Pipeline:- ${currentBuild.currentResult} ${env.JOB_NAME} ${env.BUILD_NUMBER} ${BUILD_URL}"
            cleanWs()
        }
    }
}
