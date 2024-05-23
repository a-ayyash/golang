package day62

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type Operation struct {
	on, off, toggle bool
}

func parseLine(arr []string) (Operation, Point, Point) {
	operation := Operation{}
	start := Point{0, 0}
	finish := Point{0, 0}

	if arr[0] == "toggle" {
		operation.toggle = true
		c := strings.Split(arr[1], ",")
		start.x, _ = strconv.Atoi(c[0])
		start.y, _ = strconv.Atoi(c[1])

		c = strings.Split(arr[3], ",")
		finish.x, _ = strconv.Atoi(c[0])
		finish.y, _ = strconv.Atoi(c[1])

	} else if arr[1] == "on" {
		operation.on = true
		c := strings.Split(arr[2], ",")
		start.x, _ = strconv.Atoi(c[0])
		start.y, _ = strconv.Atoi(c[1])

		c = strings.Split(arr[4], ",")
		finish.x, _ = strconv.Atoi(c[0])
		finish.y, _ = strconv.Atoi(c[1])
	} else {
		operation.off = true
		c := strings.Split(arr[2], ",")
		start.x, _ = strconv.Atoi(c[0])
		start.y, _ = strconv.Atoi(c[1])

		c = strings.Split(arr[4], ",")
		finish.x, _ = strconv.Atoi(c[0])
		finish.y, _ = strconv.Atoi(c[1])
	}

	return operation, start, finish
}

func main() {
	file, err := os.Open("../data.input")

	if err != nil {
		log.Fatal("error openning the input file for now")
	}

	scanner := bufio.NewScanner(file)
	grid := make([][]int, 1000)

	for i, _ := range grid {
		grid[i] = make([]int, 1000)
	}

	for scanner.Scan() {
		line := scanner.Text()
		line_arr := strings.Split(line, " ")
		operation_type, starting_point, finish_point := parseLine(line_arr)

		switch {
		case operation_type.on:
			{
				//fmt.Printf("ON: (%d, %d) TO (%d, %d)", starting_point.x, starting_point.y, finish_point.x, finish_point.y)
				for i := starting_point.x; i <= finish_point.x; i++ {
					for j := starting_point.y; j <= finish_point.y; j++ {
						grid[i][j] = grid[i][j] + 1
					}
				}
			}
		case operation_type.off:
			{
				//fmt.Printf("OFF: (%d, %d) TO (%d, %d)", starting_point.x, starting_point.y, finish_point.x, finish_point.y)
				for i := starting_point.x; i <= finish_point.x; i++ {
					for j := starting_point.y; j <= finish_point.y; j++ {
						if grid[i][j]-1 >= 0 {
							grid[i][j] = grid[i][j] - 1
						}
					}
				}
			}
		case operation_type.toggle:
			{
				//fmt.Printf("TOGGLE: (%d, %d) TO (%d, %d)", starting_point.x, starting_point.y, finish_point.x, finish_point.y)
				for i := starting_point.x; i <= finish_point.x; i++ {
					for j := starting_point.y; j <= finish_point.y; j++ {
						grid[i][j] = grid[i][j] + 2
					}
				}
			}

		}
	}

	counter := 0

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			counter += grid[i][j]
		}
	}

	fmt.Println(counter)
}
