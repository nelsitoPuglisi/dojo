package Truco

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_jugador_1_primera_basa_j2_juego_entonces_error_turno(t *testing.T) {
	j1 := NewUnJugador("j1")
	j2 := NewUnJugador("j2")

	j1.Levantar(NewAnchoEspada(), NewSieteOro(), NewCuatroCopa())
	j2.Levantar(NewAnchoBasto(), NewSieteEspada(), NewTresEspada())

	ronda := j1.Baja(NewAnchoEspada())
	ronda = j2.BajaLuegoDe(ronda, NewTresEspada())

	assert.Panics(t, func() {
		j2.BajaLuegoDe(ronda, NewSieteEspada())
	})
}

func Test_jugador_1_gana_ronda(t *testing.T) {
	j1 := NewUnJugador("j1")
	j2 := NewUnJugador("j2")

	j1.Levantar(NewAnchoEspada(), NewSieteOro(), NewCuatroCopa())
	j2.Levantar(NewAnchoBasto(), NewSieteEspada(), NewTresEspada())

	ronda := j1.Baja(NewAnchoEspada())
	ronda = j2.Abandonar(ronda)

	assert.Equal(t,
		NewPunto(j1, NewNoCantaron()), ronda.Punto(), "gana j1")

}

func Test_jugador_1_gana_bajada_1_y_jugador_2_Abandona(t *testing.T) {
	j1 := NewUnJugador("j1")
	j2 := NewUnJugador("j2")

	j1.Levantar(NewAnchoEspada(), NewSieteOro(), NewCuatroCopa())
	j2.Levantar(NewAnchoBasto(), NewSieteEspada(), NewTresEspada())

	ronda := j1.Baja(NewAnchoEspada())
	ronda = j2.BajaLuegoDe(ronda, NewTresEspada())

	ronda = j1.BajaLuegoDe(ronda, NewCuatroCopa())
	ronda = j2.BajaLuegoDe(ronda, NewAnchoBasto())

	ronda = j2.BajaLuegoDe(ronda, NewSieteEspada())
	ronda = j1.BajaLuegoDe(ronda, NewSieteOro())

	assert.Equal(t,
		NewPunto(j2, NewNoCantaron()), ronda.Punto(), "gana j2")

}

func Test_jugador_1_gana_bajada_1_y_jugador_2_abandona2(t *testing.T) {
	j1 := NewUnJugador("j1")
	j2 := NewUnJugador("j2")

	j1.Levantar(NewAnchoEspada(), NewSieteOro(), NewCuatroCopa())
	j2.Levantar(NewAnchoBasto(), NewCuatroEspada(), NewTresEspada())

	ronda := j1.Baja(NewAnchoEspada())
	ronda = j2.BajaLuegoDe(ronda, NewTresEspada())

	ronda = j1.BajaLuegoDe(ronda, NewCuatroCopa())
	ronda = j2.BajaLuegoDe(ronda, NewAnchoBasto())

	ronda = j2.BajaLuegoDe(ronda, NewCuatroEspada())
	ronda = j1.BajaLuegoDe(ronda, NewSieteOro())

	assert.Equal(t,
		NewPunto(j1, NewNoCantaron()), ronda.Punto(), "gana j2")

}
