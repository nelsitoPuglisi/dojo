package Room

type Area struct {
	value int
}

func NewArea(value int) *Area {
	c := new(Area)
	c.value = value
	return c
}

func (this *Area) lt(other *Area) bool {
	if other == nil {
		return false
	}
	return this.value < other.value
}

func (this *Area) eq(other *Area) bool {
	return this.value == other.value
}
