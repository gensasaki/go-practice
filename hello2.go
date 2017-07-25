// constant

package main

import "fmt"

func main() {
  // const name = "sasaki"
  // name = "tanaka" // cannot assign to name
  // fmt.Println()

  const (
    sun = iota // 0
    mon // 1
    tue // 2
  )

  fmt.Println(sun, mon, tue)
}
