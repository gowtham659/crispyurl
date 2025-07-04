package main

import (
	"urlshortnerreactgo/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main(){
	app := fiber.New()
	//hadling cors
	app.Use(cors.New())

	app.Post("/shorten",handler.ShortenUrl)
	app.Use("/:shortcode",handler.RedirectURL)
	app.Listen(":8989")
}