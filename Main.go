package main

import (
	"GO_Rest_API/pkg/handlers"
	"fmt"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	app.Get("/book/:id?", handlers.GetBook)
	app.Get("/ping", handlers.GetHealth)
	//app.Get("/books", handlers.GetAllBooks)
	app.Post("/book", handlers.AddBook)
	app.Put("/book/:id", handlers.UpdateBook)
	app.Delete("/book/:id", handlers.DeleteBook)
	fmt.Println("Now listening : 8000")
	app.Listen("8000")

}
