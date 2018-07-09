package Truco

type TresEspada struct {
}

func NewTresEspada() *TresEspada {
	r := new(TresEspada)
	return r
}

func (this TresEspada) Desafia(carta Carta) Carta {
	return &this
}

func (this TresEspada) ValorEnvido(cartas []Carta) int {
	return ValorEnvido(this, cartas)
}

func (this TresEspada) palo() string {
	return "copa"
}

func (this TresEspada) numeroEnvido() int {
	return 4
}
