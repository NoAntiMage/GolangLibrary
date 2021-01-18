package main

import (
	"tmpgo/router"

	_ "github.com/go-sql-driver/mysql"
)

//var d = db.InitDB

//var c = db.RedisConn

var r = router.Router

func main() {
	r.Run()
	//var e models.Employee
	//l := e.GetEmpToDates(10009)
	//fmt.Println(l)
	//e.UpdateEmpLeaveDate()
	//l := e.QueryRangeEmps(0, 10)
	//fmt.Println(l)
	//e.UpdateAllLeaveDate()
}
