package views

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetVersion(c *gin.Context) {
	var version string
	sql := "SELECT VERSION()"
	d.QueryRow(sql).Scan(&version)

	fmt.Println(version)
	c.JSON(http.StatusOK, gin.H{
		"db_version": version,
	})
}

func GetNow(c *gin.Context) {
	const shortForm = "2006-01-02 15:04:05"
	var nowtime string
	sql := "SELECT NOW()"
	row := d.QueryRow(sql)
	row.Scan(&nowtime)
	fmt.Println(nowtime)
	//result, _ = time.Parse(shortForm, nowtime)
	//fmt.Println(result.Format("2006-01-02"))
	c.JSON(http.StatusOK, gin.H{
		"now": nowtime,
	})
	//return
}

func GetDatabases(c *gin.Context) {
	var databases []string
	rows, err := d.Query("SHOW DATABASES")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var database string
		rows.Scan(&database)
		//fmt.Println(database)
		databases = append(databases, database)
	}
	c.JSON(http.StatusOK, gin.H{
		"databases": databases,
	})
}

func GetTables(c *gin.Context) {
	var tables []string
	rows, err := d.Query("SHOW TABLES")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var table string
		rows.Scan(&table)
		//fmt.Println(table)
		tables = append(tables, table)
	}
	c.JSON(http.StatusOK, gin.H{
		"tables": tables,
	})
}
