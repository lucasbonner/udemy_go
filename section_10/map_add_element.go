package main

import "fmt"

func main() {
	m := map[string]int{
		"Lucas":  25,
		"Lauren": 22,
	}

	//add element
	m["todd"] = 33

	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}
}
