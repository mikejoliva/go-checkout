package product

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Stocklist struct {
	Products []struct {
		Name    string `yaml:"name"`
		Price   int    `yaml:"price"`
		Special *struct {
			Type   string        `yaml:"type"`
			Params yaml.MapSlice `yaml:"params"`
		}
	} `yaml:"products"`
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
