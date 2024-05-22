package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

func getKeys(m map[rune][]rune) []rune {
  keys := make([]rune, len(m))

  i := 0
  for v := range m {
    keys[i] = v
    i++
  }

  return keys
}

type Point struct {
  pos, char int
}

func main() {
  file, err := os.Open("data.input")

  if err != nil {
    log.Fatal("Openning data file failed")
  }

  scanner := bufio.NewScanner(file)
  nice_counter := 0
  collision_array_size := 17
  collisions_array_counter := collision_array_size - 1


  for scanner.Scan() {
    found_pair := false
    found_one := false
    graph := make(map[rune][]Point)
    line := scanner.Text()

    for i, v := range line {
      if len(graph[v]) == 0 {
        graph[v] = make([]Point, collision_array_size) 
        if i+1 < len(line) {
          graph[v][0] = Point {i+1, int(line[i+1])}
        } else {
          graph[v][0] = Point {i, -1}
        } 
        graph[v][collisions_array_counter] = Point {1 , -1}
      } else {
        //check one
        latest_appearance_index := graph[v][collisions_array_counter].pos - 1
        latest_appearance := graph[v][latest_appearance_index].pos - 1
        if i - latest_appearance == 2 {
          found_one = true
        }
        //check pair
        if i+1 < len(line) {
          if slices.ContainsFunc(graph[v], func (p Point) bool {
            return p.char == int(line[i+1]) && i - p.pos >= 1
          }) {
            found_pair = true
          } 

          current := graph[v][collisions_array_counter].pos
          graph[v][current] = Point{i+1, int(line[i+1])}
          graph[v][collisions_array_counter].pos = graph[v][collisions_array_counter].pos + 1
          
        } 
      }
    }//end of single line processing

    if found_pair && found_one {
      nice_counter++
    }
  }

  fmt.Println(nice_counter)
}
