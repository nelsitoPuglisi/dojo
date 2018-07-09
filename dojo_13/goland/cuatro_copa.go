package Truco

type CuatroCopa struct {
}

func NewCuatroCopa() *CuatroCopa {
	r := new(CuatroCopa)
	return r
}

func (this CuatroCopa) Desafia(carta Carta) Carta {
	return &this
}

func (this CuatroCopa) ValorEnvido(cartas []Carta) int {
	return ValorEnvido(this, cartas)
}

func (this CuatroCopa) palo() string {
	return "espada"
}

func (this CuatroCopa) numeroEnvido() int {
	return 3
}
