package routes

import "github.com/gofiber/fiber/v2"

func RegisterUserProfile(
	app *fiber.App,
	api fiber.Router) {

	api.Get("/profile", func(c *fiber.Ctx) error {
		uid := c.Locals("user_id")
		return c.JSON(fiber.Map{"user_id": uid})
	})
}
