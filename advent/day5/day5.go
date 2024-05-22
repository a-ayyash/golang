package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

func main () {
  file, err := os.Open("data.input")

  if err != nil {
    log.Fatal("Error opening a file")
  }

  scanner := bufio.NewScanner(file)
  vowels := []string {"a", "e", "i", "o", "u"}
  vowels_counter := 0
  naughty_strings := []string {"ab", "cd", "pq", "xy"}
  nice_counter := 0
  //at_least_3_vowels := false
  two_letters_in_a_row := false
  no_naughty_strings := false

  for scanner.Scan() {
    line := scanner.Text()

    for i, v := range line {

      value := string(v)

      if vowels_counter < 3 {
        if slices.Index(vowels, value) != -1 {
          vowels_counter++
        }
      }

      if two_letters_in_a_row == false && i+1 < len(line) && value == string(line[i+1]) {
        two_letters_in_a_row = true
      }

      if no_naughty_strings == false && i+1 < len(line) {
        pair := fmt.Sprintf("%s%s", value, string(line[i+1]))

        if slices.Index(naughty_strings, pair) != -1 {
          no_naughty_strings = true
        }
      }
    }

    if vowels_counter == 3 && two_letters_in_a_row == true && no_naughty_strings == false {
      nice_counter++
    }

    vowels_counter = 0
    two_letters_in_a_row = false
    no_naughty_strings = false
  }

  fmt.Println(nice_counter)
}
