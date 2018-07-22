package Truco

type NoCantaron struct {
	puntos int
}

func NewNoCantaron() *NoCantaron {
	result := new(NoCantaron)
	result.puntos = 1
	return result
}

func (this *NoCantaron) Puntos() int {
	return this.puntos
}

func (this *NoCantaron) Siguiente() Tanto {
	return NewTruco()
}

func (this *NoCantaron) NoQuiero() Tanto {
	return this
}
