package main

func Sum(arr [5]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
