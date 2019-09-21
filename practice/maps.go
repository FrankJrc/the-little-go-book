package main

import (
	"golang.org/x/tour/wc"
	//"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	//fmt.Println(s)
	parts := strings.Split(s, " ")
	res := make(map[string]int)
	for _, v := range parts {
		res[v] ++
	}
	return res
}

func main() {
	wc.Test(WordCount)
}
