package db

import (
	"database/sql"
	"fmt"
	"tmpgo/utils/setting"

	_ "github.com/go-sql-driver/mysql"
)

var InitDB *sql.DB
var err error
var hello string

func init() {
	var (
		dbType, dbUser, dbPassword, dbHost, dbPort, dbName string
	)
	dbType = setting.DatabaseSetting.Type
	dbUser = setting.DatabaseSetting.User
	dbPassword = setting.DatabaseSetting.Password
	dbHost = setting.DatabaseSetting.Host
	dbPort = setting.DatabaseSetting.Port
	dbName = setting.DatabaseSetting.DbName

	//	InitDB, err = sql.Open("mysql", "root:123456@tcp(192.168.133.48:3307)/employees_practice")
	InitDB, err = sql.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName))
	if err != nil {
		fmt.Println(err)
	}
	err = InitDB.Ping()
	if err != nil {
		fmt.Println(err)
	}
}
