package Truco

type CincoOro struct{}

func NewCincoOro() *CincoOro {
	return new(CincoOro)
}

func (this CincoOro) Desafia(carta Carta) Carta {
	return &this

}

func (this CincoOro) ValorEnvido(cartas []Carta) int {
	return ValorEnvido(this, cartas)
}

func (this CincoOro) palo() string {
	return "oro"
}

func (this CincoOro) numeroEnvido() int {
	return 5
}
