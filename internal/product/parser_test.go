package product_test

import (
	"path"
	"reflect"
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

	t.Run("GetSpecial returns Multibuy special type from stocklist when expected", func(t *testing.T) {
		parser, err := product.NewStocklistParser(path.Join("..", "..", "test", "config", "specials_test_stocklist.yaml"))
		if err != nil {
			t.Errorf("Unexpected error loading valid stocklist: %s", err)
		}

		special, err := parser.Products[0].GetSpecial()
		if err != nil {
			t.Errorf("Unexpected error getting products special %s", err)
		}

		if _, ok := special.(product.MultibuySpecial); !ok {
			t.Errorf("Expected spcial to be of type MultibuySpecial, but found %v", reflect.TypeOf(special))
		}
	})

	t.Run("GetSpecial returns NoSpecial special type from stocklist when expected", func(t *testing.T) {
		parser, err := product.NewStocklistParser(path.Join("..", "..", "test", "config", "specials_test_stocklist.yaml"))
		if err != nil {
			t.Errorf("Unexpected error loading valid stocklist: %s", err)
		}

		special, err := parser.Products[1].GetSpecial()
		if err != nil {
			t.Errorf("Unexpected error getting products special %s", err)
		}

		if _, ok := special.(product.NoSpecial); !ok {
			t.Errorf("Expected spcial to be of type NoSpecial, but found %v", reflect.TypeOf(special))
		}
	})

	t.Run("GetSpecial returns error when unknown special type is in stocklist", func(t *testing.T) {
		parser, err := product.NewStocklistParser(path.Join("..", "..", "test", "config", "specials_test_stocklist.yaml"))
		if err != nil {
			t.Errorf("Unexpected error loading valid stocklist: %s", err)
		}

		special, err := parser.Products[1].GetSpecial()
		if err != nil {
			t.Errorf("Unexpected error getting products special %s", err)
		}

		if _, ok := special.(product.NoSpecial); !ok {
			t.Errorf("Expected spcial to be of type NoSpecial, but found %v", reflect.TypeOf(special))
		}
	})

	t.Run("GetSpecial returns error when unknown special type is in stocklist", func(t *testing.T) {
		parser, err := product.NewStocklistParser(path.Join("..", "..", "test", "config", "specials_test_stocklist.yaml"))
		if err != nil {
			t.Errorf("Unexpected error loading valid stocklist: %s", err)
		}

		_, err = parser.Products[2].GetSpecial()
		if err == nil {
			t.Error("Expected error when special type does not exist")
		}
	})
}
