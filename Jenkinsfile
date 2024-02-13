pipeline {
    agent any

    environment {
        DOCKER_IMAGE_NAME = 'muhhylmi/golang-api'
        DOCKER_REGISTRY_URL = 'hub.docker.com'
        DOCKER_PORT_MAPPING = '3000:3000'
        DOCKER_CONTAINER_NAME = 'golang-api'
        DOCKER_NETWORK = 'app1-network'
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scmGit(branches: [[name: '*/dev']], extensions: [], userRemoteConfigs: [[url: 'https://github.com/muhhylmi/golang-api.git']])
            }
        }

        stage('Build') {
            steps {
                script {
                    sh '''
                        export PATH=$PATH:/usr/local/go/bin
                        go build -o app ./bin
                    '''
                }
            }
        }

        stage('Docker Build') {
            steps {
                script {
                    // Build Docker image
                    sh "docker build -t ${DOCKER_IMAGE_NAME} ."
                }
            }
        }

        stage('Test') {
            steps {
                script {
                    // Run tests (if applicable)
                    sh '''
                        export PATH=$PATH:/usr/local/go/bin
                        go test ./...
                    '''
                }
            }
        }
        
        stage('Push to Registry') {
            steps {
                withCredentials([string(credentialsId: 'DOCKER_PASSWORD', variable: 'DOCKER_PASSWORD'), string(credentialsId: 'DOCKER_USERNAME', variable: 'DOCKER_USERNAME')]) {
                    script {
                        // Gunakan echo untuk mengirimkan kata sandi melalui stdin
                        sh "echo ${DOCKER_PASSWORD} | docker login -u ${DOCKER_USERNAME} --password-stdin"
                        sh "docker push ${DOCKER_IMAGE_NAME}:latest"                
                    }
                }
            }
        }

        stage('Deploy') {
            steps {
                script {
                    // Push Docker image to registry or deploy to Kubernetes, etc.
                    // Example: docker push your-registry/${DOCKER_IMAGE_NAME}:latest
                    sh "docker stop ${DOCKER_CONTAINER_NAME} || true"
                    sh "docker rm ${DOCKER_CONTAINER_NAME} || true"

                    // Pull the latest Docker image from the registry
                    sh "docker pull ${DOCKER_IMAGE_NAME}"

                    // Run the updated Docker container
                    sh "docker run -d --name ${DOCKER_CONTAINER_NAME} -p ${DOCKER_PORT_MAPPING} --env-file .env --network=${DOCKER_NETWORK} ${DOCKER_IMAGE_NAME}"
                }
            }
        }
    }

    post {
        success {
            echo 'Build and deployment successful! You can now access your application.'
        }
        failure {
            echo 'Build or deployment failed! Please check the logs for more information.'
        }
    }
}
