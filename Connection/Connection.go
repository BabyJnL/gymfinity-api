package Connection

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	user		= "root"
	password	= "password"
	database	= "gymfinity"
)

var Connection *sql.DB

func Connect() {
	dataSource := fmt.Sprintf("%v:%v@/%v", user, password, database)
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	log.Println("[MySQL] Connection estabilished")
	Connection = db;
}