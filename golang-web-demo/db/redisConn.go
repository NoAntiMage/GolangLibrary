package db

import (

	// "log"

	"fmt"
	"time"
	"tmpgo/utils/setting"

	"github.com/gomodule/redigo/redis"
)

var RedisConn *redis.Pool

func init() {
	Setup()
	rc := RedisConn.Get()
	defer rc.Close()
	//	v, _ := rc.Do("PING")
	//	fmt.Println(v)
}

func RedisPing() {
	rc := RedisConn.Get()
	v, _ := rc.Do("PING")
	fmt.Println(v)
}

func RedisGet(value string) (v string) {
	rc := RedisConn.Get()
	defer rc.Close()
	v, err := redis.String(rc.Do("GET", value))
	if err != nil {
		//panic(err)
		RedisReset(value)
		v = "0"
	}
	return
}

func RedisIncr(value string) (v int) {
	rc := RedisConn.Get()
	defer rc.Close()
	v, err := redis.Int(rc.Do("INCR", value))
	if err != nil {
		panic(err)
	}
	return

}

func RedisReset(value string) {
	rc := RedisConn.Get()
	defer rc.Close()
	v, err := rc.Do("SET", value, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(v)
}

func Setup() error {
	var (
		redisHost, redisPort string
	)
	redisHost = setting.RedisSetting.Host
	redisPort = setting.RedisSetting.Port
	RedisConn = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			//c, err := redis.Dial("tcp", "192.168.133.48:10109")
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%s",
				redisHost,
				redisPort,
			))
			if err != nil {
				return nil, err
			}

			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return nil
}
