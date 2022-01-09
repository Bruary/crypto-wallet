package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConnectToDB() {

	var user = os.Getenv("db_user")
	var password = os.Getenv("db_password")
	var port = os.Getenv("db_port")
	var host = os.Getenv("db_host")
	var dbName = os.Getenv("db_name")

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected!")
}
