
package main

import "fmt"
import "math"

type Point struct {
    X, Y float64
}

func (p Point) Display() {
    fmt.Println(p.X, p.Y)
}

func (p1 Point) Distance(p2 Point) float64 {
    return math.Sqrt((p1.X-p2.X)*(p1.X-p2.X) + (p1.Y-p2.Y)*(p1.Y-p2.Y))
}

func main() {
    x, y:= Point{X:4, Y:5}, Point{X:7, Y:1}
    x.Display()
    y.Display()
    fmt.Println(x.Distance(y))
}
