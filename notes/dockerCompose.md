# What is Docker Compose?

- tool for defining and running multi-container docker applications.
- use yml files to configure application services: **docker-compose.yml**
- can start all services with single command: **docker-compose up**
- can stop all services with single command: **docker-compose down**

# Docker-Compose using scale

> scale up selected services when required.

- **docker-compose up -d --scale database=4**

- it creates 4 instances of database.

# Steps to follow

> Install docker compose

- already installed on windows and mac with docker
- check by using **docker-compose -v** command.
  > 2 ways to install for linux system
  - Latest [Release](https://github.com/docker/compose/releases)
  - Using PIP (**pip install -U docker-compose**)

> Create docker [compose file](example/docker-compose.yml) at any location on your system.

- **docker-compose.yml**

> Check the validity of file by command.

- **docker-compose config**

> Run docker-compose.yml file by command.

- **docker-compose up**

> Bring down application by command.

- **docker-compose down**
