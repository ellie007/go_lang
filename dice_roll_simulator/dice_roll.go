package main

import "fmt"
import "math/rand"

func main() {
  var upperBound int = 7

  diceOne := "Your 1st dice roll: " + fmt.Sprintf("%d", rand.Intn(upperBound))
  diceTwo := "Your 2nd dice roll: " + fmt.Sprintf("%d", rand.Intn(upperBound))

  fmt.Println(diceOne)
  fmt.Println(diceTwo)
}

