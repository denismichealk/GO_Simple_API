package MS_SQL_DataAccess

import (
	"context"
	_ "context"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

var db *sql.DB

var server = "CHAPCHAP-001\\SQLEXPRESS"
var port = 1433
var user = "HP"
var password = ""
var database = "local_db"

func CreateDBConnection() (*sql.DB, error) {
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	var err error

	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")

	return db, nil
}
