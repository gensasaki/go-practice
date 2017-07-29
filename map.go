// map

package main

import "fmt"

func main() {
  // m := make(map[string]int)
  // m["sasaki"] = 200
  // m["gen"] = 300

  m := map[string]int{"sasaki": 100,"gen": 200}
  fmt.Println(m)
  fmt.Println(len(m))
  delete(m, "sasaki")
  fmt.Println(m)
  v, ok := m["gen"]
  fmt.Println(v)
  fmt.Println(ok)

  n := map[int]bool{1: true, 2: false}
  fmt.Println(n)
}
