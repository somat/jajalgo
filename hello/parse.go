package main

import "strconv"
import "fmt"

func main() {
  f, _ := strconv.ParseFloat("1.234", 64)
  fmt.Println(f)
}
