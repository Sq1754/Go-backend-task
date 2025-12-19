package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sq1754/user-age-api/internal/handler"
)

func RegisterUserRoutes(app *fiber.App, h *handler.UserHandler) {
	app.Post("/users", h.CreateUser)
	app.Get("/users/:id", h.GetUserByID)
	app.Put("/users/:id", h.UpdateUser)
	app.Delete("/users/:id", h.DeleteUser)
	app.Get("/users", h.ListUsers)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

}
