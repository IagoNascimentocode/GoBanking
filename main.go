package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/accounts/services"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/database"
)

func main() {
	router := gin.Default()

	database.ConnectDataBase()

	router.GET("/accounts", services.FindAccounts)
	router.POST("/accounts", services.CrerateAccounts)
	router.GET("/accounts/:id", services.FindAccountsByID)
	router.GET("/accounts/:id/balance", services.FindBalanceByID)

	router.Run()
}
