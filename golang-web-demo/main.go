package main

import (
	"tmpgo/models"

	_ "github.com/go-sql-driver/mysql"
)

//var d = db.InitDB

//var c = db.RedisConn

//var r = router.Router

func main() {
	//	r.Run()
	var e models.Employee
	e.UpdateAllEmpLeaveDate()
}
