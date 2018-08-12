package Truco

type TresOro struct {
}

func NewTresOro() *TresOro {
	r := new(TresOro)
	return r
}

func (this TresOro) Desafia(carta Carta) Carta {
	return &this
}

func (this TresOro) palo() string {
	return "oro"
}

func (this TresOro) numeroEnvido() int {
	return 3
}
