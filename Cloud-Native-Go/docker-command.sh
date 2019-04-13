## Get golang:1.12.3-alpine3.9
docker pull golang:1.12.3-alpine3.9

## Build Dockerfile
docker build -t cloud-native-go:1.0.0 .

## docker tag [DOCKER_REPOSITORY_NAME:TAG] [DOCKER_ACCOUNT/[DOCKER_REPOSITORY_NAME:TAG]]
docker tag cloud-native-go:1.0.0 jieqiong/cloud-native-go:1.0.0

## Login the docker hub and do the push thing
docker login
docker push jieqiong/cloud-native-go:1.0.0

## Run cloud-native-go container
docker run -it -p 8080:8080 cloud-native-go:1.0.0

## Get the container id
docker ps
## Get the container ip: `2e8ba2e6df90` is the container id
docker inspect 2e8ba2e6df90 | grep '"IPAddress"' | head -n 1