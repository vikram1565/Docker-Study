# Docker Introduction

**What is docker**

- Docker is the worlds leading software container platform.
- Docker is a tool designed to make it easier to deploy & run applications by using containers.
- Containers allow a developer to package up an application with all of the parts it needs, such as libraries & other dependencies & ship it all out as one container in standard manner.

**Where does docker operate**

- Docker is present in the entire application workflow, but its main use in deployment.

![introduction](assets/dockerIntro1.PNG)

**Docker Installation**

- [Install Docker](https://docs.docker.com/get-docker/)

**Docker Benefits**

- Build app only once

  - An application inside a container can run on any system that has Docker installed. So there is no need to build and configure app multiple times on different platforms.

- More sleep & less worry

  - With Docker you test your application inside a container and ship it inside a container.
  - This means the environment in which you test is identical to the one on which the app will run in production

- Portability

  - Docker containers can run on any platforms. It can run on your local system, Amazon ec2, Google Cloud Platform(GCP), Rackspace server, VirtualBox..etc.

- Version Control

  - Like Git, Docker has in-built version control system. Docker containers work just like GIT repositories, allowing you to commit changes to your Docker images and version control them.

- Isolation

  - With docker every application works in isolation in its own container and does not interferes with other applications running on the same system. so multiple containers can run on same system without interference.
  - For removal also you can simply delete the container and it will not leave behind any files or traces on the system.

- Productivity
  - Docker allows faster and more efficient deployments without worrying about running your app on different platforms. It increases productivity many folds.
