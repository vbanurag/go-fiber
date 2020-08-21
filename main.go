package main

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/vbanurag/go-fiber/configuration"
	"github.com/vbanurag/go-fiber/routes"
)

func main() {
	app := fiber.New()
	config, err := configuration.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	api := app.Group("/api")

	routes.RoutesHandler(api)

	app.Use(func(c *fiber.Ctx) {
		c.SendStatus(404)
		if err := c.Render("errors/404", fiber.Map{}); err != nil {
			c.Status(500).Send(err.Error())
		}
	})

	_err := app.Listen(config.Port)

	if _err != nil {
		// Exit the application
		log.Fatal(_err)
	}
}
