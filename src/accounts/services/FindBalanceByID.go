package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/accounts/models"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/database"
)

func FindBalanceByID(c *gin.Context) {

	var account models.Account

	if err := database.DB.Where("id = ?", c.Param("id")).First(&account).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": account.Balance})
}
