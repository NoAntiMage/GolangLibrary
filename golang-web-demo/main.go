package main

import (
	"tmpgo/db"
	"tmpgo/router"

	_ "github.com/go-sql-driver/mysql"
)

var d = db.InitDB

var c = db.RedisConn

var r = router.Router

func main() {
	r.Run()
}
