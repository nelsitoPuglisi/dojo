package Truco

type AnchoBasto struct{}

func NewAnchoBasto() *AnchoBasto {
	return new(AnchoBasto)
}

func (this AnchoBasto) Desafia(carta Carta) Carta {
	switch tipada := carta.(type) {
	case *AnchoEspada:
		return tipada
	}
	return &this
}

func (this AnchoBasto) palo() string {
	return "basto"
}

func (this AnchoBasto) numeroEnvido() int {
	return 1
}
