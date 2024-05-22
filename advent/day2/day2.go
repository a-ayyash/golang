package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getSmallest(arr []int) int {
  smallest := arr[0]
  for _, v := range arr {
    if v < smallest {
      smallest = v 
    }
  }

  return smallest
}

func multiplyBy2(arr []int) []int {
  for i := range arr {
    arr[i] = arr[i] * 2
  }

  return arr
}

func constructDims(arr []int) []int {
  result_arr := make([]int, len(arr))
  result_arr[0] = arr[0] * arr[1]
  result_arr[1] = arr[1] * arr[2]
  result_arr[2] = arr[0] * arr[2]

  return result_arr;
}

func convertArrToInt(str_arr []string) []int {
  int_arr := make([]int, len(str_arr))
  for i, v := range str_arr {
    strint, err := strconv.Atoi(v)
    int_arr[i] = strint

    if err != nil {
      log.Fatal("YELEHWEEEEEE")
    }
  }


  return int_arr;
}

func add(arr []int) int {
  sum := 0
  for i := range arr {
    sum = sum + arr[i]
  }

  return sum
}

func calculateWrappingPaper(line string) int {
    comps := strings.Split(line, "x")
    comps_int := convertArrToInt(comps)
    comps_int = constructDims(comps_int)
    smallest := getSmallest(comps_int)
    comps_int = multiplyBy2(comps_int)
    return add(comps_int) + smallest
}

func get2Smallest(arr []int) int {
  slices.Sort(arr)
  return 2*(arr[0] + arr[1])
}

func cubicVolume(arr []int) int {
  sum := 1
  for i := range arr {
    sum = sum * arr[i]
  }

  return sum
}

func calculateRibbon(line string) int {
  comps := strings.Split(line, "x")
  comps_int := convertArrToInt(comps)
  sum := (comps_int[0]+comps_int[1]+comps_int[2]) - slices.Max([]int{comps_int[0],comps_int[1],comps_int[2]})
  return (2*sum) + (comps_int[0]*comps_int[1]*comps_int[2])
}

func main() {
  file, err := os.Open("data.input")
  if err != nil {
    log.Fatal(err)
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)

  grand_total_wrap := 0
  grand_total_ribbon := 0
  counter := 0

  fmt.Println(calculateRibbon("2x3x4"))
  for scanner.Scan() {
    line := scanner.Text()
    grand_total_ribbon += calculateRibbon(line)
    grand_total_wrap = grand_total_wrap + calculateWrappingPaper(line)
    counter++
  }

  fmt.Println(grand_total_wrap)
  fmt.Println(grand_total_ribbon)
  fmt.Println(counter)

  if err := scanner.Err(); err != nil {
    fmt.Fprintf(os.Stderr, "reading standard input: ", err)
  }
}
