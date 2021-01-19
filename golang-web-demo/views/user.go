package views

import (
	"fmt"
	"net/http"
	"strconv"
	"tmpgo/models"

	"github.com/gin-gonic/gin"
)

func GetEmployeeName(c *gin.Context) {
	id := c.Query("id")
	Id, _ := strconv.Atoi(id)
	var empInstance models.Employee
	name := empInstance.GetEmpName(Id)
	fmt.Println(name)
	c.JSON(http.StatusOK, gin.H{
		"name": name,
	})
}

func GetEmployeeDetail(c *gin.Context) {
	id := c.Query("id")
	Id, _ := strconv.Atoi(id)
	var empInstance models.Employee
	emp := empInstance.GetEmpProfile(Id)
	//fmt.Println(emp)
	c.JSON(http.StatusOK, gin.H{
		"result": emp,
	})
}

func GetEmployeeSum(c *gin.Context) {
	var e models.Employee
	num := e.CountEmp()
	c.JSON(http.StatusOK, gin.H{
		"count": num,
	})
}

func GetRangeEmps(c *gin.Context) {
	offset := c.DefaultQuery("offset", "0")
	pagesize := c.DefaultQuery("pagesize", "10")
	Offset, _ := strconv.Atoi(offset)
	PageSize, _ := strconv.Atoi(pagesize)
	var e models.Employee
	l := e.QueryRangeEmps(Offset, PageSize)
	c.JSON(http.StatusOK, gin.H{
		"emp_list": l,
	})
}

//func UpdateToDate(c *gin.Context) {
//	jsonf := make(map[string]interface{})
//	c.BindJSON(&jsonf)
//	id, ok := jsonf["id"]
//	if ok {
//		var e models.Employee
//		leaveDate := e.GetEmpLeaveDate(id)
//		e.UpdateEmpLeaveDate(id, leaveDate)
//	} else {
//		fmt.Println("key <id> not exist.")
//	}
//	c.JSON(http.StatusOK, gin.H{
//		"message": "ok",
//	})
//}
