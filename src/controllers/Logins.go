package controllers

import (
	"github.com/gofiber/fiber"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/database"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/models"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/services"
)

func Login(c *fiber.Ctx) {

	var credential models.Credential
	if err := c.BodyParser(&credential); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.NewError(400))
		return
	}

	login := models.Credential{
		Cpf:    credential.Cpf,
		Secret: credential.Secret,
	}

	var account models.Account
	if err := c.JSON(&account); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.NewError(400))
		return
	}

	if err := database.DB.Where("cpf = ?", login.Cpf).First(&account).Error; err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.NewError(401, "Password or cpf incorrect"))
		return
	}

	if err := services.CheckPasswordHash(credential.Secret, account.Secret); err {
		c.Status(fiber.StatusBadRequest).JSON(fiber.NewError(401, "Password or cpf incorrect"))
		return
	}

	token, err := services.NewJWTServices().GenerateToken(account.ID)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.NewError(500))
		return
	}

	if err := database.DB.Where("cpf = ?", login.Cpf).First(&account).Update("token", token).Error; err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.NewError(500))
		return
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "jwt"
	cookie.Value = token

	c.Cookie(cookie)
	c.Status(200).JSON(token)
}
