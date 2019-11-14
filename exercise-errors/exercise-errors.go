package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) string {
	if x < 0 {
		return fmt.Sprint(ErrNegativeSqrt(x))
	}
	z := 1.0
	for i := 0; i < 10; i++ {
		old := z
		z -= (z*z - x) / (2 * z)
		if i > 1 && old < z {
			break
		}
	}
	return fmt.Sprintf("%f", z)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
