/*
Create your own type "person" which will have an underlying type of "struct" so that it can store:
	first name
	last name
	favorite ice cream flavors
Create two values of TYPE person. Print out the values, ranging over the elmeents in the slice
*/
package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	favIceCream []string
}

func main() {
	p1 := person{
		firstName: "Lucas",
		lastName:  "Boner",
		favIceCream: []string{
			"chocolate",
			"vanilla",
			"cookie dough",
		},
	}

	p2 := person{
		firstName: "Lauren",
		lastName:  "Guthrie",
		favIceCream: []string{
			"strawberry",
			"vanilla",
			"raspberry",
		},
	}

	for i, v := range p1.favIceCream {
		fmt.Println(i, v)
	}

	for i, v := range p2.favIceCream {
		fmt.Println(i, v)
	}
}
