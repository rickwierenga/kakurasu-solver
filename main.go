package main

import (
  "fmt"
)

func main() {
  b := NewBoard(5)
  b.Print()
  fmt.Println(b.Validate())
}
