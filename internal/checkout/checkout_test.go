package checkout_test

import (
	"path"
	"testing"

	"github.com/mikejoliva/go-checkout/internal/checkout"
)

func TestSuite_No_Special(t *testing.T) {
	t.Run("Checkout creation returns error when stocklist is missing", func(t *testing.T) {
		_, err := checkout.NewCheckout("missing")
		if err == nil {
			t.Error("Expected NewCheckout to return error when stocklist path is invalid")
		}
	})

	t.Run("Checkout creation does not return error when stocklist exists", func(t *testing.T) {
		_, err := checkout.NewCheckout(path.Join("..", "..", "test", "config", "valid_stocklist.yaml"))
		if err != nil {
			t.Errorf("Unexpected error creating checkout instance: %s", err)
		}
	})

	t.Run("Checkout Scan returns error when SKU is longer than 1 char", func(t *testing.T) {
		checkout, err := checkout.NewCheckout(path.Join("..", "..", "test", "config", "valid_stocklist.yaml"))
		if err != nil {
			t.Errorf("Unexpected error creating checkout instance: %s", err)
		}

		err = checkout.Scan("AA")
		if err == nil {
			t.Errorf("Expected error scanning SKU longer than 1 digit %s", err)
		}
	})

	t.Run("Checkout Scan returns error when SKU does not exist", func(t *testing.T) {
		checkout, err := checkout.NewCheckout(path.Join("..", "..", "test", "config", "valid_stocklist.yaml"))
		if err != nil {
			t.Errorf("Unexpected error creating checkout instance: %s", err)
		}

		err = checkout.Scan("Z")
		if err == nil {
			t.Errorf("Expected error scanning SKU which doesn't exist")
		}
	})

	t.Run("Checkout TotalPrice returns expected price", func(t *testing.T) {
		checkout, err := checkout.NewCheckout(path.Join("..", "..", "test", "config", "valid_stocklist.yaml"))
		if err != nil {
			t.Errorf("Unexpected error creating checkout instance: %s", err)
		}

		for i := 0; i < 3; i++ {
			err = checkout.Scan("A")
			if err != nil {
				t.Errorf("Unexpected error scanning SKU: %s", err)
			}
		}

		price, err := checkout.GetTotalPrice()
		if err != nil {
			t.Errorf("Unexpected error calculating total price %s", err)
		}

		if price != 130 {
			t.Errorf("Expected price to be 130, got: %d", price)
		}
	})
}
