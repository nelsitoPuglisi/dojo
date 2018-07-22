package Truco

type CuatroCopa struct {
}

func NewCuatroCopa() *CuatroCopa {
	r := new(CuatroCopa)
	return r
}

func (this CuatroCopa) Desafia(carta Carta) Carta {
	return carta
}

func (this CuatroCopa) palo() string {
	return "espada"
}

func (this CuatroCopa) numeroEnvido() int {
	return 3
}
