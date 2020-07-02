# Docker Overview

- Docker is the worlds leading software container platform.
- Docker is a tool designed to make it easier to deploy & run applications by using containers.

# Installation

- [Install docker](https://docs.docker.com/get-docker/)

- Check docker version

```
docker version
```

- Run sample application

```
docker run hello-world
```

# Images, Containers

- Docker images are templates used to create docker containers.
- Containers is running instance of images.

- See all images

```
docker images
```

- See all containers

```
docker ps -a
```

- See all running containers

```
docker ps
```

- Run docker image

```
docker run imageName/ID
```

# Create Dockerfile

- Create dockerfile.
- Build an image from dockerfile.

- Create Dockerfile and add these lines in dockerfile

```
FROM golang:latest
MAINTAINER  "Vikram Ingawale" <vikram.ingawale91@gmail.com>
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
CMD ["app"]
```

# Build an image from the Dockerfile

- Build the image using the docker build command.

```
docker build -t myhelloworld .
```

- `-t` parameter gives tag to yout image.
- dot(.) at the end means this command reads the Dockerfile from the current directory otherwise specify Dockerfile path.

# create a container from your image & run

- to create container use command.

```
docker run -d -p 80:80 imageName
```

- `-d` means run this detached, as a daemon not dependent on the terminal session.
- `-p` means map ports.
- access it with http://localhost:80

# Push docker image

- Create [docker hub](https://hub.docker.com/) account.
- Run command in docker cli

```
docker login
```

- Add a docker tag to image.

```
docker tag <imageName/ID>  <docker hub username>/<target imageName>:<version or tag>
```

- Push your image.

```
docker push <docker hub username>/<target imageName>
```

- Look at your docker hub account. Your image should be there.

- See images on your machine

```
docker images
```

# Pull docker image

- To pull image from docker hub use below commands

```
docker run -d -p 80:80 <docker hub username>/<imageName>
```

```
docker run imageName/ID
```

# Stop running container

```
docker stop containerName/ID
```

# Remove image

```
docker rmi -f <imageName/ID>
```

- `-f` is for forcefully delete.

# Remove container

```
docker rm containerName/ID
```

> ## For more details please refer [notes](notes) section
