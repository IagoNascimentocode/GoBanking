package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/accounts/services"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/database"
)

func main() {
	router := gin.Default()

	database.ConnectDataBase()

	router.GET("/Accounts", services.FindAccounts)
	router.GET("/Accounts/:id/balance", services.FindAccountsByID)
	router.GET("/findBalance/:id", services.FindBalanceByID)
	router.POST("/createAccount", services.CrerateAccounts)

	router.Run()
}
