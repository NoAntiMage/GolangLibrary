package main

import (
	"database/sql"
	"fmt"

	//	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.133.48:3308)/gin")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}

	var version string
	err2 := db.QueryRow("SELECT VERSION()").Scan(&version)

	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println(version)
}
