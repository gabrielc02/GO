package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func pingSite(url string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	ch <- fmt.Sprintf("Site [%s] respondeu em [%v]", url, time.Since(start))
}

func main() {
	var wg sync.WaitGroup
	Chan := make(chan string)
	url := []string{"www.url1.com", "www.url2.com", "www.url3.com"}

	for _, ping := range url {
		wg.Add(1)
		go pingSite(ping, Chan, &wg)
	}

	go func() {
		wg.Wait()
		close(Chan)
	}()

	for res := range Chan {
		fmt.Println(res)
	}
}
