/*
Take code from previous exercie
store values of type person in a map
with the key of last name

access each value in the map
print out the values, ranging over the slice
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

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v.firstName)
		fmt.Println(v.lastName)
		for i, val := range v.favIceCream {
			fmt.Println(i, val)
		}
		fmt.Println("------------------")
	}
}
