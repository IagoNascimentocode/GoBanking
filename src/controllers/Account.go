package controllers

import (
	"github.com/gofiber/fiber"

	"gitlab.com/IagoNascimentocode/StoneBanking/src/database"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateAccount(c *fiber.Ctx) {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(data["secret"]), 8)

	account := models.Account{
		Name:   data["name"],
		Cpf:    data["cpf"],
		Secret: string(hash),
	}

	database.DB.Create(&account)

	c.Status(201).JSON(account)

}

func FindAccounts(c *fiber.Ctx) {

	var accounts []models.Account

	database.DB.Find(&accounts)

	c.Status(200).JSON(accounts)
}

func FindAccountsByID(c *fiber.Ctx) {

	var account models.Account

	if err := database.DB.Where("id = ?", c.Params("id")).First(&account).Error; err != nil {
		c.Status(400)
	}

	c.Status(200).JSON(account)
}

func FindBalanceByID(c *fiber.Ctx) {

	var account models.Account
	if err := c.JSON(&account); err != nil {
		c.Status(400)
	}

	if err := database.DB.Where("id = ?", c.Params("id")).First(&account).Error; err != nil {
		c.Status(400)
	}

	c.Status(200).JSON(account.Balance)
}
