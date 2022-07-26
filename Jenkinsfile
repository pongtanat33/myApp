pipeline {
    agent any
    environment {
        registry = "pongtanat/todoapp"
        registryCredential = 'docker_hub'
        production_server_ip = "root@167.99.70.176"
    }
    stages {
        stage('Build latest version') {
            steps { 
                script { 
                    dockerImage = docker.build registry + ":latest" 
                }
            }
        }
        stage('Push latest version') {
            steps { 
                script { 
                    docker.withRegistry( '', registryCredential ) { 
                        dockerImage.push() 
                    }
                } 
            }
        }
        stage('Test run docker') {
            steps {
                sh 'docker run -d -p 1234:1234 -p 80:80 --name myApp pongtanat/todoapp:latest'
				sh 'docker rm -f myApp'
                sh 'docker image rm -f pongtanat/todoapp:latest'
            }
        }
        stage('Prepare docker-compose file'){
            steps{
                sshagent(['producton_credentials']) {
                    sh 'ssh -o StrictHostKeyChecking=no ${production_server_ip} mkdir -p /home/dong/todoapp'
                    sh 'scp -o StrictHostKeyChecking=no docker-compose.yml ${production_server_ip}:/home/dong/todoapp/docker-compose.yml'
                    sh 'scp -o StrictHostKeyChecking=no .env ${production_server_ip}:/home/dong/todoapp/.env'
                }
            }
        }
        stage('Deploy to production server') {
            steps {
                sshagent(['producton_credentials']) {
                    sh 'ssh -o StrictHostKeyChecking=no ${production_server_ip} docker-compose -f /home/dong/todoapp/docker-compose.yml down'
                    sh 'ssh -o StrictHostKeyChecking=no ${production_server_ip} docker pull pongtanat/todoapp'
                    sh 'ssh -o StrictHostKeyChecking=no ${production_server_ip} docker-compose -f \
                    /home/dong/todoapp/docker-compose.yml --env-file /home/dong/todoapp/.env  up -d'
                 }
            }
        }
    }
}