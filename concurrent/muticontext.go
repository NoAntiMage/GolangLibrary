package main

import (
	"context"
	"fmt"
	"time"
)

func watch(ctx context.Context, job string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(job, "stop")
			return
		default:
			fmt.Println(job, "running")
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "job1")
	go watch(ctx, "job2")
	go watch(ctx, "job3")

	time.Sleep(10 * time.Second)
	fmt.Println("all done")
	cancel()
	time.Sleep(3 * time.Second)
}