package main

import (
	"fmt"
)

func main() {
	// ...
	var a, b int
	fmt.Println("Please, enter your numbers ")
	_, _ = fmt.Scan(&a, &b)
	sum := a + b
	fmt.Printf("Сумма: %d", sum)
}
