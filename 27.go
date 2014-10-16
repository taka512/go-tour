package main

import "fmt"

// Vertex is
type Vertex struct {
	X int
	Y int
}

func main() {
	p := Vertex{1, 2}
	q := &p
	q.X = 1e8
	fmt.Println(p)
}
