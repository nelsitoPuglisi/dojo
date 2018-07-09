package Truco

type AnchoBasto struct{}

func NewAnchoBasto() *AnchoBasto {
	return new(AnchoBasto)
}

func (this AnchoBasto) Desafia(carta Carta) Carta {
	return &this
}

func (this AnchoBasto) ValorEnvido(cartas []Carta) int {
	return ValorEnvido(this, cartas)
}

func (this AnchoBasto) palo() string {
	return "basto"
}

func (this AnchoBasto) numeroEnvido() int {
	return 1
}
