package Truco

type Carta interface {
	Desafia(carta Carta) Carta

	numeroEnvido() int

	palo() string
}

func ValorEnvido(carta Carta, otrasCartas []Carta) int {
	resultado := carta.numeroEnvido()

	for _, c := range otrasCartas {
		if c == carta {
			continue
		}

		if carta.palo() != c.palo() {
			continue
		}

		resultado = resultado + c.numeroEnvido()
	}

	return resultado
}
