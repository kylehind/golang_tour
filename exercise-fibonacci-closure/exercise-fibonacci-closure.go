package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34
func fibonacci() func() int {
	prev1 := 0
	prev2 := 1
	return func() int {
		x := prev1
		prev1 = prev2
		prev2 = x + prev1
		return x
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println("NUM OUT", f())
	}
}
