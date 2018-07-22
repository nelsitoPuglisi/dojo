package Truco

type SeisOro struct{}

func NewSeisOro() *SeisOro {
	return new(SeisOro)
}

func (this SeisOro) Desafia(carta Carta) Carta {
	return &this
}

func (this SeisOro) palo() string {
	return "oro"
}

func (this SeisOro) numeroEnvido() int {
	return 6
}
