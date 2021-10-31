package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup //counter
	wg.Add(1)             //increment by 1

	go func() {
		count("sheep")
		wg.Done() //decrement the counter by 1
	}()

	wg.Wait() //block until count is 0
}

func count(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
