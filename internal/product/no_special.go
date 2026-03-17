package product

type NoSpecial struct{}

func NewNoSpecial() NoSpecial {
	return NoSpecial{}
}

func (s NoSpecial) Calculate(quantity int) (int, int, error) {
	return 0, quantity, nil
}
