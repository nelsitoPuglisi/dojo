package trips

type Bus struct {
	section []*Section
}

func NewBus(section ...*Section) *Bus {
	b := new(Bus)
	b.section = section
	return b
}

func (b *Bus) pricePerDistance() float64 {
	return 2
}

func (b *Bus) Cost() float64 {
	accumulatedDistance := 0.0
	for _, s := range b.section {
		accumulatedDistance += s.Distance()
	}

	return accumulatedDistance * b.pricePerDistance()
}

func (b *Bus) Flight() float64 {
	return 100.0
}
