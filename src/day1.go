package src

import (
  "fmt"
  "os"
  "bufio"
  "unicode"
  "strconv"
)

func Day1() {
  data, _ := os.Open("src/day1_input.txt")

  defer data.Close()

  scanner := bufio.NewScanner(data)

  var result int

  for scanner.Scan() {
    // fmt.Println(scanner.Text())

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

    // fmt.Println(acc)
    accInt, _ := strconv.Atoi(acc)
    result = result + accInt
  }

  fmt.Println(result)
}
