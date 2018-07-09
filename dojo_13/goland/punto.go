package Truco

type Punto struct {
	jugador *Jugador
	puntaje int
}

func NewPunto(jugador *Jugador, puntaje int) *Punto {
	result := new(Punto)
	result.jugador = jugador
	result.puntaje = puntaje
	return result
}
