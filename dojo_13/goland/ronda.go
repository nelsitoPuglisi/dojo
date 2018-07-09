package Truco

type Ronda struct {
	bajadas []Bajada
}

func NewRonda(bajada Bajada) *Ronda {
	result := new(Ronda)

	result.bajadas = []Bajada{bajada}

	return result
}

func (this *Ronda) y(j *Jugador, carta Carta) Bajada {
	this.bajadas[0].y(j, carta)
	return this.enCurso.y(j, carta)
}

func (this *Ronda) Punto() *Punto {
	return NewPunto(nil, 1)
}
