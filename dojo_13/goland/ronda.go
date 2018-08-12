package Truco

type Ronda struct {
	basas []*Basa

	puntaje Tanto
}

func NewRonda(j Jugador, c Carta) *Ronda {
	result := new(Ronda)

	result.basas = []*Basa{NewBasa(j, c)}
	result.puntaje = NewNoCantaron()

	return result
}

func (this *Ronda) basaActual() *Basa {
	return this.basas[len(this.basas)-1]
}

func (this *Ronda) esNueva(basa *Basa) bool {
	for _, b := range this.basas {
		if b == basa {
			return false
		}
	}
	return true
}

func (this *Ronda) BajaUna(j Jugador, carta Carta) *Ronda {

	otraBasa := this.basaActual().BajaUna(j, carta)

	if !this.esNueva(otraBasa) {
		return this
	}

	this.basaActual().verificarTurno(j)

	this.basas = append(this.basas, NewBasaConPrevia(this.basaActual(), j, carta))
	return this
}

func (this *Ronda) Abandona(j Jugador) *Ronda {
	this.basas = append(this.basas, this.basaActual().abandona(j))
	return this
}

func (this *Ronda) Truco() *Ronda {
	this.puntaje = this.puntaje.Siguiente()
	return this
}

func (this *Ronda) Retruco() *Ronda {
	this.puntaje = this.puntaje.Siguiente()
	return this
}

func (this *Ronda) Quiero() *Ronda {
	return this
}

func (this *Ronda) NoQuiero(j Jugador) *Ronda {
	this.Abandona(j)
	this.puntaje = this.puntaje.NoQuiero()
	return this
}

func (this *Ronda) Punto() *Punto {
	return this.basaActual().Punto(this.puntaje)
}
