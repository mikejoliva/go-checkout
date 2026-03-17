package checkout

import (
	"fmt"

	"github.com/mikejoliva/go-checkout/internal/product"
)

type Checkout struct {
	stocklist *product.Stocklist
	scanned   map[string]*product.Product
}

func NewCheckout(stocklist_path string) (*Checkout, error) {
	stocklist, err := product.NewStocklistParser(stocklist_path)
	if err != nil {
		return nil, fmt.Errorf("failed to create stocklist parser: %s", err)
	}

	// Populate the scanned map with all known products from the stocklist
	scanned := make(map[string]*product.Product, len(stocklist.Products))
	for _, prod := range stocklist.Products {
		special, err := prod.GetSpecial()
		if err != nil {
			return nil, fmt.Errorf("error retrieving product special offer %s", err)
		}
		scanned[prod.Name] = product.NewProduct(prod.Price, special)
	}

	return &Checkout{
		stocklist,
		scanned,
	}, nil
}

func (c *Checkout) Scan(SKU string) error {
	if len(SKU) != 1 {
		return fmt.Errorf("unexpected SKU length %d for input %s", len(SKU), SKU)
	}

	if _, found := c.scanned[SKU]; !found {
		return fmt.Errorf("unexpected item ID: %s", SKU)
	}

	c.scanned[SKU].Scanned()
	return nil
}

func (c *Checkout) GetTotalPrice() (int, error) {
	total := 0
	for _, product := range c.scanned {
		price, err := product.Total()
		if err != nil {
			return 0, err
		}
		total += price
	}
	return total, nil
}
