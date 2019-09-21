package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	res, prev:= 1, 0
	return func() int {
		res, prev = prev, res+prev
		return res
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 9; i++ {
		fmt.Println(f())
	}
}
