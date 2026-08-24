package main

import "fmt"

func IndexInt (s []int, x int) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func IndexString (s []string, x string) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	si := []int{10, 20, 15, -10}
	ss := []string{"hello", "world", "type", "bear"}

	fmt.Println(IndexInt(si, 10))
	fmt.Println(IndexString(ss, "world"))
}