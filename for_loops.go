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
}
