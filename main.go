package main

import (
	"fmt"
	"os"
	"path"

	"github.com/mikejoliva/go-checkout/internal/checkout"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Expected exactly 1 argument with SKU code(s)!\n")
		os.Exit(1)
	}

	checkout, err := checkout.NewCheckout(path.Join("config", "stocklist.yaml"))
	if err != nil {
		fmt.Printf("Failed to create checkout: %s\n", err)
		os.Exit(1)
	}

	product_codes := os.Args[1]
	for _, code := range product_codes {
		err = checkout.Scan(string(code))
		if err != nil {
			fmt.Printf("Failed to scan code %s: %s\n", string(code), err)
			os.Exit(1)
		}
	}

	total, err := checkout.GetTotalPrice()
	if err != nil {
		fmt.Printf("Failed to calculate total price: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Total price: %d\n", total)
}
