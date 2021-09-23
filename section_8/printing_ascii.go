//loop over 33-122, print them out as text
package main

import "fmt"

func main() {
	a := 33

	for a < 123 {
		fmt.Printf("%c\n", a)
		a++
	}
}
