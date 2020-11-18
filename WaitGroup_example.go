package main

// WaitGroup

import (
	"fmt"
	"sync"
	"time"
)

func useSleep() {
	for i := 0; i < 100; i++ {
		go fmt.Println(i)
	}
	time.Sleep(time.Second)
}

func useChannel() {
	c := make(chan bool, 100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			c <- true
		}(i)
	}
	for i := 0; i < 100; i++ {
		<-c
	}
}

func useWaitGroup() {
	wg := sync.WaitGroup()
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func passArg(i int, wg *sync.WaitGroup) {
	fmt.Println()
	wg.Done()
}

func main() {
	wg := sync.WaitGroup()
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go passArg(i, &wg)
	}
	wg.Wait()
}
