package Truco

type Nadie struct {
}

func NewNadie() *Nadie {
	return new(Nadie)
}

func (this *Nadie) Levantar(cartas ...Carta) {
	return
}

func (this *Nadie) Baja(carta Carta) *Ronda {
	return NewRonda(this, carta)
}

func (this *Nadie) BajaLuegoDe(ronda *Ronda, carta Carta) *Ronda {
	return ronda
}

func (this *Nadie) Abandonar(ronda *Ronda) *Ronda {
	return ronda
}

func (this *Nadie) Truco(ronda *Ronda) *Ronda {
	return ronda
}

func (this *Nadie) Retruco(ronda *Ronda) *Ronda {
	return ronda
}

func (this *Nadie) Quiero(ronda *Ronda) *Ronda {
	return ronda
}

func (this *Nadie) NoQuiero(ronda *Ronda) *Ronda {
	return ronda
}

func (this *Nadie) IncrementarEn(valor int, inc int) int {
	return valor
}
