// slice2

package main

import "fmt"

func main() {
  // s := make([]int, 3) // [0 0 0]
  s := []int{1, 3, 5}
  // append
  s = append(s, 8, 2, 10)
  // copy
  t := make([]int, len(s))
  n := copy(t, s) // returns the number of something copied
  fmt.Println(s)
  fmt.Println(t)
  fmt.Println(n)
}
