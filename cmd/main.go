package cmd

import (
	"GO_Rest_API/pkg/handlers"
	"github.com/gofiber/fiber"
	_ "github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	app.Get("/book/:id?", handlers.GetBook)
	//app.Get("/books", handlers.GetAllBooks)
	app.Post("/book", handlers.AddBook)
	app.Put("/book/:id", handlers.UpdateBook)
	app.Delete("/book/:id", handlers.DeleteBook)
	app.Listen("8000")
}
