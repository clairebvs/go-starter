package main

import (
  "fmt"
)

func main() {
  var x float64
  var y float64

  x = 4
  y = 2

  fmt.Printf("x=%v, type of %T\n", x, x)
  fmt.Printf("y=%v, type of %T\n", y, y)

  var mean float64
  mean = (x + y)/2
  fmt.Printf("The result is %v, type of %T\n", mean, mean)
}
