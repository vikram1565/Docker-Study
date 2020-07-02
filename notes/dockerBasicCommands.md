# Docker Basic Commands

**Basic**

> docker version

- to get the information of docker client & docker server use this command.

> docker -v

- this command gives us the docker version.
- you can also use `docker --version` to get docker version.

> docker info

- This command gives us the detailed information about docker that is installed on your machine.

> docker --help

- use this command to get information of any other commands.

> docker login

- use this command to login to your https://hub.docker.com
- login with your docker id to push and pull images from docker hub.

**Images**

> docker images

- this command gives us the list of all images.
- to get docker image id use `docker images -q` command.

> docker pull

- this command is used to pull the images from docker hub.
- to pull docker image use `docker pull imageName:tag` command.

> docker rmi

- this command is used to remove one or more images.
- pass docker image id to rmi command to remove particular image.

> docker inspect

- return low -level information on docker objects.

**Containers**

> docker ps

- used to list the containers.
- for more options related to ps use `docker ps --help` command.

> docker run

- this command is used to run in a new container.

> docker start

- this command is used to start one or more stopped containers.

> docker stop

- this command is used to stop one or more running containers.

**System**

> docker stats

- this command is used to display a live stream of containers resource usage statistics.

> docker system df

- this command is used to show docker disk usage.

> docker system prune

- this command is used to remove unused data.
- be careful while using this command.
