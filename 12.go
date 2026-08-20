package main

import "fmt"

func main() {
	sum := 1

	for {
		if (sum > 100) {
			break
		}

		sum += sum
	}

	fmt.Println(sum)
}