package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/accounts/models"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/database"
	"golang.org/x/crypto/bcrypt"
)

func CrerateAccounts(c *gin.Context) {

	var input models.InputAccount

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(input.Secret), 8)

	account := models.Account{Name: input.Name, Cpf: input.Cpf, Secret: string(hash)}

	database.DB.Create(&account)

	c.JSON(http.StatusOK, gin.H{"data": account})

}
