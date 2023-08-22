package main

//sqlx example according to the doc.
//reference: http://jmoiron.github.io/sqlx/

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB

func InitDB() (err error) {
	dsn := "./webgo.db"
	db, err = sqlx.Open("sqlite3", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxIdleConns(200)
	db.SetMaxIdleConns(30)

	return
}

func Close() {
	db.Close()
}

// CREATE TABLE "place" (
//   "country" text,
//   "city" text NULL,
//   "telcode" integer,
//   "time_zone" text
// );
type Place struct {
	Country       string
	City          sql.NullString
	TelephoneCode int `db:"telcode"`
	// TimeZone      string `db:"TimeZone"`
	TimeZone string `db:"time_zone"`
}

func WithExec() {
	sqlc := `
		SELECT sqlite_version() as version;
	`
	_, err := db.Exec(sqlc)
	if err != nil {
		fmt.Println(err)
	}
}

func WithMustExec() {
	sqlc := `INSERT INTO place (country, city, telcode) VALUES(?,?,?)`
	result := db.MustExec(sqlc, "China", "Wuhan", 27)
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("id : ", id)

	rowCount, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("row affected: ", rowCount)
}

// return row results.
func WithQuery() {
	rows, err := db.Query("SELECT country, city, telcode FROM place")
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var country string
		var city sql.NullString
		var telcode int
		err = rows.Scan(&country, &city, &telcode)
		fmt.Printf("%v %v %v\n", country, city, telcode)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}
}

func WithQueryx() {
	rows, err := db.Queryx("SELECT * FROM place")
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var p Place
		err = rows.StructScan(&p)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%+v\n", p)
	}
}

//fetch one row
func WithQueryRow() {
	row := db.QueryRow("SELECT * FROM place WHERE telcode = ? ", 8)
	var telcode int
	err := row.Scan(&telcode)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("telcode : %v", telcode)
}

func WithQueryRowx() {
	var p Place
	db.QueryRowx("SELECT city, telcode FROM place LIMIT 1").StructScan(&p)

}

func WithGet() {
	var p Place
	db.Get(&p, "SELECT * FROM place LIMIT 1")
	fmt.Printf("%+v\n", p)

}

func WithSelect() {
	pp := []Place{}
	db.Select(&pp, "select * FROM place WHERE telcode > ?", 3)
	fmt.Printf("%+v\n", pp)
}

func WithTx() {
	tx := db.MustBegin()
	_, err := tx.Exec("INSERT INTO place (country, city, telcode) VALUES(?,?,?)", "China", "Wuhan", 28)
	if err != nil {
		fmt.Println(err)
	}
	err = tx.Commit()
	// err = tx.Rollback()
	if err != nil {
		fmt.Println(err)
	}
}

func WithPreparex() {
	stmt, err := db.Preparex(`SELECT * FROM place WHERE telcode=?`)
	if err != nil {
		fmt.Println(err)
	}
	var p Place
	err = stmt.Get(&p, 27)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", p)
}

func WithInRange() {
	var tels = []int{9, 27, 28}
	query, args, err := sqlx.In("SELECT * FROM place WHERE telcode IN (?)", tels)
	if err != nil {
		fmt.Println(err)
	}
	query = db.Rebind(query)
	rows, err := db.Queryx(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	// WithSliceScan(rows)
	WithMapScan(rows)

}

// Named prefix with : mark
func WithNamedQuery() {
	p := Place{Country: "China"}
	// p := map[string]any{"country": "China"}
	rows, err := db.NamedQuery("SELECT * FROM place WHERE country=:country", p)
	if err != nil {
		fmt.Println(err)
	}
	WithMapScan(rows)

}

func WithNamedPrepare() {
	p := Place{Country: "China"}
	pp := []Place{}

	stmt, err := db.PrepareNamed(`SELECT * FROM place WHERE country = :country`)
	if err != nil {
		fmt.Println(err)
	}
	err = stmt.Select(&pp, p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", pp)

}

func WithSliceScan(rows *sqlx.Rows) {
	for rows.Next() {
		cols, err := rows.SliceScan()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%+v\n", cols)
	}
}

func WithMapScan(rows *sqlx.Rows) {
	for rows.Next() {
		res := make(map[string]any)
		err := rows.MapScan(res)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
	}
}

func main() {
	InitDB()
	defer Close()

	// do some biz here..

}
