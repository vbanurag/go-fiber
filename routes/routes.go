package routes

import (
	"github.com/gofiber/fiber"
)

func RoutesHandler(api fiber.Router) {
	users(api)
}
