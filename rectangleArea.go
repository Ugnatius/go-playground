package main

import (
	"fmt"
)

func main() {
	// ...
	var a float64
	var b float64
	_, _ = fmt.Scan(&a, &b)
	var area = a * b
	fmt.Printf("%f", area)
}
