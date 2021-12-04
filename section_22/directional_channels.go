package main

import "fmt"

func main() {
	c := make(chan<- int, 2)  //send only channel /only can SEND unto the channel
	c2 := make(<-chan int, 2) //receive only channel

	c <- 42
	c <- 43

	//fmt.Println(<-c) //these wont work with send only
	//fmt.Println(<-c) //these wont work with send only
	fmt.Println("-----")
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", c2)
}
