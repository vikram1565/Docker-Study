# What are docker containers? How to create docker containers?

- Containers are running instance of docker images.
- A container image is a lightweight, stand-alone, executable package of a piece of software that includes everything needed to run it.

> Features of containers

- Lightweight.
- less resources are required.
- Booting of containers is very fast.
- Can start,stop,kill,remove containers very easily & quickly.
- OS resources can be shared within docker.

> Containers run on the same machine sharing the same OS kernel , this makes it faster.

> You can use command **docker container create** to create container in stopped state.

**Useful Docker Containers Commands**

- docker ps
- docker ps -a

- docker run imageName/imageID

- docker start containerName/containerID
- docker stop containerName/containerID

- docker pause containerName/containerID
- docker unpause containerName/containerID

- docker top containerName/containerID
- docker stats containerName/containerID
- docker attach containerName/containerID

- docker kill containerName/containerID
- docker rm containerName/containerID

- docker history containerName/containerID
