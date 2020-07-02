# What is docker volume?

- By default all files created inside a container are stored on a writable container layer.
- The data doesn't persist when that container is no longer running.
- A container's writable layer is tightly coupled to the host machine where the container is running. you can't easily move data somewhere else.
- Docker has two options for containers to store files in the host machine, so that the files are persisted even after the container stops.
- Volumes are stored in a part of the host filesystem which is managed by docker and isolated from the core functionality of the host machine.
- Non-Docker processes should not modify this part of the filesystem.
- Volumes are the best way to persist data in Docker.
- A given volume can be mounted into multiple containers simultaneously.
- When no running container is using a volume , the volume is still available to docker and is not removed automatically. you can remove unused volumes using **docker volume prune**
- When you mount a volume, it may be named or anonymous.
- Anonymous volumes are not given an explicit name when they are first mounted into container.
  -Volumes also support the use of volume drivers, which allow you to store your data on remote hosts or cloud providers, among other possibilities.

# Useful commands

> docker volume

- get the docker volume information.

> docker volume create volume_name

- this command is used to create volume_name volume.

> docker volume ls

- this command is used to list the docker volumes.

> docker volume inspect volume_name

- this command is used to display detailed information of one or more volumes.

> docker volume prune

- this command is used to remove all unused local volumes.

> docker volume rm volume_name

- this command is used to remove one or more volumes.
