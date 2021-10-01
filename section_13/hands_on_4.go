/*
Create and use an anonymous struct
*/
package main

import "fmt"

func main() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "Lucas",
		last:  "Bonner",
		age:   25,
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	fmt.Println(p1.age)
}
