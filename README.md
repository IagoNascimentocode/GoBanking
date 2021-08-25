## 📄 About
Um site de portifólio para um aventureiro e seu *Mavic*.

---

## 🚀 Technologies


The project was developed using the following technologies:

 - [Docker](https://docs.docker.com/get-started/overview/) - Production Container
 - [Go](https://golang.org/doc/) - Language
 - [Gorm](https://gorm.io/docs/) - Database Orchestrator
 - [Fiber](https://docs.gofiber.io/) - Framework 
 - [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) - Cryptographic library 
 - [Json Web token](https://docs.gofiber.io/) - Library for managing authentication tokens
 ---
<br>

##  📁 How to download and run the project on a local machine 

<br>

```bash

    #Clone repository
    $ git clone https://github.com/IagoNascimentocode/GoBanking.git

    #Enter directory  
    $ cd GoBanking

    #Install dependencies
    $ go mod download 

    #Start Project
    $ go run main.go
```

## 🐋 Run an API with the Docker
<br>

```bash
   # Run DockerFile with the command. It will build our image.
   $ docker build -t imageName .

   # Up the container and run the project
   $ docker-compose up

   #See the containers that are running
   $ docker ps

   #Access the terminal of the application running in the docker
   $ docker logs banking
```