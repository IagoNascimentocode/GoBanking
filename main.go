package main

import (
	"github.com/gofiber/fiber"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/database"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/routes"
)

func main() {
	app := fiber.New()

	database.ConnectDataBase()

	routes.Setup(app)

	app.Listen(":8080")
}
