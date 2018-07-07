package Room

type Capacity struct {
	value int
}

func NewCapacity(value int) *Capacity {
	c := new(Capacity)
	c.value = value
	return c
}

func (c Capacity) lt(other *Capacity) bool {
	return c.value < other.value
}

func (c Capacity) eq(other *Capacity) bool {
	return c.value == other.value
}
