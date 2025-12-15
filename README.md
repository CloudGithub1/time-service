🕒 time-service

A minimal Go-based microservice that returns the current timestamp and the client’s IP address in JSON format.

Features

Built using Go (static binary)
Ultra-lightweight container using scratch base image
Runs as non-root user inside the container
Simple to build & run using only:

docker build
docker run

1️⃣ Prerequisites
Ensure the following tools are installed locally:

Git
Download: https://git-scm.com/downloads

Docker Desktop
Download: https://www.docker.com/products/docker-desktop/

Verify Docker is running:

docker --version

2️⃣ Clone the Repository
git clone https://github.com/CloudGithub1/time-service.git
cd time-service

3️⃣ Build the Docker Image
docker build -t time-service .


This will:
Build the Go application
Create a static binary
Produce a tiny scratch-based container image named time-service

4️⃣ Run the Container
docker run -p 8080:8080 time-service

Expected log:
time-service starting on :8080
The container will continue running in your terminal.

5️⃣ Test the Service
Option A — Browser

Open:

http://localhost:8080/

Option B — curl
curl http://localhost:8080/


Expected response:

{
  "timestamp": "2025-01-01T12:00:00Z",
  "ip": "127.0.0.1"
}

6️⃣ Verify NON-ROOT Container (Mandatory)

Open a new terminal and run:
docker image inspect time-service --format '{{.Config.User}}'
Expected output:
10001

This confirms the container is running as a non-root user, following container security best practices.

7️⃣ Test Using Published Docker Image (Optional)
If you don’t want to build locally, you can pull the image directly from Docker Hub.

Pull the image
docker pull clouddockerhub1/time-service:v1.0

Run the container
docker run -p 8080:8080 clouddockerhub1/time-service:v1.0

Test
Browser:

http://localhost:8080/
Or:
curl http://localhost:8080/

8️⃣ Optional: Using This Image in ECS or EKS
🟦 A) Amazon ECS

Example container definition:

{
  "containerDefinitions": [
    {
      "name": "time-service",
      "image": "clouddockerhub1/time-service:latest",
      "portMappings": [
        {
          "containerPort": 8080,
          "protocol": "tcp"
        }
      ]
    }
  ]
}

Deploy the new task definition revision to your ECS service.

🟩 B) Kubernetes (EKS)

Update deployment.yaml:

spec:
  containers:
    - name: time-service
      image: clouddockerhub1/time-service:latest
      ports:
        - containerPort: 8080


Apply:

kubectl apply -f deployment.yaml


Expose the service:

kubectl expose deployment time-service \
  --type=LoadBalancer \
  --port=80 \
  --target-port=8080

9️⃣ Cleanup (Optional)

Stop running containers:

docker ps
docker stop <container-id>

Remove local images:

docker rmi time-service

✅ Completed Requirements

✔ Minimal microservice returning timestamp + IP
✔ Dockerfile with best practices
✔ Runs as non-root user
✔ Very small image size (scratch)
✔ Clear beginner-friendly documentation
✔ Ready for ECS / EKS
✔ Build & run using only:

docker build
docker run

