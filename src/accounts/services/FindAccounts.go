package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/accounts/models"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/database"
)

func FindAccounts(c *gin.Context) {

	var accounts []models.Account

	database.DB.Find(&accounts)

	c.JSON(http.StatusOK, gin.H{"data": accounts})
}
