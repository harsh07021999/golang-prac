package models

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func ConnectDB() {
	// os.Getenv("DB")
	env, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal("Error reading env file")
	}
	// fmt.Println(env["DB"])
	dsn := "user=" + env["user"] + " dbname=" + env["dbname"] + " sslmode=" + env["sslmode"] + " password=" + env["password"] + " host=" + env["host"]
	// fmt.Println(dsn)
	db, err := sqlx.Connect(env["DB"], dsn)

	if err != nil {
		log.Fatal("Error connecting DB", err)
	}
	DB = db
	// defer db.Close()

	// if err := db.Ping(); err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	log.Println("Successfuly connected")
	// }
}
