package main

import (
	"github.com/user/golang_tour/exercise-slices/pic"
)

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for x := range pic {
		pic[x] = make([]uint8, dx)
		for y := range pic[x] {
			z := uint8(x ^ y)
			pic[x][y] = z
		}
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
