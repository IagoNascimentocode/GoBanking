package main

import (
	"github.com/gofiber/fiber"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/database"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/routes"
)

func main() {

	database.ConnectDataBase()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8080")
}
