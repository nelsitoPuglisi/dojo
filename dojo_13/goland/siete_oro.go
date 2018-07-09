package Truco

type SieteOro struct{}

func NewSieteOro() *SieteOro {
	return new(SieteOro)
}

func (this SieteOro) Desafia(carta Carta) Carta {
	switch tipada := carta.(type) {
	case *AnchoBasto:
		return tipada
	case *SieteEspada:
		return tipada
	}
	return &this
}

func (this SieteOro) ValorEnvido(cartas []Carta) int {
	return ValorEnvido(this, cartas)
}

func (this SieteOro) palo() string {
	return "oro"
}

func (this SieteOro) numeroEnvido() int {
	return 7
}
