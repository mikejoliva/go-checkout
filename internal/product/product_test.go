package product_test

import (
	"testing"

	"github.com/mikejoliva/go-checkout/internal/product"
)

func TestSuite_Product(t *testing.T) {
	t.Run("Scanned product increments quantity", func(t *testing.T) {
		product := product.NewProduct(50, product.NewMultibuySpecial(3, 130))
		product.Scanned()

		if product.Quantity() != 1 {
			t.Errorf("Expected product to have quantity of 1, got: %d", product.Quantity())
		}
	})

	t.Run("Product with Multibuy matches expected price", func(t *testing.T) {
		product := product.NewProduct(50, product.NewMultibuySpecial(3, 130))

		// Scan 3 times to hit multibuy price
		product.Scanned()
		product.Scanned()
		product.Scanned()

		price, err := product.Total()
		if err != nil {
			t.Errorf("Unexpected error from calculating product total: %s", err)
		}

		if price != 130 {
			t.Errorf("Expected price to be 130, got: %d", price)
		}
	})

	t.Run("Product with NoSpecial matches expected price", func(t *testing.T) {
		product := product.NewProduct(50, product.NewNoSpecial())

		product.Scanned()
		product.Scanned()
		product.Scanned()

		price, err := product.Total()
		if err != nil {
			t.Errorf("Unexpected error from calculating product total: %s", err)
		}

		if price != 150 {
			t.Errorf("Expected price to be 130, got: %d", price)
		}
	})
}
