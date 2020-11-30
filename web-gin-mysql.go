package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

type Person struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
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

func (this Person) get() (person Person, err error) {
	row := db.QueryRow("SELECT id, first_name, last_name FROM user WHERE id=?", this.Id)
	err = row.Scan(&person.Id, &person.FirstName, &person.LastName)
	if err != nil {
		return
	}
	return
}

func (this Person) add() (Id int, err error) {
	fmt.Println(this.FirstName, this.LastName)
	stmt, err := db.Prepare("INSERT INTO user(first_name, last_name) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	rs, err := stmt.Exec(this.FirstName, this.LastName)
	if err != nil {
		log.Fatal(err)
	}

	id, err := rs.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}
	Id = int(id)
	defer stmt.Close()
	return
}

func (this Person) del() (rows int, err error) {
	stmt, err := db.Prepare("DELETE FROM person WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	rs, err := stmt.Exec(this.Id)
	if err != nil {
		log.Fatalln(err)
	}

	row, err := rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()
	rows = int(row)
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

	router.GET("/person/:id", func(c *gin.Context) {
		var result gin.H
		id := c.Param("id")
		Id, err := strconv.Atoi(id)
		if err != nil {
			log.Fatal(err)
		}
		p := Person{
			Id: Id,
		}
		person, err := p.get()
		if err != nil {
			result = gin.H{
				"result": nil,
				"count":  0,
			}
		} else {
			result = gin.H{
				"result": person,
				"count":  1,
			}
		}
		c.JSON(http.StatusOK, result)
	})

	router.POST("/person", func(c *gin.Context) {
		var p Person
		err := c.BindJSON(&p)
		if err != nil {
			log.Fatal(err)
		}
		Id, err := p.add()
		fmt.Println("id= ", Id)
		name := p.FirstName + " " + p.LastName
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%s insert successful", name),
		})
	})

	router.DELETE("/person/:id", func(c *gin.Context) {
		id := c.Param("id")
		Id, err := strconv.ParseInt(id, 10, 32)
		if err != nil {
			log.Fatalln(err)
		}
		p := Person{Id: int(Id)}
		rows, err := p.del()
		if err != nil {
			log.Fatalln(err)

		}
		fmt.Println("delete rows: ", rows)
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprint("delete user %s successfully", id),
		})
	})

	router.Run()
}
