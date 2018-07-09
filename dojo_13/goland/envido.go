package Truco

type Envido struct {
	cartas []Carta
}

func NewEnvido(carta1, carta2, carta3 Carta) *Envido {
	result := new(Envido)
	result.cartas = []Carta{carta1, carta2, carta3}
	return result
}

func NewEnvidoFrom(cartas []Carta) *Envido {
	result := new(Envido)
	result.cartas = cartas
	return result
}

func (this *Envido) Valor() int {
	maximo := 0

	for _, c := range this.cartas {
		valor := ValorEnvido(c, this.cartas)

		if valor > maximo {
			maximo = valor
		}
	}

	return maximo + 20
}
