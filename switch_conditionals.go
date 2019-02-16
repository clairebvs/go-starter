package main

import (
  "fmt"
)

func main() {
  x := 3

  switch x {
  case 1:
    fmt.Println("one")
  case 2:
    fmt.Println("two")
  case 3:
    fmt.Println("three")
  default:
    fmt.Println("many")
  }

  switch {
  case x > 100:
    fmt.Println("Norway")
  default:
    fmt.Println("Brazil")
  }
}
