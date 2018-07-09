package Truco

type AnchoEspada struct{}

func NewAnchoEspada() *AnchoEspada {
	return new(AnchoEspada)
}

func (this AnchoEspada) Desafia(carta Carta) Carta {
	return &this
}

func (this AnchoEspada) ValorEnvido(cartas []Carta) int {
	return ValorEnvido(this, cartas)
}

func (this AnchoEspada) palo() string {
	return "espada"
}

func (this AnchoEspada) numeroEnvido() int {
	return 1
}
