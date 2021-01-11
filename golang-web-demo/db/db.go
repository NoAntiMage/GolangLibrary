package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var InitDB *sql.DB
var err error
var hello string

func init() {
	hello = "world"
	InitDB, err = sql.Open("mysql", "root:123456@tcp(192.168.133.48:3308)/employees_practice")
	if err != nil {
		fmt.Println(err)
	}
	err = InitDB.Ping()
	if err != nil {
		fmt.Println(err)
	}
}
