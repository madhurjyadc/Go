package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex {
	"Bell Labs": Vertex{
		40.235324, -23.23432,
	},

	"Google": Vertex {
		63.23423, 24.5123,
	},
}

func main() {
	fmt.Println(m)
}