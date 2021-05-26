package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("stop")
				return

			default:
				fmt.Println("running")
				time.Sleep(2 * time.Second)
			}

		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("all done")
	stop <- true

	time.Sleep(3 * time.Second)
}