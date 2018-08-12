package Truco

type CuatroEspada struct {
}

func NewCuatroEspada() *CuatroEspada {
	r := new(CuatroEspada)
	return r
}

func (this CuatroEspada) Desafia(carta Carta) Carta {
	return carta
}

func (this CuatroEspada) palo() string {
	return "espada"
}

func (this CuatroEspada) numeroEnvido() int {
	return 3
}
