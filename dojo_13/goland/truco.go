package Truco

type Truco struct {
	puntos int
}

func NewTruco() *Truco {
	result := new(Truco)

	result.puntos = 2

	return result
}

func (this *Truco) Siguiente() Tanto {
	return NewRetruco()
}

func (this *Truco) NoQuiero() Tanto {
	return NewNoCantaron()
}
