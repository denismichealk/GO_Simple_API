package cmd

import (
	"github.com/gofiber/fiber"
	_ "github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	app.Get("/book/:id?", GetBook)
	app.Get("/books", GetAllBooks)
	app.Post("/book", AddBook)
	app.Put("/book/:id", UpdateBook)
	app.Delete("/book/:id", DeleteBook)
	app.Listen("8000")
}
