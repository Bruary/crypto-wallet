package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func ConnectToDB() {

	var user = os.Getenv("db_user")
	var password = os.Getenv("db_password")
	var port = os.Getenv("db_port")
	var host = os.Getenv("db_host")
	var dbName = os.Getenv("db_name")

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	// open database
	Db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}

	// close database
	defer Db.Close()

	// check db
	err = Db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected!")

	createUsersTableQuery := `CREATE TABLE IF NOT EXISTS Users(
		ID int PRIMARY KEY,
		First_Name TEXT NOT NULL,
		Last_Name TEXT NOT NULL,
		Full_Name TEXT NOT NULL,
		Email TEXT NOT NULL UNIQUE,
		Age INT NOT NULL,
		Created_at TIMESTAMP,
		Updated_at TIMESTAMP,
		Deteled_at TIMESTAMP
		)`

	_, err = Db.Exec(createUsersTableQuery)
	if err != nil {
		panic(err)
	}

	fmt.Println("Users Table Created/Found.")

}
