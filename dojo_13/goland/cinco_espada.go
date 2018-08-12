package Truco

type CincoEspada struct{}

func NewCincoEspada() *CincoEspada {
	return new(CincoEspada)
}

func (this CincoEspada) Desafia(carta Carta) Carta {
	return &this

}

func (this CincoEspada) palo() string {
	return "espada"
}

func (this CincoEspada) numeroEnvido() int {
	return 5
}
