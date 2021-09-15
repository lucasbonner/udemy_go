package main

import "fmt"

var y = 42

//Declared "z" is of type string
var z = "Shaken, not stirred"

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	//z = 43 #> will not work, z is of type String, and golang is a STATIC LANGUAGE, not dynamic

	fmt.Println(z)
	fmt.Printf("%T\n", z)
}
