package routes

import (
	"github.com/gofiber/fiber"
	"gitlab.com/IagoNascimentocode/StoneBanking/src/controllers"
)

func Setup(app *fiber.App) {

	app.Post("/accounts", controllers.CreateAccount)
	app.Get("/accounts", controllers.FindAccounts)
	app.Get("/accounts/:id", controllers.FindAccountsByID)
	app.Get("/accounts/:id/balance", controllers.FindBalanceByID)
	app.Post("/login/", controllers.Login)
}
