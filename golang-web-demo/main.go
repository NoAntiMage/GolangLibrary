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

//func main() {
//	d := setting.DatabaseSetting
//	fmt.Println(d)
//	r := setting.RedisSetting
//	fmt.Println(r)
//}
