package Truco

type Triada struct {
	cartas []Carta
}

func NewTriada(carta1, carta2, carta3 Carta) *Triada {
	result := new(Triada)
	result.cartas = []Carta{carta1, carta2, carta3}
	return result
}

func (this *Triada) envido() *Envido {
	return NewEnvidoFrom(this.cartas)
}

func (this *Triada) DesafiaEnvido(otra *Triada) *Triada {
	if this.envido().Valor() < otra.envido().Valor() {
		return otra
	}

	return this
}
