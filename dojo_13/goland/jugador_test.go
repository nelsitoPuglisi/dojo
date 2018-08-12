package Truco

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_tira_cartas_que_posee(t *testing.T) {
	j1 := NewUnJugador("jo")

	j1.Levantar(NewCuatroCopa(), NewTresEspada(), NewSieteEspada())

	assert.NotPanics(t, func() {
		j1.Baja(NewSieteEspada())
	})
}

func Test_tira_cartas_que_no_posee(t *testing.T) {
	j1 := NewUnJugador("jo")

	j1.Levantar(NewCuatroCopa(), NewTresEspada(), NewSieteEspada())

	assert.Panics(t, func() {
		j1.Baja(NewAnchoBasto())
	})
}

func Test_no_puede_tirar_dos_veces_misma_carta(t *testing.T) {
	j1 := NewUnJugador("jo")

	j1.Levantar(NewCuatroCopa(), NewTresEspada(), NewSieteEspada())

	j1.Baja(NewSieteEspada())
	assert.Panics(t, func() {
		j1.Baja(NewSieteEspada())
	})
}
