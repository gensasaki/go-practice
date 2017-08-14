// structure
// method (function based on data type)

package main

import "fmt"

type user struct {
  name string
  score int
}

func (u user) show() {
  fmt.Printf("name:%s, score:%d\n", u.name, u.score)
}

func (u user) hit() {
  u.score++
}

func main() {
  // u := new(user)
  // (*u).name = "sasaki"
  // u.name = "sasaki"
  // u.score = 20

  u := user{name:"sasaki", score:20}
  u.show()
  fmt.Println(u)
}
