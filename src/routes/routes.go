package routes

import (
	"github.com/gofiber/fiber"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/controllers"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/middlewares"
)

func Setup(app *fiber.App) {

	app.Post("/accounts", controllers.CreateAccount)
	app.Post("/login", controllers.Login)
	app.Get("/accounts", middlewares.Auth, controllers.FindAccounts)
	app.Get("/accounts/:id", middlewares.Auth, controllers.FindAccountsByID)
	app.Get("/accounts/:id/balance", middlewares.Auth, controllers.FindBalanceByID)
}
