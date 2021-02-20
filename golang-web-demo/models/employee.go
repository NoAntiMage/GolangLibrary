package models

import (
	"fmt"
	"strconv"
	"time"
	"tmpgo/db"
	"tmpgo/utils"
)

type Employee struct {
	EmpNo     int
	BirthDate time.Time
	FirstName string
	LastName  string
	Gender    string
	HireDate  time.Time
	LeaveDate time.Time
}

func (this Employee) GetEmpName(id int) (name string) {
	var first_name string
	var last_name string
	sql := "SELECT `first_name`,`last_name` FROM `employees` WHERE `emp_no` = ?"
	row := d.QueryRow(sql, id)
	row.Scan(&first_name, &last_name)
	name = fmt.Sprintf("%v %v", first_name, last_name)
	return
}

func (this Employee) GetEmpProfile(id int) (emp Employee) {
	var BirthDate string
	var HireDate string
	//	var LeaveDate string
	sql := "SELECT * FROM `employees` WHERE `emp_no` = ?"
	row := d.QueryRow(sql, id)
	fmt.Println(row)
	err := row.Scan(&emp.EmpNo, &BirthDate, &emp.FirstName, &emp.LastName, &emp.Gender, &HireDate) //	&LeaveDate

	if err != nil {
		panic(err)
	}
	//	fmt.Println(p.EmpNo)
	//	fmt.Println(BirthDate)
	emp.BirthDate = utils.ParseDate(BirthDate)
	emp.HireDate = utils.ParseDate(HireDate)
	//	this.LeaveDate = utils.ParseDate(LeaveDate)
	return
}

func (this Employee) CountEmp() (num int) {
	sql := "SELECT count(emp_no) FROM employees"
	row := d.QueryRow(sql)
	row.Scan(&num)
	return
}

func (this Employee) GetEmpToDates(id int) (list []string) {
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

func (this Employee) UpdateEmpLeaveDate(id int, leaveDate string) {
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

func (this Employee) UpdateRangeEmpsLeaveDates(empList []int) {
	fmt.Println(empList)
	for _, id := range empList {
		leaveDate := this.GetEmpLeaveDate(id)
		//fmt.Println(id)
		//fmt.Println(leaveDate)
		if leaveDate == "9999-12-31" {
			break
		}
		this.UpdateEmpLeaveDate(id, leaveDate)
	}
}

func (this Employee) QueryRangeEmps(offset int, pageSize int) (empNoList []int) {
	sql := "SELECT emp_no FROM employees LIMIT ?,?"
	stmt, err := d.Prepare(sql)
	if err != nil {
		panic(err)
	}
	rows, err := stmt.Query(offset, pageSize)
	if err != nil {
		panic(err)
	}
	fmt.Println("dealing with offset: " + strconv.Itoa(offset))

	for rows.Next() {
		var empNo int
		rows.Scan(&empNo)
		empNoList = append(empNoList, empNo)
	}
	//	fmt.Println(empNoList)
	return
}

func (this Employee) GetEmpLeaveDate(id int) (maxDate string) {
	l := this.GetEmpToDates(id)
	maxDate = utils.MaxDateInList(l)
	return
}

func (this Employee) UpdateAllEmpLeaveDate() {
	num := this.CountEmp()
	//fmt.Println(num)
	redisKey := "page_offset"
	pageOffset, err := strconv.Atoi(db.RedisGet(redisKey))
	if err != nil {
		panic(err)
	}
	pageSize := 500
	pageLeft := (num / pageSize) - pageOffset + 1

	fmt.Println("start at page: " + strconv.Itoa(pageOffset))

	for i := 0; i < pageLeft; i++ {
		targetEmpList := this.QueryRangeEmps(pageOffset*pageSize, pageSize)
		this.UpdateRangeEmpsLeaveDates(targetEmpList)
		pageOffset += 1
		db.RedisIncr(redisKey)
	}

	db.RedisReset(redisKey)
}
