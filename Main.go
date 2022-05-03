package main

import (
	"GO_Rest_API/pkg/handlers"
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber"
)

var db *sql.DB

var server = "<your_server.database.windows.net>"
var port = 1433
var user = "<your_username>"
var password = "<your_password>"
var database = "<your_database>"

func main() {

	// Build connection string
	//connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
	//	server, user, password, port, database)
	//
	//var err error
	//
	//// Create connection pool
	//db, err = sql.Open("sqlserver", connString)
	//if err != nil {
	//	log.Fatal("Error creating connection pool: ", err.Error())
	//}
	//ctx := context.Background()
	//err = db.PingContext(ctx)
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	fmt.Printf("Connected!\n")
	app := fiber.New()
	app.Get("/book", handlers.GetBook)
	app.Get("/books", handlers.GetAllBooks)
	app.Get("/ping", handlers.GetHealth)
	app.Post("/book", handlers.AddBook)
	app.Put("/book", handlers.UpdateBook)
	app.Delete("/book", handlers.DeleteBook)
	fmt.Println("Now listening : 8000")
	app.Listen("8000")

}
