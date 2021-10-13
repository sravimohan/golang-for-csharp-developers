package main

import (
	"fmt"
	"net/http"
	"sync"
)

// "Do not communicate by sharing memory; instead, share memory by communicating."

func main() {
	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	ch := make(chan string)
	defer close(ch)

	var wg sync.WaitGroup
	wg.Add(len(apis))

	for _, api := range apis {
		apiToCheck := api
		go func() {
			checkApi(apiToCheck, ch)
			wg.Done()
		}()
	}

	for i := 0; i < len(apis)*2; i++ {
		fmt.Println(<-ch)
	}

	wg.Wait()
}

func checkApi(api string, ch chan string) {
	ch <- fmt.Sprintf("Checking:%s", api)
	_, err := http.Get(api)
	if err == nil {
		ch <- fmt.Sprintf("OK:%s", api)
	} else {
		ch <- fmt.Sprintf("BAD:%s", api)
	}
}
