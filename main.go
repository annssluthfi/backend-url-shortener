package main

import (
	"api-url-shortener/database"
	"api-url-shortener/handler"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.GetDBConnection()
	app := fiber.New()

	app.Get("/links", handler.GetAllLinks)
	app.Get("/links/:shortened_link", handler.RedirectByShortenedLink)
	app.Post("/links", handler.CreateLink)
	app.Put("/links/:shortened_link", handler.UpdateByShortenedLink)
	app.Delete("/links/:shortened_link", handler.DeleteByShortenedLink)
	app.Listen(":8080")

}
