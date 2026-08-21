package main

import "fmt"

func main() {
	var nums[] int
	for i := 0; i < 20; i++ {
		nums = append(nums, i)
		fmt.Printf("Length: %d Capacity: %d\n", len(nums), cap(nums))
	}
}