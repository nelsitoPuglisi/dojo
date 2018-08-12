package Truco

type Ninguna struct{}

func NewNinguna() *Ninguna {
	return new(Ninguna)
}

func (this Ninguna) Desafia(carta Carta) Carta {
	return carta
}

func (this Ninguna) ValorEnvido(cartas []Carta) int {
	return ValorEnvido(this, cartas)
}

func (this Ninguna) palo() string {
	return "ninguna"
}

func (this Ninguna) numeroEnvido() int {
	return 0
}
