//create a program that uses switch statement with the switch expression specified as
// variable of TYPE string with IDENTIFIER "favsport"

package main

import "fmt"

func main() {
	favSport := "skateboarding"
	switch favSport {
	case "baseball":
		fmt.Println("yeah i like BASEBALL")
	case "football":
		fmt.Println("yeah I like FOOTBALL")
	case "soccer":
		fmt.Println("yeah i like futbol")
	case "basketball":
		fmt.Println("yeah i HOOP on em")
	case "skateboarding":
		fmt.Println("doOooo a KICKFLIP")
	}
}
