package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/accounts/controllers"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/database"
)

func main() {
	router := gin.Default()

	database.ConnectDataBase()

	router.GET("/", controllers.FindAccounts)

	router.Run()
}
