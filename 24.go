package main

import (
	"fmt"
	"math"
)

// Sqrt is
func Sqrt(x float64) float64 {
	var z float64 = 1
	for i := 0; i < 100; i++ {
		calcZ := z - (((z * z) - x) / (2 * z))
		if math.Abs(z-calcZ) < 0.0001 {
			fmt.Println(i)
			return calcZ
		}
		z = calcZ

	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
