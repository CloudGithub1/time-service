🕒 time-service

A minimal Go-based microservice that returns the current timestamp and the client’s IP address in JSON format.
Example output:
{
  "timestamp": "2025-01-01T12:00:00Z",
  "ip": "127.0.0.1"
}

Features
Built using Go (static binary)
Ultra-lightweight container using scratch base image
Runs as non-root user inside the container

Simple to build & run using only:
docker build
docker run

1. Prerequisites
Before running this project locally, ensure the following tools are installed:
Git
Download: https://git-scm.com/downloads
Docker Desktop
Download: https://www.docker.com/products/docker-desktop/
Verify Docker is running:
docker --version

2. Clone the Repository
git clone https://github.com/CloudGithub1/time-service.git
cd time-service

3. Build the Docker Image
docker build -t time-service .

This will:
Build the Go application
Create a static binary
Produce a tiny scratch container image named time-service

4. Run the Container
docker run -p 8080:8080 time-service
Expected log:
time-service starting on :8080
The container will continue running in your terminal.

5. Test the Service
Option A — Browser
Visit:
http://localhost:8080/

Option B — curl
curl http://localhost:8080/

Expected response:

{
  "timestamp": "2025-01-01T12:00:00Z",
  "ip": "127.0.0.1"
}

6. Verify NON-ROOT Container (Mandatory)

Open a new terminal:
docker image inspect time-service --format '{{.Config.User}}'
Expected output:
10001

This confirms the container is running as a non-root user, following best security practices.

7. Test Using Published Docker Image (Optional)
If you do not want to build the image locally, test directly from Docker Hub.

Step 1 — Pull the image
docker pull clouddockerhub1/time-service:v1.0

Step 2 — Run the container
docker run -p 8080:8080 clouddockerhub1/time-service:v1.0

Step 3 — Test in browser
http://localhost:8080/
Or:
curl http://localhost:8080/


8. Optional: Using This Image in ECS or EKS
 A) Deploying in Amazon ECS

Update your task definition JSON:

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

Deploy the updated task definition revision to your ECS service.

 B) Deploying in Kubernetes (EKS)

Add/update your deployment.yaml:

spec:
  containers:
    - name: time-service
      image: clouddockerhub1/time-service:latest
      ports:
        - containerPort: 8080

Apply:

kubectl apply -f deployment.yaml
Expose it using a LoadBalancer Service:
kubectl expose deployment time-service --type=LoadBalancer --port=80 --target-port=8080

9. Cleanup (Optional)
Stop all running containers:
docker ps
docker stop <container-id>

Remove local images:
docker rmi time-service

Completed Requirements

✔ Minimal microservice returning timestamp + IP
✔ Proper Dockerfile with best practices
✔ Image runs as non-root user
✔ Very small image size (scratch)
✔ Clear documentation for beginners
✔ Ready for cloud platforms (ECS/EKS)
✔ Build & run using only two commands:
docker build + docker run