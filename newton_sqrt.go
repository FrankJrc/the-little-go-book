package main

import "fmt"
import "math"

func guess(x float64) float64 {
    var pre,result float64 = 0.0, 1.0;
	for math.Abs(pre-result) > 0.0001 {
		pre = result
		result -= (result*result - x) / (2*result)
	}
	return result
}

func main()  {
	var original float64 = 90.0
	var result float64 = guess(original)
	var difference float64 = math.Abs(result - math.Sqrt(original))
	fmt.Println(difference)
	if difference < 0.0001 {
		fmt.Println("perfect")
	} else if difference < 0.01 {
		fmt.Println("not too bad")
	} else {
		fmt.Println("pls improve")
	}
}
