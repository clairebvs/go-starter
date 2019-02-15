package main

import (
  "fmt"
)

func main() {
  x, y := 1.0, 2.0

  fmt.Printf("x=%v, type of %T\n", x, x)
  fmt.Printf("y=%v, type of %T\n", y, y)

  mean := (x + y) / 2.0
  fmt.Printf("The result is %v, type of %T\n", mean, mean)
}
