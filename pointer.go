// pointer
// variable that indicates address
// cannot calculate

package main

import "fmt"

func main() {
  a := 5
  var pa *int
  pa = &a // &a = address of a
  // value of data that is in pa = *pa
  fmt.Println(pa)
  fmt.Println(*pa)
}
