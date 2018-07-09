package Truco

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_map_cards_challenge_and(t *testing.T) {
	j1 := NewJugador()
	j2 := NewJugador()

	d := NewUnaBajada(j1, NewSieteEspada())
	d.y(j2, NewTresEspada())

	assert.NotPanics(t, func() {
		d.y(j1, NewAnchoEspada())
	})
}

func Test_map_cards_challenge_and_avoid_j2_to_play(t *testing.T) {
	j1 := NewJugador()
	j2 := NewJugador()

	d := NewUnaBajada(j1, NewSieteEspada())
	d.y(j2, NewTresEspada())

	assert.Panics(t, func() {
		d.y(j2, NewCuatroCopa())
	})
}
