package main

import "fmt"

func main() {
	var putNumber float64
	for {
		fmt.Println("Please enter a number:")
		_, err := fmt.Scan(&putNumber)
		if err != nil {
			fmt.Println("Please enter a valid number, not a word")
			var dump string
			_, _ = fmt.Scan(&dump)
			continue
		}

		if putNumber > 0 {
			fmt.Printf("Your number %.2f is greater than 0!", putNumber)
			break
		} else {
			fmt.Println("Please enter a positive valid number ")
		}
	}
	fmt.Println("Program is sucsesfull")
}
