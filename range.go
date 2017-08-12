
package main
import "fmt"

func main() {
  // s := []int{2, 3, 4}
  // for i, v := range s {
  //   fmt.Println(i, v)
  // }

  // for _, v := range s {
  //   fmt.Println(v)
  // }

  m := map[string]int{"sasaki": 200, "gen": 300}
  for k, v := range m {
    fmt.Println(k, v)
  }
}
