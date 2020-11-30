package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

type Person struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:last_name`
}

func (this Person) getAll() (persons []Person, err error) {
	rows, err := db.Query("SELECT id,first_name, last_name FROM user")
	if err != nil {
		return
	}
	for rows.Next() {
		var person Person
		rows.Scan(&person.Id, &person.FirstName, &person.LastName)
		persons = append(persons, person)
	}
	defer rows.Close()
	return
}

func main() {
	var err error
	db, err = sql.Open("mysql", "root:123456@tcp(192.168.133.48:3308)/gin")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	router := gin.Default()
	router.GET("/person", func(c *gin.Context) {
		p := Person{}
		persons, err := p.getAll()
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"result": persons,
			"count":  len(persons),
		})
	})
	router.Run()
}
