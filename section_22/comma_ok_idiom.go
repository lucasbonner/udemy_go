package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 42
		close(c)
	}()

	v, ok := <-c //comma ok idiom

	fmt.Println(v, ok)
}
