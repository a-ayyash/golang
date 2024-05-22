package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
  x, y int
}

func getSantaHouses(line string) int {

  grid := make(map[Point]int) 
  grid[Point{0,0}] = 0
  current_point := Point{0,0}

  for _, v := range line {
    switch string(v) {
      case "^": {
        current_point.y++
      } 
      case "v": {
        current_point.y--
      }
      case ">": {
        current_point.x++
      }
      case "<":
        current_point.x--
      }

      if grid[current_point] == 0 {
        grid[current_point] = 1
      } 
  }

  return len(grid)
}

func getSantaAndRobotHouses(line string) int { 
  grid := make(map[Point]int) 
  grid[Point{0,0}] = 0

  current_point := Point{0,0}
  current_robo_point := Point{0,0}

  for i, v := range line {
    if i % 2 == 0 {
      switch string(v) {
        case "^": {
          current_robo_point.y++
        } 
        case "v": {
          current_robo_point.y--
        }
        case ">": {
          current_robo_point.x++
        }
        case "<":
          current_robo_point.x--
        }
      if grid[current_robo_point] == 0 {
        grid[current_robo_point] = 1
      } 

    } else {
      switch string(v) {
        case "^": {
          current_point.y++
        } 
        case "v": {
          current_point.y--
        }
        case ">": {
          current_point.x++
        }
        case "<":
          current_point.x--
        }
      if grid[current_point] == 0 {
        grid[current_point] = 1
      } 
    }

  }

  return len(grid) 
}

func main() {
  file, err := os.Open("data.input")
  if err != nil {
    log.Fatal("Error reading file")
  }

  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    line := scanner.Text()
    fmt.Println(getSantaHouses(line))
    fmt.Println(getSantaAndRobotHouses(line))
  }
}
