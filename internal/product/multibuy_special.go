package product

import (
	"fmt"
	"math"
)

type MultibuySpecial struct {
	threshold int
	price     int
}

func NewMultibuySpecial(threshold int, price int) MultibuySpecial {
	return MultibuySpecial{
		threshold: threshold,
		price:     price,
	}
}

func (s MultibuySpecial) Calculate(quantity int) (int, int, error) {
	if s.threshold == 0 {
		return 0, 0, fmt.Errorf("special threshold cannot be 0")
	}
	return int(math.Floor(float64(quantity/s.threshold))) * s.price,
		quantity % s.threshold,
		nil
}
