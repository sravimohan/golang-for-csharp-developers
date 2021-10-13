package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		Work("Task 1", 2)
		wg.Done()
	}()

	go func() {
		Work("Task 2", 1)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Done")
}

func Work(name string, sleep time.Duration) {
	time.Sleep(sleep * time.Second)
	fmt.Println(name, "completed")
}
