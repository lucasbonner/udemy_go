/*
in addition to the main goroutine, launch two additional goroutines
each additional goroutine should print something
use waitgroups to make sure each goroutine finishes b4 program exits
*/

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go bar()

	wg.Wait()
}

func foo() {
	fmt.Println("FOOFOO")
	wg.Done()
}

func bar() {
	fmt.Println("BARBAR")
	wg.Done()
}
