package src

import (
  "fmt"
  "os"
  "bufio"
  "unicode"
  "strconv"
  "slices"
)

func Day1Part1() {
  data, _ := os.Open("src/day1_part1_input.txt")

  defer data.Close()

  scanner := bufio.NewScanner(data)

  var result int

  for scanner.Scan() {
    acc := ""
    for _, char := range scanner.Text() {
      if unicode.IsDigit(char) {
        acc = acc + string(char)
      }
    }

    if len(acc) == 1 {
      acc = acc + acc
    } else if len(acc) > 2 {
      lastValue := acc[len(acc)-1:]
      acc = acc[0:1] + lastValue
    }

    accInt, _ := strconv.Atoi(acc)
    result = result + accInt
  }

  fmt.Println(result)
}

func Day1Part2() {
  data, _ := os.Open("src/day1_part2_input.txt")

  defer data.Close()

  scanner := bufio.NewScanner(data)

  numbers := []string {
    "one",
    "two",
    "three",
    "four",
    "five",
    "six",
    "seven",
    "eight",
    "nine",
  }

  // var result int

  for scanner.Scan() {
    acc := ""
    for _, char := range scanner.Text() {
      if unicode.IsDigit(char) {
        acc = acc + string(char)
      } else if slices.Contains(numbers, string(char)) {
        // get index of char in numbers string and add 1
        // regex?
        acc = acc + numbers[0:i + 1]
      }
    }

    fmt.Println(acc)

    // if len(acc) == 1 {
    //   acc = acc + acc
    // } else if len(acc) > 2 {
    //   lastValue := acc[len(acc)-1:]
    //   acc = acc[0:1] + lastValue
    // }

    // accInt, _ := strconv.Atoi(acc)
    // result = result + accInt
  }

  // fmt.Println(result)
}
