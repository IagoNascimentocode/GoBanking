package controllers

import (
	"github.com/gofiber/fiber"

	"gitlab.com/IagoNascimentocode/StoneBanking/src/database"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/models"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/services"
)

func CreateAccount(c *fiber.Ctx) {

	var account_ models.Account
	if err := c.BodyParser(&account_); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.NewError(400))
		return
	}

	if err := database.DB.Where("cpf = ?", account_.Cpf).First(&account_).Error; err == nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.NewError(401, "User already exists"))
		return
	}

	var hash, _ = services.HashPassword(account_.Secret)

	account := models.Account{
		Name:   account_.Name,
		Cpf:    account_.Cpf,
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
	if err := c.BodyParser(&account); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.NewError(400))
		return
	}

	if err := database.DB.Where("id = ?", c.Params("id")).First(&account).Error; err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.NewError(400, "User does not exist"))
		return
	}

	c.Status(200).JSON(account.Balance)
}
