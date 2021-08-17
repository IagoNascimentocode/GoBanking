package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/accounts/models"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/database"
)

func CrerateAccounts(c *gin.Context) {

	var input models.InputAccount

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account := models.Account{Name: input.Name, Secret: input.Secret, Cpf: input.Cpf}

	database.DB.Create(&account)

	c.JSON(http.StatusOK, gin.H{"data": account})

}
