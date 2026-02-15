package main

import "fmt"

func main() {
	var number int
	_, _ = fmt.Scan(&number)
	isEven := number%2 == 0
	fmt.Printf("Число %d чётное: %t\n", number, isEven)

}
