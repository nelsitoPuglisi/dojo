package Truco

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_jugador_1_gana_ronda(t *testing.T) {
	j1 := NewJugador()
	j2 := NewJugador()

	j1.levantar(NewAnchoEspada(), NewSieteOro(), NewCuatroCopa())
	j2.levantar(NewAnchoBasto(), NewSieteEspada(), NewTresEspada())

	bajada := j1.baja(NewAnchoEspada())              // Bajada: j1, Espada
	bajada = j2.bajaLuegoDe(bajada, NewTresEspada()) // Ronda: b(j1, 1E, j2, 3E)

	bajada = j1.bajaLuegoDe(bajada, NewSieteOro()) // Ronda:
	bajada = j2.bajaLuegoDe(bajada, NewSieteEspada())

	bajada = j2.bajaLuegoDe(bajada, NewAnchoBasto())
	bajada = j1.bajaLuegoDe(bajada, NewCuatroCopa())

	assert.Equal(t,
		bajada.Punto(), NewPunto(j1, 1), "gana j2")

}
