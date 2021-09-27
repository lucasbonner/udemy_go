package main

import "fmt"

func main() {
	m := map[string]int{
		"Lucas":  25,
		"Lauren": 22,
	}

	fmt.Println(m)

	//delete entry

	delete(m, "Lucas")
	fmt.Println(m)

	//if delete key that doesn't exist; no errors thrown
	delete(m, "Ian Fleming")
	fmt.Println(m)
}
