package product_test

import (
	"path"
	"testing"

	"github.com/mikejoliva/go-checkout/internal/product"
)

func TestSuite_Parser(t *testing.T) {
	t.Run("Parser loads valid stocklist without error", func(t *testing.T) {
		_, err := product.NewStocklistParser(path.Join("..", "..", "test", "config", "valid_stocklist.yaml"))
		if err != nil {
			t.Errorf("Unexpected error loading valid stocklist: %s", err)
		}
	})

	t.Run("Parser returns error when stocklist doesn't exist", func(t *testing.T) {
		_, err := product.NewStocklistParser(path.Join("..", "..", "test", "config", "missing_stocklist.yaml"))
		if err == nil {
			t.Error("Expected missing stocklist to return error")
		}
	})

	t.Run("Parser returns error when stocklist has malformed schema", func(t *testing.T) {
		_, err := product.NewStocklistParser(path.Join("..", "..", "test", "config", "bad_schema_stocklist.yaml"))
		if err == nil {
			t.Error("Expected error loading stocklist with invalid schema")
		}
	})
}
