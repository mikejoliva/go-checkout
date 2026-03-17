package product

type ISpecial interface {
	Calculate(quantity int) (price int, remainder int, err error)
}

type Product struct {
	price    int
	quantity int
	special  ISpecial
}

func NewProduct(price int, special ISpecial) *Product {
	return &Product{
		price:    price,
		quantity: 0,
		special:  special,
	}
}

func (p *Product) Scanned() {
	p.quantity++
}

func (p *Product) Quantity() int {
	return p.quantity
}

func (p *Product) Total() (int, error) {
	total, remainder, err := p.special.Calculate(p.quantity)
	if err != nil {
		return 0, err
	}
	total += remainder * p.price
	return total, nil
}
