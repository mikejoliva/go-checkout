package product

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// I wish Go supported generic member functions so this could be part of the
// the Special struct.
func getKeyFromMapSlice[T any](k string, s yaml.MapSlice) (T, error) {
	var casted T
	for _, item := range s {
		if item.Key == k {
			casted, ok := item.Value.(T)
			if !ok {
				return casted, fmt.Errorf("failed to cast value: %v", item.Value)
			}
			return casted, nil
		}
	}
	return casted, fmt.Errorf("failed to find key %s in slice: %#v", k, s)
}

type StocklistProduct struct {
	Name    string `yaml:"name"`
	Price   int    `yaml:"price"`
	Special *struct {
		Type   string        `yaml:"type"`
		Params yaml.MapSlice `yaml:"params"`
	}
}

func (p *StocklistProduct) GetSpecial() (ISpecial, error) {
	if p.Special == nil {
		return NewNoSpecial(), nil
	}

	switch p.Special.Type {
	case string(MultiBuy):
		threshold, err := getKeyFromMapSlice[int]("threshold", p.Special.Params)
		if err != nil {
			return nil, err
		}
		price, err := getKeyFromMapSlice[int]("price", p.Special.Params)
		if err != nil {
			return nil, err
		}
		return NewMultibuySpecial(threshold, price), nil
	default:
		return nil, fmt.Errorf("unknown special type: %s", p.Special.Type)
	}
}

type Stocklist struct {
	Products []StocklistProduct `yaml:"products"`
}

func NewStocklistParser(path string) (*Stocklist, error) {
	raw_list, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open stocklist file: %s", err)
	}

	stocklist := &Stocklist{}
	err = yaml.Unmarshal(raw_list, stocklist)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal stocklist: %s", err)
	}

	return stocklist, nil
}
