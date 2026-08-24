package main

import "fmt"

func main() {
	var i interface{} = "hello"

	// Here we are type asserting - converting an interface value to a string one.
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f,ok)

	f = i.(float64)
	fmt.Println(f)
}