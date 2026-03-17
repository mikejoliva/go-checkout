package product_test

import (
	"testing"

	"github.com/mikejoliva/go-checkout/internal/product"
)

func TestSuite_Multibuy_Special(t *testing.T) {
	t.Run("Calculate handles divide by 0", func(t *testing.T) {
		special := product.NewMultibuySpecial(0, 10)
		_, _, err := special.Calculate(10)
		if err == nil {
			t.Error("Expected Calculate to catch divide-by-zero error")
		}
	})

	t.Run("Returns expected value when threshold is exact", func(t *testing.T) {
		special := product.NewMultibuySpecial(3, 130)
		price, remainder, err := special.Calculate(3)
		if err != nil {
			t.Errorf("Unexpected error from multibuy calculate: %s", err)
		}
		if remainder != 0 {
			t.Errorf("Expected multibuy remainder to be 0, got: %d", remainder)
		}
		if price != 130 {
			t.Errorf("Expected multibuy price to be 130, got: %d", remainder)
		}
	})

	t.Run("Returns expected value when threshold is satisfied exactly twice", func(t *testing.T) {
		special := product.NewMultibuySpecial(3, 130)
		price, remainder, err := special.Calculate(6)
		if err != nil {
			t.Errorf("Unexpected error from multibuy calculate: %s", err)
		}
		if remainder != 0 {
			t.Errorf("Expected multibuy remainder to be 0, got: %d", remainder)
		}
		if price != 260 {
			t.Errorf("Expected multibuy price to be 260, got: %d", remainder)
		}
	})

	t.Run("Returns expected value when remainder is present", func(t *testing.T) {
		special := product.NewMultibuySpecial(3, 130)
		price, remainder, err := special.Calculate(4)
		if err != nil {
			t.Errorf("Unexpected error from multibuy calculate: %s", err)
		}
		if remainder != 1 {
			t.Errorf("Expected multibuy remainder to be 1, got: %d", remainder)
		}
		if price != 130 {
			t.Errorf("Expected multibuy price to be 130, got: %d", remainder)
		}
	})
}
