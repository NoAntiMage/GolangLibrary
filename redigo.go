package main

import (
	"fmt"
	// "log"

	"github.com/gomodule/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "192.168.133.48:10109")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	v, err := c.Do("SET", "name", "wujimaster")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(v)
	v, err = redis.String(c.Do("GET", "name"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
}

var Pool *redis.Pool

func redisPool() {
	Pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "192.168.133.48:10109")
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
}
