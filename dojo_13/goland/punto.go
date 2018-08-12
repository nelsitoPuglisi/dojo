package Truco

type Punto struct {
	jugador Jugador
	puntaje Tanto
}

func NewPunto(j Jugador, puntaje Tanto) *Punto {
	result := new(Punto)

	result.jugador = j
	result.puntaje = puntaje

	return result
}
