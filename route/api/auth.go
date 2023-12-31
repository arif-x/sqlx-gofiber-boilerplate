package api

import (
	controllers "github.com/arif-x/sqlx-gofiber-boilerplate/app/http/controller/auth"
	"github.com/gofiber/fiber/v2"
)

func Auth(a *fiber.App) {
	auth := a.Group("/api/v1/auth")

	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)
}
