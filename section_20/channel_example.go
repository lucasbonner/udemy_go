package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go count("sheep", c)

	//send && receive blocking operations

	//code here will stop and wait to receive
	for { //for loop will "deadlock" unless you close channel in count function
		msg, open := <-c

		if !open {
			break
		}
		fmt.Println(msg)
	}
}

//channel -> pipe send||receive message
func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		//code here will stop and wait to receive
		c <- thing //send value of thing over the channel
		time.Sleep(time.Millisecond * 500)
	}

	close(c) //close channel
}
