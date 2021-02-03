package setting

import (
	"fmt"

	"github.com/go-ini/ini"
)

type Database struct {
	Type     string
	User     string
	Password string
	Host     string
	Port     string
	DbName   string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host string
	Port string
}

var cfg *ini.File

var RedisSetting = &Redis{}

func init() {
	var err error
	cfg, err = ini.Load("conf/app.conf")
	if err != nil {
		fmt.Println("setting.Setup failed")
	}

	mapTo("database", DatabaseSetting)
	mapTo("redis", RedisSetting)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		fmt.Println(err)
	}
}
