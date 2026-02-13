package main

import (
	"fmt"
	"math"
)

func main() {
	price := 19.99
	discount := 7
	discountAmount := price * (float64(discount) / 100)
	priceLast := price - discountAmount
	roundedPrice := math.Round(priceLast*100) / 100
	fmt.Printf("Price without discount is %.2f\n", price)
	fmt.Printf("Discount %d%%: %.2f.\n", discount, discountAmount)
	fmt.Printf("Total to pay %.2f.", roundedPrice)
}
