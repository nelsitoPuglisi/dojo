package Truco

type SieteEspada struct{}

func NewSieteEspada() *SieteEspada {
	return new(SieteEspada)
}

func (this SieteEspada) Desafia(carta Carta) Carta {
	switch tipada := carta.(type) {
	case *AnchoEspada:
		return tipada
	}
	return &this
}

func (this SieteEspada) palo() string {
	return "espada"
}

func (this SieteEspada) numeroEnvido() int {
	return 7
}
