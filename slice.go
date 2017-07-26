// slice

package main

import "fmt"

func main() {
  a := [5]int{2, 10, 8, 15, 5}
  s := a[2:4] // [8, 15] not include 5
  s[1] = 12 // slice refers to the original array so also the value of the original orray changes in the same way as javascript
  // t := a[2:] // from the 3rd to till the end
  // u := a[:4] // from the top to the 4th
  fmt.Println(a)
  fmt.Println(s)
  fmt.Println(len(s))
  fmt.Println(cap(s)) // Returns the maximum ammount that the array can storage. from the value it started to refers to to the end of the original one
}
