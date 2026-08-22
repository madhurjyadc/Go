package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value: ", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value: ", m["Answer"])

	// delete(m, "Answer")
	// fmt.Println("The value: ", m["Answer"])

	// v,ok := m["Answer2"]
	// fmt.Println("The value: ", v, "Present? ", ok)

	for key, value := range m {
		fmt.Println(key, value)
	}
	
}