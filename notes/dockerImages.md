# What are docker images? How to run docker images?

- Docker images are templates used to create docker containers.
- Containers is running instance of images.
- Docker images are stored in registries(ex. Docker Hub). it can be stored locally or remote.
- Docker can build images automatically by reading the instructions from a DockerFile.
- A single image can be used to create multiple containers.

**Useful Docker Images Commands**

- docker images -help
- docker images
- docker pull imageName
- docker pull imageName:tag
- docker images -q
- docker images -f "dangling=false"
- docker images -f "dangling=false" -q
- docker inspect
- docker history imageName
- docker run imageName/id
- docker rmi image
- docker rmi -f image
