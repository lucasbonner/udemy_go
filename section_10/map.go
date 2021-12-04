package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)

	fmt.Println(m["James"])

	fmt.Println(m["Barnabas"]) //0
	//reference key not there, return 0

	//comma ok idiom
	v, ok := m["Barnabas"]

	fmt.Println("v", v)
	fmt.Println("ok", ok)

	if v, ok := m["James"]; ok {
		fmt.Println("THIS IS THE IF PRINT", v)
	}
}
