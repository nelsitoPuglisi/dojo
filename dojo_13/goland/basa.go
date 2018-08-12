package Truco

type Basa struct {
	jugadores []Jugador
	cartas    []Carta

	previa *Basa
}

func NewBasaConPrevia(previa *Basa, j Jugador, c Carta) *Basa {
	result := NewBasa(j, c)
	result.previa = previa
	return result

}

func NewBasa(j Jugador, c Carta) *Basa {
	result := new(Basa)

	result.jugadores = []Jugador{j}
	result.cartas = []Carta{c}

	return result
}

func NewBasaWithoutFrom(jugadorAquitar Jugador, base *Basa) *Basa {
	// This code considers only two players
	// Must be re-thought for more players
	indiceJugador := -1
	for inx, j := range base.jugadores {
		if j != jugadorAquitar {
			indiceJugador = inx
			break
		}
	}

	return NewBasa(base.jugadores[indiceJugador],
		base.cartas[indiceJugador])
}

func (this *Basa) BajaUna(j Jugador, carta Carta) *Basa {

	if len(this.jugadores) < 2 {
		this.agregar(j, carta)
		return this
	}

	this.verificarTurno(j)

	return NewBasaConPrevia(this, j, carta)
}

func (this *Basa) agregar(jugador Jugador, carta Carta) {
	this.jugadores = append(this.jugadores, jugador)
	this.cartas = append(this.cartas, carta)
}

func (this *Basa) verificarTurno(j Jugador) {
	quienPuedeJugar := this.dueno(this.cartaGanadora())

	if quienPuedeJugar != j {
		panic("jugador x no puede jugar")
	}
}

func (this *Basa) abandona(j Jugador) *Basa {
	return NewBasaWithoutFrom(j, this)
}

func (this *Basa) Punto(puntaje Tanto) *Punto {
	puntosXjugador := make(map[Jugador]int)

	this.puntosGanados(&puntosXjugador)

	var maximo int = 0
	var indice Jugador = NewNadie()

	for j, v := range puntosXjugador {
		if maximo < v {
			maximo = v
			indice = j
		}
	}

	return NewPunto(indice,
		puntaje /* un punto sin truco, retruco, etc */)
}

func (this *Basa) puntosGanados(pjn *map[Jugador]int) {
	ganador := this.dueno(this.cartaGanadora())

	(*pjn)[ganador] = ganador.(SumadorPuntos).IncrementarEn(
		(*pjn)[ganador], 1)

	if this.previa != nil {
		this.previa.puntosGanados(pjn)
	}
}

func (this *Basa) cartaGanadora() Carta {
	result := this.cartas[0]

	for _, c := range this.cartas[1:] {
		result = result.Desafia(c)
	}

	return result
}

func (this *Basa) dueno(carta Carta) Jugador {
	for i, c := range this.cartas {
		if c == carta {
			return this.jugadores[i]
		}
	}
	return NewNadie()
}
