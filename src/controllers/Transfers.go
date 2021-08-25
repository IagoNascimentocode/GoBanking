package controllers

import (
	"github.com/gofiber/fiber"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/database"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/models"
)

func Transfers(c *fiber.Ctx) {
	var token = c.Cookies("jwt")

	var account models.Account

	if err := database.DB.Where("token = ?", token).First(&account).Error; err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.NewError(400, "User does not exist"))
		return
	}

	var transfer models.Transfer
	if err := c.BodyParser(&transfer); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.NewError(400))
		return
	}

	accountDestination := models.Transfer{
		Account_origin_id:      account.ID,
		Account_destination_id: transfer.Account_destination_id,
		Amount:                 transfer.Amount,
	}

	database.DB.Create(&accountDestination)

	if accountDestination.Amount > account.Balance {
		c.Status(fiber.StatusBadRequest).JSON(fiber.NewError(400, "insufficient funds"))
		return
	}
	account.Balance = account.Balance - accountDestination.Amount

	var reciver models.Account
	if err := database.DB.Where("id = ?", accountDestination.Account_destination_id).First(&reciver).Error; err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.NewError(401, "User does not exist"))
		return
	}
	reciver.Balance = reciver.Balance + accountDestination.Amount

	database.DB.Save(account)
	database.DB.Save(reciver)

	c.JSON(accountDestination)
}

func ListUserOperations(c *fiber.Ctx) {
	var token = c.Cookies("jwt")

	var account models.Account

	if err := database.DB.Where("token = ?", token).First(&account).Error; err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.NewError(401, "User does not exist"))
		return
	}

	var trasnfers []models.Transfer

	if err := database.DB.Where("account_origin_id = ?", account.ID).Find(&trasnfers).Error; err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.NewError(401, "you don't have operations"))
		return
	}

	c.JSON(trasnfers)
}
