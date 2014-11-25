package main

import "fmt"
import "math/rand"
import "strconv"

func main() {
  firstDice := strconv.Itoa(roll_dice)
  secondDice := strconv.Itoa(roll_dice)
  fmt.Println("Your 1st roll: " + firstDice)
  fmt.Println("Your 2nd roll: " + secondDice)
}

func roll_dice() int {
  return rand.Intn(6)
}

