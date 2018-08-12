package Truco

type Retruco struct {
	puntos int
}

func NewRetruco() *Retruco {
	result := new(Retruco)

	result.puntos = 2

	return result
}

func (this *Retruco) Siguiente() Tanto {
	return this
}

func (this *Retruco) NoQuiero() Tanto {
	return NewTruco()
}
