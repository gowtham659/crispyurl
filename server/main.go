package main

import (
	handler "urlshortnerreactgo/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	//hadling cors
	//app.Use(cors.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000/, https://react-go-oracle-app-1.onrender.com/",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("request success")
	})
	app.Post("/shorten", handler.ShortenUrl)
	app.Use("/:shortcode", handler.RedirectURL)
	app.Listen(":8989")
}
