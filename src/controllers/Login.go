package controllers

import (
	"github.com/gofiber/fiber"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/database"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/models"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/services"
)

func Login(c *fiber.Ctx) {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.NewError(400))
		return
	}

	credential := models.Credentials{
		Cpf:    data["cpf"],
		Secret: data["secret"],
	}

	var account models.Account

	if err := c.JSON(&account); err != nil {
		c.Status(400)
		return
	}

	if err := database.DB.Where("cpf = ?", credential.Cpf).First(&account).Error; err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.NewError(401, "Password or cpf incorrect"))
		return
	}

	if account.ID == 0 {
		c.Status(fiber.StatusNotFound).JSON(fiber.NewError(401, "User not found"))
		return
	}

	if err := services.CheckPasswordHash(credential.Secret, account.Secret); err {
		c.Status(fiber.StatusBadRequest).JSON(fiber.NewError(401, "Password or cpf incorrect compare"))
		return
	}

	token, err := services.NewJWTServices().GenerateToken(account.ID)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.NewError(500))
		return
	}

	c.JSON(token)

}
