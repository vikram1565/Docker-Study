# Docker Architecture

**Docker Workflow**

![DockerWorkflow](assets/dockerWorkflow.PNG)

**Virtualization vs Containerization**

![VirtualizationVScontainerization](assets/virtualizationVScontainerization.PNG)

- In Virtualization, VMs having there own os and does not uses the host os.
  VMs resource allocation is fixed & does not change as per applications need. so there is lot of wastage of memory & space.

- In Containerization, containers are light weight alternative to VMs. we having containers and uses the host os. so the space, memory & other resources are not fixed. It will be taken as per the applications need so there is no overhead and its very fast.

**Docker Architecture**

![DockerArchitecture](assets/dockerArchitecture.PNG)

**Docker File**

- A dockerfile is a text file that contains the necessary commands to assemble an image. Once a Dockerfile is created, the developer uses the docker build command to create an image based on the commands within the file.

**Docker Images**

- An image is a read-only template with instructions for creating a Docker container. Often, an image is based on another image, with some additional customization. For example, you may build an image which is based on the ubuntu image, but installs the Apache web server and your application, as well as the configuration details needed to make your application run.

- You might create your own images or you might only use those created by others and published in a registry. To build your own image, you create a Dockerfile with a simple syntax for defining the steps needed to create the image and run it. Each instruction in a Dockerfile creates a layer in the image. When you change the Dockerfile and rebuild the image, only those layers which have changed are rebuilt. This is part of what makes images so lightweight, small, and fast, when compared to other virtualization technologies.

**Docker Containers**

- A container is a runnable instance of an image. You can create, start, stop, move, or delete a container using the Docker API or CLI. You can connect a container to one or more networks, attach storage to it, or even create a new image based on its current state.

- By default, a container is relatively well isolated from other containers and its host machine. You can control how isolated a container’s network, storage, or other underlying subsystems are from other containers or from the host machine.

- A container is defined by its image as well as any configuration options you provide to it when you create or start it. When a container is removed, any changes to its state that are not stored in persistent storage disappear.

**Docker Registry/Hub**

- A Docker registry stores Docker images. Docker Hub is a public registry that anyone can use, and Docker is configured to look for images on Docker Hub by default. You can even run your own private registry. If you use Docker Datacenter (DDC), it includes Docker Trusted Registry (DTR).

- When you use the docker pull or docker run commands, the required images are pulled from your configured registry. When you use the docker push command, your image is pushed to your configured registry.

**Docker Client**

- The Docker client (docker) is the primary way that many Docker users interact with Docker. When you use commands such as docker run, the client sends these commands to dockerd, which carries them out. The docker command uses the Docker API. The Docker client can communicate with more than one daemon.

**Docker Daemon(Server)**

- The Docker daemon (dockerd) listens for Docker API requests and manages Docker objects such as images, containers, networks, and volumes. A daemon can also communicate with other daemons to manage Docker services.

**Docker Engine**

Docker Engine is a client-server application with these major components -

- A server which is a type of long-running program called a daemon process (the dockerd command).

- A REST API which specifies interfaces that programs can use to talk to the daemon and instruct it what to do.

- A command line interface (CLI) client (the docker command).

  ![DockerEngine](assets/dockerEngine.PNG)

- The CLI uses the Docker REST API to control or interact with the Docker daemon through scripting or direct CLI commands. Many other Docker applications use the underlying API and CLI.

* The daemon creates and manages Docker objects, such as images, containers, networks, and volumes.
* Dangling image is one that is not tagged and is not referenced by any container.
