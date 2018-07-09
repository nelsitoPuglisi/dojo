package Truco

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_tira_cartas_que_posee(t *testing.T) {
	j1 := NewJugador()

	j1.levantar(NewCuatroCopa(), NewTresEspada(), NewSieteEspada())

	assert.NotPanics(t, func() {
		j1.baja(NewSieteEspada())
	})
}

func Test_tira_cartas_que_no_posee(t *testing.T) {
	j1 := NewJugador()

	j1.levantar(NewCuatroCopa(), NewTresEspada(), NewSieteEspada())

	assert.Panics(t, func() {
		j1.baja(NewAnchoBasto())
	})
}

func Test_no_puede_tirar_dos_veces_misma_carta(t *testing.T) {
	j1 := NewJugador()

	j1.levantar(NewCuatroCopa(), NewTresEspada(), NewSieteEspada())

	j1.baja(NewSieteEspada())
	assert.Panics(t, func() {
		j1.baja(NewSieteEspada())
	})
}
