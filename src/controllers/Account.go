package controllers

import (
	"github.com/gofiber/fiber"

	"gitlab.com/IagoNascimentocode/StoneBanking/src/database"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/models"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/services"
)

func CreateAccount(c *fiber.Ctx) {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.NewError(400))
		return
	}

	var hash, _ = services.HashPassword(data["secret"])

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
		c.Status(fiber.StatusBadRequest).JSON(fiber.NewError(400, "User does not exist"))
		return
	}

	c.Status(200).JSON(account)
}

func FindBalanceByID(c *fiber.Ctx) {

	var account models.Account
	if err := c.JSON(&account); err != nil {
		c.Status(400)
		return
	}

	if err := database.DB.Where("id = ?", c.Params("id")).First(&account).Error; err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.NewError(400, "User does not exist"))
		return
	}

	c.Status(200).JSON(account.Balance)
}

func Login(c *fiber.Ctx) {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.NewError(400))
	}

	credential := models.Credentials{
		Cpf:    data["cpf"],
		Secret: data["secret"],
	}

	var account models.Account

	if err := c.JSON(&account); err != nil {
		c.Status(400)
	}

	if err := database.DB.Where("cpf = ?", credential.Cpf).First(&account).Error; err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.NewError(400, "Password or cpf incorrect"))
		return
	}

	if err := services.CheckPasswordHash(credential.Secret, account.Secret); err {
		c.Status(fiber.StatusBadRequest).JSON(fiber.NewError(400, "Password or cpf incorrect compare"))
		return
	}

	c.Status(200).JSON(account.Secret)

}
