package routes

import (
	"github.com/gofiber/fiber"
	"github.com/vbanurag/go-fiber/controllers"
)

func users(api fiber.Router) {
	users := api.Group("/users")

	users.Get("/", controllers.GetAllUsers)
	users.Get("/:id", controllers.GetUser)
	users.Post("/", controllers.AddUser)
	users.Put("/:id", controllers.EditUser)
	// users.Delete("/:id", Controller.DeleteUser)
}
