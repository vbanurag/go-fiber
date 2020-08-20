package routes

import (
	"github.com/gofiber/fiber"
	"github.com/vbanurag/go-fiber/controllers"
)

func RoutesHandler(api fiber.Router) {
	users(api)
}

func users(api fiber.Router) {
	users := api.Group("/users")

	users.Get("/", controllers.GetAllUsers)
	// users.Get("/:id", Controller.GetUser)
	// users.Post("/", Controller.AddUser)
	// users.Put("/:id", Controller.EditUser)
	// users.Delete("/:id", Controller.DeleteUser)
}
