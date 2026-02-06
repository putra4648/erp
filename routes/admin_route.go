package routes

import (
	"putra4648/erp/configs/middleware"

	"github.com/casbin/casbin/v3"
	"github.com/gofiber/fiber/v2"
)

func RegisterAdminRoutes(
	app *fiber.App,
	api fiber.Router,
	enforcer *casbin.Enforcer) {

	admin := api.Group("/admin")

	admin.Use(middleware.PermissionMiddleware(enforcer)) // Use Casbin for authorization

	admin.Get("/dashboard", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "Welcome Admin!"})
	})
}
