/*
get this code working
*/
package main

import (
	"fmt"
)

func main() {
	// c := make(chan int, 1)//buffered channel
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
