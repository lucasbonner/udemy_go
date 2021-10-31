package main

import "fmt"

func main() {
	c := make(chan string, 2) //2 is capacity
	c <- "hello"
	c <- "world"

	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)
}
