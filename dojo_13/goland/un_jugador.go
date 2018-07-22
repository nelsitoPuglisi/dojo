package Truco

type UnJugador struct {
	nombre string
	cartas []Carta
}

func NewUnJugador(nombre string) *UnJugador {
	result := new(UnJugador)
	result.cartas = []Carta{}
	result.nombre = nombre
	return result
}

func (this *UnJugador) Levantar(cartas ...Carta) {
	this.cartas = cartas
	return
}

func (this *UnJugador) Baja(carta Carta) *Ronda {
	this.usar(carta)
	return NewRonda(this, carta)
}

func (this *UnJugador) BajaLuegoDe(ronda *Ronda, carta Carta) *Ronda {
	this.usar(carta)
	return ronda.BajaUna(this, carta)
}

func (this *UnJugador) Abandonar(ronda *Ronda) *Ronda {
	return ronda.Abandona(this)
}

func (this *UnJugador) usar(carta Carta) {
	ubicacion := this.errorSiNoTieneCarta(carta, this.cartas, "absent")
	this.cartas = append(this.cartas[:ubicacion], this.cartas[ubicacion+1:]...)
	return
}

func (this *UnJugador) Truco(ronda *Ronda) *Ronda {
	return ronda.Truco()
}

func (this *UnJugador) Retruco(ronda *Ronda) *Ronda {
	return ronda.Retruco()
}

func (this *UnJugador) Quiero(ronda *Ronda) *Ronda {
	return ronda.Quiero()
}

func (this *UnJugador) NoQuiero(ronda *Ronda) *Ronda {
	return ronda.NoQuiero(this)
}

func (this *UnJugador) errorSiNoTieneCarta(carta Carta, cartas []Carta, msg string) int {
	for i, c := range cartas {
		if c == carta {
			return i
		}
	}

	panic(msg)
}

func (this *UnJugador) IncrementarEn(valor int, inc int) int {
	return valor + inc
}
