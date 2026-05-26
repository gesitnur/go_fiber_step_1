package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	var err error
	dsn := "host=localhost user=gesit password=12345678 dbname=gemini_pengeluaran port=5432 sslmode=disable"
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Println("error while starting the psql server: ", err)
		return err
	}
	if err = DB.Ping(); err != nil {
		log.Println("error making a test ping to the server: ", err)
		return err
	}
	log.Println("Database connected successfully")
	return nil
}
