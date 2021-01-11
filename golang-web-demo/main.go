package main

import (
	"fmt"
	"strconv"
	"time"
	"tmpgo/db"
	"tmpgo/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
)

var d = db.InitDB
var c = db.RedisConn

func main() {
	//GetVersion()
	//getDatabases()
	//	getEmpName(10001)
	//	getEmpProfile(10001)
	//	a := getNow()
	//	fmt.Println(a)
	//	getEmpProfile(10001)
	updateAllLeaveDate()
	//	redisPing()
	//redisGet("page_offset")
	//redisIncr("page_view")
	//redisReset("page_view")
}

func getVersion() {
	var version string
	sql := "SELECT VERSION()"
	d.QueryRow(sql).Scan(&version)
	fmt.Println(version)
}

func getNow() (result time.Time) {
	//var nowtime string
	const shortForm = "2006-01-02 15:04:05"
	var nowtime string
	sql := "SELECT NOW()"
	row := d.QueryRow(sql)
	row.Scan(&nowtime)
	fmt.Println(nowtime)
	result, _ = time.Parse(shortForm, nowtime)
	fmt.Println(result.Format("2006-01-02"))
	return
}

func parseDateTime(thetime string) (result time.Time) {
	const dateTimeForm = "2006-01-02 15:04:05"
	result, _ = time.Parse(dateTimeForm, thetime)
	return
}

func parseDate(thetime string) (result time.Time) {
	const dateForm = "2006-01-02"
	result, _ = time.Parse(dateForm, thetime)
	return
}

func getDatabases() {
	//	var databases string[]
	rows, err := d.Query("SHOW DATABASES")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var database string
		rows.Scan(&database)
		fmt.Println(database)
	}
}

func getTables() {
	rows, err := d.Query("SHOW TABLES")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var table string
		rows.Scan(&table)
		fmt.Println(table)
	}
}

func getEmpName(id int) {
	var first_name string
	sql := "SELECT `first_name` FROM `employees` WHERE `emp_no` = ?"
	row := d.QueryRow(sql, id)
	row.Scan(&first_name)
	fmt.Println(first_name)

}

func getEmpProfile(id int) {
	var p models.Employee
	var BirthDate string
	var HireDate string
	var LeaveDate string
	sql := "SELECT * FROM `employees` WHERE `emp_no` = ?"
	row := d.QueryRow(sql, id)
	fmt.Println(row)
	err := row.Scan(&p.EmpNo, &BirthDate, &p.FistName, &p.LastName, &p.Gender, &HireDate, &LeaveDate)
	if err != nil {
		panic(err)
	}
	//	fmt.Println(p.EmpNo)
	//	fmt.Println(BirthDate)
	p.BirthDate = parseDate(BirthDate)
	p.HireDate = parseDate(HireDate)
	p.LeaveDate = parseDate(LeaveDate)
	fmt.Println(p)
}

func updateAllLeaveDate() {
	num := countEmp()
	const pageSize = 500
	pageOffset, err := strconv.Atoi(redisGet("page_offset"))
	pageLeft := (num / pageSize) - pageOffset + 1
	//page := 3

	sql := "SELECT emp_no FROM employees LIMIT ?,?"

	//offset := 0
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("start at: " + strconv.Itoa(pageOffset))
	for i := 0; i < pageLeft; i++ {
		stmt, err := d.Prepare(sql)
		if err != nil {
			panic(err)
		}

		rows, err := stmt.Query(pageOffset*pageSize, pageSize)
		if err != nil {
			panic(err)
		}
		fmt.Printf("dealing with offset: " + strconv.Itoa(pageOffset))

		var empNoList []int
		for rows.Next() {
			var empNo int
			rows.Scan(&empNo)
			empNoList = append(empNoList, empNo)
		}
		fmt.Println(empNoList)
		for _, empNo := range empNoList {
			//fmt.Println(empNo)
			l := getEmpToDate(empNo)
			maxDate := maxDateInList(l)
			updateEmpLeaveDate(empNo, maxDate)
		}
		stmt.Close()
		pageOffset += 1
		redisIncr("page_offset")
	}

}

func updateEmpLeaveDate(id int, leaveDate string) {
	sql := "UPDATE employees SET leave_date = ? WHERE emp_no=?"
	//	stmt, err := d.Prepare(sql)
	//	if err != nil {
	//		panic(err)
	//	}
	//	_, err = stmt.Exec(leaveDate, id)
	rs, err := d.Exec(sql, leaveDate, id)
	if err != nil {
		panic(err)
	}
	RowsAff, _ := rs.RowsAffected()
	if RowsAff != 0 {
		fmt.Println("update id:" + strconv.Itoa(id))
		fmt.Println(rs.RowsAffected())
	}
}

func getEmpToDate(id int) (list []string) {
	sql := "SELECT to_date FROM dept_emp WHERE emp_no=?"
	rows, err := d.Query(sql, id)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var toDate string
		rows.Scan(&toDate)
		list = append(list, toDate)
	}
	return
}

func maxDateInList(list []string) (maxDate string) {
	tmp := parseDate("1900-01-01")
	for _, thedate := range list {
		result := parseDate(thedate)
		if result.After(tmp) {
			tmp = result
		}
	}
	maxDate = tmp.Format("2006-01-02")
	return
}

func countEmp() (num int) {
	sql := "SELECT count(emp_no) FROM employees"
	row := d.QueryRow(sql)
	row.Scan(&num)
	return

}

func redisPing() {
	rc := db.RedisConn.Get()
	rc.Do("PING")
}

func redisGet(value string) (v string) {
	rc := db.RedisConn.Get()
	defer rc.Close()
	v, err := redis.String(rc.Do("GET", value))
	if err != nil {
		//panic(err)
		redisReset(value)
		v = "0"
	}
	return
}

func redisIncr(value string) (v int) {
	rc := db.RedisConn.Get()
	defer rc.Close()
	v, err := redis.Int(rc.Do("INCR", value))
	if err != nil {
		panic(err)
	}
	return

}

func redisReset(value string) {
	rc := db.RedisConn.Get()
	defer rc.Close()
	v, err := rc.Do("SET", value, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(v)
}
