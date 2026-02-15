package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("Enter numbers")
	_, _ = fmt.Scan(&a, &b, &c)
	if (a+b > c) && (a+c > b) && (b+c > a) {
		fmt.Println("Exists")
	} else {
		fmt.Println("Not exist")
	}
}
