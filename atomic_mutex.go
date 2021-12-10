package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func MutexRun() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	count := int64(0)
	t := time.Now()
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func(i int) {
			mutex.Lock()
			count++
			wg.Done()
			mutex.Unlock()
		}(i)
	}

	wg.Wait()

	fmt.Printf("mutex run cost time: %d, count is: %d\n", time.Now().Sub(t), count)
}

func AtomicRun() {
	var wg sync.WaitGroup
	count := int64(0)
	t := time.Now()
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func(i int) {
			atomic.AddInt64(&count, 1)
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Printf("atomic run cost time: %d, count is :%d\n", time.Now().Sub(t), count)
}

func main() {
	MutexRun()
	AtomicRun()
}

/*
reference:
https://blog.betacat.io/post/golang-atomic-value-exploration/
https://blog.haohtml.com/archives/25881
https://www.intel.com/content/dam/www/public/us/en/documents/manuals/64-ia-32-architectures-software-developer-instruction-set-reference-manual-325383.pdf
https://en.wikipedia.org/wiki/Compare-and-swap
*/
