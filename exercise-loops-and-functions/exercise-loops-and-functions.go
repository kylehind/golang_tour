package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		old := z
		z -= (z*z - x) / (2 * z)
		if i > 1 && old < z {
			break
		}
	}
	return z
}

func main() {
	fmt.Println("Sqrt")
	fmt.Println(Sqrt(3))

	fmt.Println("math.Sqrt")
	fmt.Println(math.Sqrt(3))
}
