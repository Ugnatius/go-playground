package main

import (
	"fmt"
	"math"
)

func main() {
	number := 3.75
	floor := math.Floor(number)
	ceil := math.Ceil(number)
	round := math.Round(number)
	trunc := math.Trunc(number)

	fmt.Println("Origin number is ", number)
	fmt.Println(floor)
	fmt.Println(ceil)
	fmt.Println(round)
	fmt.Println(trunc)
}
