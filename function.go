// function

package main

import "fmt"

// func hi(name string) string {
//   fmt.Println("hi!" + name)
//   msg := "hi!" + name
//   return msg
// }

// or

func hi(name string) (msg string) {
  msg = "hi!" + name
  return
}

func main() {
  // hi("sasaki")
  fmt.Println(hi("sasaki"))
}
