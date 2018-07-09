package Truco

type UnaBajada struct {
	jugador1 *Jugador
	carta1   Carta

	jugador2 *Jugador
	carta2   Carta
}

func NewUnaBajada(j *Jugador, c Carta) Bajada {
	result := new(UnaBajada)

	result.jugador1 = j
	result.carta1 = c

	result.jugador2 = nil
	result.carta2 = nil

	return result
}

func (this *UnaBajada) y(j *Jugador, carta Carta) Bajada {

	if this.jugador2 == nil {
		this.jugador2 = j
		this.carta2 = carta
		return NewRonda(this)
	}

	this.panicIfNotYourTurn(j)

	return NewRonda(NewUnaBajada(j, carta))
}

func (this *UnaBajada) panicIfNotYourTurn(j *Jugador) {
	winnerCard := this.carta1.Desafia(this.carta2)
	if winnerCard == this.carta1 && j != this.jugador1 {
		panic("jugador 1 puede jugar")
	}

	if winnerCard == this.carta2 && j != this.jugador2 {
		panic("jugador 2 puede jugar")
	}
}

func (this *UnaBajada) Punto() *Punto {
	return NewPunto(nil, 1)
}
