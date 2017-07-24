// comment

/*
comment
comment
*/

// when first character of variable name is a lower case, it is visible only in that package
// when first character of variable name is a capital case, is is vilible in everywhere

// basic data types
// string "hello"
// int 53
// float64 10.2
// bool true false
// nil
//
// var s string // ""
// var a int // 0
// var f bool // false


package main

import "fmt"

func main() {
  // var msg string
  // msg = "Hello, world"
  // var msg = "Hello, world"
  // msg := "Hello, world"
  // fmt.Println(msg)


  a := 10
  b := 12.3
  c := "hoge"
  var d bool

  fmt.Printf("a: %d, b: %f, c: %s, d: %t\n", a, b, c, d)

  // var a, b int
  // a, b = 10, 15
  // a, b := 10, 15
  //
  // var (
  //   c int
  //   d string
  // )
  // c = 20
  // d = "hoge"
}
