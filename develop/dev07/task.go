package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	channels := []<-chan time.Time{
		time.After(1 * time.Second),
		time.After(2 * time.Second),
		time.After(3 * time.Second),
		time.After(4 * time.Second),
		time.After(5 * time.Second),
	}

	for ch := range or(channels...) {
		fmt.Println(ch)
	}
}

func or(channels ...<-chan time.Time) <-chan time.Time {
	var wg sync.WaitGroup
	out := make(chan time.Time)

	for _, c := range channels {
		wg.Add(1)
		go func(c <-chan time.Time) {
			defer wg.Done()
			out <- <-c
		}(c)
	}

	go func() {
		defer close(out)
		wg.Wait()
	}()

	return out
}
