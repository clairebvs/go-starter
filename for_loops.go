package main

import (
  "fmt"
)

func main() {
  for i := 0; i < 3; i++ {
    fmt.Println(i)
  }

  for i := 0; i < 5; i++ {
    if i < 2 {
      continue
    }
    fmt.Println(i)
  }

  for i := 0; i < 5; i++ {
    if i > 3 {
      break
    }
    fmt.Println(i)
  }

  // manual incrementation
  // starting point
  a := 0
  for a < 5 {
    fmt.Println(a)
    a++ 
  }
}
