package Truco

type SieteOro struct{}

func NewSieteOro() *SieteOro {
	return new(SieteOro)
}

func (this SieteOro) Desafia(carta Carta) Carta {
	switch tipada := carta.(type) {
	case *AnchoEspada:
		return tipada
	case *AnchoBasto:
		return tipada
	case *SieteEspada:
		return tipada
	}
	return &this
}

func (this SieteOro) palo() string {
	return "oro"
}

func (this SieteOro) numeroEnvido() int {
	return 7
}
