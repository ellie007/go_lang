package main

import "fmt"
import "math/rand"
import "time"

func generateRandNum(upperBound int) int {
  return rand.Intn(upperBound) + 1
}

func main() {
  rand.Seed(time.Now().Unix())

  diceOne := "Your 1st dice roll: " + fmt.Sprintf("%d", generateRandNum(6))
  diceTwo := "Your 2nd dice roll: " + fmt.Sprintf("%d", generateRandNum(6))

  fmt.Println(diceOne)
  fmt.Println(diceTwo)
}






