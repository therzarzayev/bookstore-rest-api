def COLOR_MAP = [
    'SUCCESS': 'good',
    'FAILURE': 'danger',
]
pipeline {
    agent any

    tools {
        go 'go22.0'
    }

    environment {
        DOCKER_REGISTRY = 'registry.digitalocean.com'
        REGISTRY_NAMESPACE = 'therzarzayev'
        IMAGE_NAME = 'bookapi' 
        DOCKER_CREDENTIALS_ID = 'digitalocean-api-token'
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

        stage('Build Image'){
            agent{
                label 'docker-node'
            }
            steps{
                script{
                    dockerImage = docker.build("${DOCKER_REGISTRY}/${REGISTRY_NAMESPACE}/${IMAGE_NAME}")
                }
            }
        }

        stage('Push Image'){
            agent{
                label 'docker-node'
            }
            steps{
                script{
                    docker.withRegistry("https://${DOCKER_REGISTRY}", DOCKER_CREDENTIALS_ID) {
                        dockerImage.push("${env.BUILD_NUMBER}")
                        dockerImage.push("latest")
                    }
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
