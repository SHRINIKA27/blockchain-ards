package main

import (

  "fmt"

)

func add_values(a int, b int) int {

  var c = a + b

  return c
}

func main() {

  var a, b, c int

  a=1
  b=2

  c = add_values(a , b)

  fmt.Printf("Sum: %d \n", c)

}
