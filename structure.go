// structure

package main

import "fmt"

type user struct {
  name string
  score int
}

func main() {
  // u := new(user)
  // (*u).name = "sasaki"
  // u.name = "sasaki"
  // u.score = 20

  u := user{name:"sasaki", score:20}
  fmt.Println(u)
}
