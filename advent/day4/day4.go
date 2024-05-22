package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
)

func main() {
  fmt.Println("hello")
  text := "iwrupvqb"
  hasher := md5.New()
  io.WriteString(hasher, text)
  fmt.Println(fmt.Sprintf("%02x", hasher.Sum(nil)))

  found := -1

  for i := 1; i < 10000000; i++ {
    test_str := fmt.Sprintf("%s%d", text, i)
    hasher.Reset()
    io.WriteString(hasher, test_str)
    hash := fmt.Sprintf("%02x", hasher.Sum(nil))
    v, err := strconv.Atoi(hash[:6])
    if v == 0 && err == nil {
      fmt.Println(hash)
      fmt.Println(v)
      fmt.Println(err)
      found = i
      break
    }
  }

  fmt.Println(found)
}
