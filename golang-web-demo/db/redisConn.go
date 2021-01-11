package db

import (

	// "log"

	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

var RedisConn *redis.Pool

func init() {
	Setup()
	rc := RedisConn.Get()
	defer rc.Close()
	v, _ := rc.Do("PING")
	fmt.Println(v)
}

func Setup() error {
	RedisConn = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "192.168.133.48:10109")
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
