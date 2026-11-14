package main

import (
	"fmt"
	"strings" // standard package for working with text
)

// product name -> price; declared outside main(), so every function can use it
var productPrices = map[string]float64{
	"Dress":  49.99,
	"Makeup": 199.99,
	"Watch":  99.99,
	"Shoes":  89.99,
}

// returns two values: the price and whether the item was found
func calculateItemPrice(itemCode string) (float64, bool) {
	// Check if the item exists in the product list.
	basePrice, found := productPrices[itemCode] // found is false if the key is not in the map
	if found {
		return basePrice, true // return ends the function here
	}

	// If not found, check if it is a sale item.
	if strings.HasSuffix(itemCode, "Sale") { // true if itemCode ends with "Sale"
		originalItemCode := strings.TrimSuffix(itemCode, "Sale") // "WatchSale" becomes "Watch"
		basePrice, found = productPrices[originalItemCode] // "=" not ":=", both variables already exist

		if found {
			salePrice := basePrice * 0.8 // Apply a 20% discount.

			fmt.Printf( // %.2f prints a number with 2 decimal places
				"Item '%s' is on sale! Original price: $%.2f, Sale price: $%.2f\n",
				originalItemCode,
				basePrice,
				salePrice,
			)

			return salePrice, true
		}
	}

	// Item not found.
	fmt.Printf("Item code '%s' not found in the product list.\n", itemCode)
	return 0, false
}

func main() {
	fmt.Println("----- Simple Sale Order Processor -----")

	orderItems := []string{
		"Dress",
		"Makeup",
		"WatchSale",
		"ShoesSale",
		"Hat",
	}

	var subtotal float64 // zero value is 0, ready to add prices to

	fmt.Println("Processing order items...")
	fmt.Println() // no arguments: prints only a newline, which leaves a blank line

	for _, itemCode := range orderItems { // loops over each item in the slice
		price, found := calculateItemPrice(itemCode)

		if found { // items that were not found are skipped
			fmt.Printf("Added %-10s : $%.2f\n", itemCode, price) // %-10s pads the text to 10 characters, left-aligned
			subtotal += price // same as subtotal = subtotal + price
		}
	}

	fmt.Println("--------------------------------------")
	fmt.Printf("Subtotal for the order: $%.2f\n", subtotal)
}
