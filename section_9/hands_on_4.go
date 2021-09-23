///Create a for loop, years alive, no condition at begining

package main

import "fmt"

func main() {
	x := 1996

	for {
		fmt.Println(x)
		x++
		if x > 2021 {
			break
		}
	}
}
