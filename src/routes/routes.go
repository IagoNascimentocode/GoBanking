package routes

import (
	"github.com/gofiber/fiber"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/controllers"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/middlewares"
)

func Setup(app *fiber.App) {

	app.Post("/accounts", controllers.CreateAccount)
	app.Post("/login", controllers.Login)
	app.Get("/accounts", middlewares.Authenticate, controllers.FindAccounts)
	app.Get("/accounts/:id", middlewares.Authenticate, controllers.FindAccountsByID)
	app.Get("/accounts/:id/balance", middlewares.Authenticate, controllers.FindBalanceByID)
	app.Post("/trasnfers", middlewares.Authenticate, controllers.Transfers)
	app.Get("/trasnfers", middlewares.Authenticate, controllers.ListUserOperations)
}
