package Truco

type Jugador interface {
	Levantar(cartas ...Carta)
	Baja(carta Carta) *Ronda
	BajaLuegoDe(ronda *Ronda, carta Carta) *Ronda
	Abandonar(ronda *Ronda) *Ronda
	Truco(ronda *Ronda) *Ronda
	Retruco(ronda *Ronda) *Ronda
	Quiero(ronda *Ronda) *Ronda
	NoQuiero(ronda *Ronda) *Ronda
}
