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

func (this TresEspada) palo() string {
	return "espada"
}

func (this TresEspada) numeroEnvido() int {
	return 3
}
