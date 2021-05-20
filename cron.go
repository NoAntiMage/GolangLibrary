package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()
	c.AddFunc("@every 1s", func() {
		fmt.Println("tick tuck every 1 sec")
	})

	c.Start()
	time.Sleep(time.Second * 10)
}

// reference: https://github.com/robfig/cron/blob/master/doc.go