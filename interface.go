// interface

package main

import "fmt"

func show(t interface{}) {
  // type assertion
  // _, ok := t.(japanese)
  // if ok {
  //   fmt.Println("I am Japanese")
  // } else {
  //   fmt.Println("I am not Japanese")
  // }

  // type switch
  switch t.(type) {
  case japanese:
    fmt.Println("I am Japanese")
  default:
    fmt.Println("I am not Japanese")
  }
}

type greeter interface {
  greet()
}
type japanese struct {}
type american struct {}

func (j japanese) greet() {
  fmt.Println("Konnnichiwa!")
}

func (a american) greet() {
  fmt.Println("Hello!")
}

func main() {
  greeters := []greeter{japanese{}, american{}}
  for _, greeter := range greeters {
    greeter.greet()
    show(greeter)
  }
}
