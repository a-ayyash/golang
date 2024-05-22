package main

import (
	"fmt"
	"math"
)

//receive function as parameter
func compute(fn func(float64, float64) float64) float64 {
  return fn(3,4)
}

//return closure
func addr() func(int) int {
  sum := 0
  return func(x int) int {
    sum = sum + x
    return sum
  }
}

func fibbonaci() func() int {
  sum  := -1
  sum2 := 0

  return func() int {
    if sum == -1 || sum == 0 {
      sum = sum + 1
      return sum
    }

    r := sum + sum2
    sum2 = sum
    sum = r 
    return r
  }
}

func main() {
  sqrt := func(x, y float64) float64 {
    return math.Sqrt(x*x + y*y)
  }

  pos, minus := addr(), addr()

  fmt.Println(compute(sqrt))
  fmt.Println(compute(math.Pow))

  for i := 0; i < 10; i++ {
    fmt.Println(pos(i))
    fmt.Println(minus(-2*i))
  }

  fmt.Println("----------------------")
  fib := fibbonaci()

  for i := 0; i < 10; i++ {
    fmt.Println("%d, ", fib())
  }
}
