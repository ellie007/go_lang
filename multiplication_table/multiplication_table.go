package main

import "fmt"

func main() {
  multiplicationTable(12)
}

func multiplicationTable(i int) int {

  for j:= 1; j <= i; j++ {
    var slice []int
    for _, value := range makeBase(i) {
      slice = append(slice, value * j)
    }
    fmt.Println(slice)
  }

  return 0
}

func makeBase(i int) []int {
  var baseSlice = make([]int, i)

  for j:=0; j < i; j++ {
    baseSlice[j] = j + 1
  }

  return baseSlice
}

