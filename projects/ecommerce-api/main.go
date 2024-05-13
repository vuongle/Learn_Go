package main

import (
	"database/sql"
	"ecommerce-api/cmd/api"
	"ecommerce-api/config"
	"ecommerce-api/db"
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {

	// Init database
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Env.DBUser,
		Passwd:               config.Env.DBPassword,
		Addr:                 config.Env.DBAddr,
		DBName:               config.Env.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}
	initStorage(db)

	// Init the server
	server := api.NewAPIServer(
		":8080", nil,
	)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("CONNECTED TO MYSQL")
}
