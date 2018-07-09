package Truco

type Jugador struct {
	cartas []Carta
}

func NewJugador() *Jugador {
	result := new(Jugador)
	result.cartas = []Carta{}
	return result
}

func (this *Jugador) levantar(cartas ...Carta) {
	this.cartas = cartas
	return
}

func (this *Jugador) baja(carta Carta) Bajada {
	this.usar(carta)
	return NewUnaBajada(this, carta)
}

func (this *Jugador) bajaLuegoDe(bajada Bajada, carta Carta) Bajada {
	this.usar(carta)
	return bajada.y(this, carta)
}

func (this *Jugador) usar(carta Carta) {
	ubicacion := this.panicIfAbsent(carta, this.cartas, "absent")
	this.cartas = append(this.cartas[:ubicacion], this.cartas[ubicacion+1:]...)
	return
}

func (this *Jugador) panicIfAbsent(carta Carta, cartas []Carta, msg string) int {
	for i, c := range cartas {
		if c == carta {
			return i
		}
	}

	panic(msg)
}
