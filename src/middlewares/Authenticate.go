package middlewares

import (
	"github.com/gofiber/fiber"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/services"
)

func Authenticate(c *fiber.Ctx) {
	/* 	const Beadrer_schema = "Beare "
	   	header := c.Get("Authorization")
	   	if header == " " {
	   		c.Status(fiber.StatusBadRequest).JSON(fiber.NewError(401, "Header is not Exists"))
	   		return
	   	}

	   	token := header[len(Beadrer_schema):]
	*/
	token := c.Cookies("jwt")

	if !services.NewJWTServices().ValidateToken(token) {
		c.Status(fiber.StatusBadRequest).JSON(fiber.NewError(401, "Token invalid"))
		return

	}
		c.Next()
}
