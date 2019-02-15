package main

import (
  "fmt"
)

func main() {
  x := 100

  if x > 50 {
    fmt.Println("x is bigger than you think")
  }

  if x > 50 {
    fmt.Println("Japan")
  } else {
    fmt.Println("Austria")
  }

  if x == 50 {
    fmt.Println("Japan")
  } else if x > 150 {
    fmt.Println("Austria")
  } else {
    fmt.Println("Peru")
  }
}
