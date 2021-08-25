## 📄 About
Go banking is a transfer api between Internal accounts of a digital bank.

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
<br>

## 🐋 Run an API with the Docker

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

<br>

# <div align=center><strong>🗺️ Routes and features</strong></div>


```bash
app.Post("/accounts", controllers.CreateAccount)
 ```
* This route is responsible for creating a user and saving it in the database.
* The CPF must be unique, a user must not be able to create an account with an existing CPF.
<br>
<br>

```bash
app.Post("/login", controllers.Login)
```
* This route is responsible for authenticating the user.
* Generates and returns a Json Web Token.
<br>
<br>

```bash
app.Get("/accounts", middlewares.Authenticate, controllers.FindAccounts)
```
* This route is responsible for listing all saved accounts.
* Only an Authenticated user can list all users.
<br>
<br>

```bash
app.Get("/accounts/:id", middlewares.Authenticate, controllers.FindAccountsByID)
```
* This route is responsible for listing a user that is specified by the ID parameter.
* Only an Authenticated user can list all users.
<br>
<br>

```bash
app.Get("/accounts/:id/balance", middlewares.Authenticate, controllers.FindBalanceByID)
```
* This route is responsible for listing the balance of an account passing the user id as a parameter.
* Only an Authenticated user can list all users.
<br>
<br>

```bash
app.Post("/trasnfers", middlewares.Authenticate, controllers.Transfers)
```
* This route must be able to transfer an amount from the authenticated user to a user defined with the parameter ID in the request body.
* The amount cannot be greater than the balance of the authenticated user.
* Only an Authenticated user can list all users.
<br>
<br>

```bash
app.Get("/trasnfers", middlewares.Authenticate, controllers.Transfers)
```
* It is routed and responsible for returning all sending operations that the authenticated user did.
* Only an Authenticated user can list all users.
<br>
<br>
