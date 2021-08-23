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

	var t models.Transfer
	if err := c.BodyParser(&t); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.NewError(400, "PRoblema com o t"))
		return
	}

	accountDestination := models.Transfer{
		Account_origin_id:      account.ID,
		Account_destination_id: t.Account_destination_id,
		Amount:                 t.Amount,
	}

	c.JSON(accountDestination)
}
