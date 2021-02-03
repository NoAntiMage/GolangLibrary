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
//	var (
//		dbType, dbUser, dbPassword, dbHost, dbPort, dbName string
//	)
//	dbType = setting.DatabaseSetting.Type
//	dbUser = setting.DatabaseSetting.User
//	dbPassword = setting.DatabaseSetting.Password
//	dbHost = setting.DatabaseSetting.Host
//	dbPort = setting.DatabaseSetting.Port
//	dbName = setting.DatabaseSetting.DbName
//
//	fmt.Println(dbType)
//	//	InitDB, err = sql.Open("mysql", "root:123456@tcp(192.168.133.48:3307)/employees_practice")
//	fmt.Printf("%s:%s@tcp(%s:%s)/%s",
//		dbUser,
//		dbPassword,
//		dbHost,
//		dbPort,
//		dbName)
//}
