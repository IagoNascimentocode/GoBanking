package middlewares

import (
	"github.com/gofiber/fiber"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/services"
)

/* func Auth() fiber.Handler { */

func Auth(c *fiber.Ctx) {
	const Beadrer_schema = "Bearer "
	header := c.Get("Authorization")
	if header == " " {
		c.Status(fiber.StatusBadRequest).JSON(fiber.NewError(401, "Header is not Exists"))
		return
	}

	token := header[len(Beadrer_schema):]

	if !services.NewJWTServices().ValidateToken(token) {
		c.Status(fiber.StatusBadRequest).JSON(fiber.NewError(401, "Token invalid"))
		return
	}

	c.Next()
}
