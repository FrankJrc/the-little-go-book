package main

import "golang.org/x/tour/pic"
//import "fmt"

func Pic(dx, dy int) [][]uint8 {
	//fmt.Println(dx, dy)
	var res [][]uint8 = make([][]uint8, dx)
	for x := range res {
		res[x] = make([]uint8, dy)
		for i := range res[x] {
			res[x][i] = uint8((x*i))
		}
	}
	return res
}

func main() {
	pic.Show(Pic)
}
