package main

import (
	"fmt"
	"log"
	"slices"

	"ayyash.io/greetings"
	"rsc.io/quote"
)
type Point struct {
  x, y float32
}

func main()  {

  graph := make(map[string][]Point)

  fmt.Println(len(graph["h"]))
  graph["h"] = make([]Point, 10)
  graph["h"][0] = Point {1,11}
  graph["h"][1] = Point {2,22}
  graph["h"][2] = Point {3,33}

  fmt.Println(graph)
  found := slices.Contains(graph["h"], Point {3, 33})
  fmt.Println(found)






  return
  log.SetPrefix("greetings: ")
  log.SetFlags(0)
  message, err := greetings.Hello("Ahmad")

  if err != nil {
    log.Fatal(err)
  }

  names := []string {
    "ahmad", 
    "yousef",
    "shadi",
  }

  for i := range "abcdefg"{
    fmt.Println(i)
  }

  fmt.Println("*****")
  myString := "abcdefg"
  for _, v := range myString {
    fmt.Println(string(v))
  }
  myMap := make(map[int]Point)
  myMap[11] = Point {1.1,2.5}
  myMapLiteral := map[int]string {
    20: "hello",
  }


  fmt.Println(myMap)
  fmt.Println(myMapLiteral[20])
  fmt.Println("*****")

  fmt.Println("----")
  messages, err := greetings.Hellos(names)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(messages)
  fmt.Println("----")

  fmt.Println(message)
  fmt.Println("Hello, World!")
  fmt.Println(quote.Go())
}
