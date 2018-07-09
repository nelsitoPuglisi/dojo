package Truco

type Bajada interface {
	y(j *Jugador, carta Carta) Bajada
	Punto() *Punto
}
