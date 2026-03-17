package product_test

import (
	"testing"

	"github.com/mikejoliva/go-checkout/internal/product"
)

func TestSuite_No_Special(t *testing.T) {
	t.Run("Returns 0 price and same remainder as quantity passed in", func(t *testing.T) {
		special := product.NewNoSpecial()
		price, remainder, err := special.Calculate(10)
		if err != nil {
			t.Errorf("Unexpected error from NoSpecial calulcate: %s", err)
		}
		if price != 0 {
			t.Errorf("Expected price from NoSpecial to be 0, got: %d", price)
		}
		if remainder != 10 {
			t.Errorf("Expected remainder from NoSpecial to be 10, got: %d", price)
		}
	})
}
