package Truco

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_j1_canta_truco_j2_gana_los_puntos_son_2_para_j2(t *testing.T) {
	j1 := NewUnJugador("j1")
	j2 := NewUnJugador("j2")

	j1.Levantar(NewAnchoEspada(), NewSieteOro(), NewCuatroCopa())
	j2.Levantar(NewAnchoBasto(), NewSieteEspada(), NewTresEspada())

	ronda := j1.Baja(NewAnchoEspada())
	ronda = j2.BajaLuegoDe(ronda, NewTresEspada())

	ronda = j1.Truco(ronda)
	ronda = j2.Quiero(ronda)

	ronda = j1.BajaLuegoDe(ronda, NewCuatroCopa())
	ronda = j2.BajaLuegoDe(ronda, NewAnchoBasto())

	ronda = j2.BajaLuegoDe(ronda, NewSieteEspada())
	ronda = j1.BajaLuegoDe(ronda, NewSieteOro())

	assert.Equal(t,
		NewPunto(j2, NewTruco()), ronda.Punto(), "j2 gana truco")
}

func Test_j1_canta_truco_j2_no_quiere_gana_1_punto_j1(t *testing.T) {
	j1 := NewUnJugador("j1")
	j2 := NewUnJugador("j2")

	j1.Levantar(NewAnchoEspada(), NewSieteOro(), NewCuatroCopa())
	j2.Levantar(NewAnchoBasto(), NewSieteEspada(), NewTresEspada())

	ronda := j1.Baja(NewAnchoEspada())
	ronda = j2.BajaLuegoDe(ronda, NewTresEspada())

	ronda = j1.Truco(ronda)
	ronda = j2.NoQuiero(ronda)

	assert.Equal(t,
		NewPunto(j1, NewNoCantaron()), ronda.Punto(), "j1 gana 1")
}

func Test_j1_canta_truco_j2_no_quiere_gana_2_punto_j2_por_primera(t *testing.T) {
	j1 := NewUnJugador("j1")
	j2 := NewUnJugador("j2")

	j1.Levantar(NewAnchoEspada(), NewSieteOro(), NewCuatroCopa())
	j2.Levantar(NewAnchoBasto(), NewSieteEspada(), NewTresEspada())

	ronda := j1.Baja(NewCuatroCopa())
	ronda = j2.BajaLuegoDe(ronda, NewTresEspada())

	ronda = j1.Truco(ronda)
	ronda = j2.NoQuiero(ronda)

	assert.Equal(t, NewPunto(j1, NewNoCantaron()),
		ronda.Punto(), "j1 gana 1")
}

func Test_j1_canta_truco_j2_canta_retruco_gana_los_puntos_son_2_para_j2(t *testing.T) {
	j1 := NewUnJugador("j1")
	j2 := NewUnJugador("j2")

	j1.Levantar(NewAnchoEspada(), NewSieteOro(), NewCuatroCopa())
	j2.Levantar(NewAnchoBasto(), NewSieteEspada(), NewTresEspada())

	ronda := j1.Baja(NewAnchoEspada())
	ronda = j2.BajaLuegoDe(ronda, NewTresEspada())

	ronda = j1.Truco(ronda)
	ronda = j2.Quiero(ronda)

	ronda = j1.BajaLuegoDe(ronda, NewCuatroCopa())
	ronda = j2.BajaLuegoDe(ronda, NewAnchoBasto())

	ronda = j2.Retruco(ronda)
	ronda = j1.Quiero(ronda)

	ronda = j2.BajaLuegoDe(ronda, NewSieteEspada())
	ronda = j1.BajaLuegoDe(ronda, NewSieteOro())

	assert.Equal(t,
		NewPunto(j2, NewRetruco()), ronda.Punto(), "j2 gana truco")
}

func Test_j1_canta_truco_j2_canta_retruco_j1_no_quiere_los_puntos_son_2_para_j2(t *testing.T) {
	j1 := NewUnJugador("j1")
	j2 := NewUnJugador("j2")

	j1.Levantar(NewAnchoEspada(), NewSieteOro(), NewCuatroCopa())
	j2.Levantar(NewAnchoBasto(), NewSieteEspada(), NewTresEspada())

	ronda := j1.Baja(NewAnchoEspada())
	ronda = j2.BajaLuegoDe(ronda, NewTresEspada())

	ronda = j1.Truco(ronda)
	ronda = j2.Quiero(ronda)

	ronda = j1.BajaLuegoDe(ronda, NewCuatroCopa())
	ronda = j2.BajaLuegoDe(ronda, NewAnchoBasto())

	ronda = j2.Retruco(ronda)
	ronda = j1.NoQuiero(ronda)

	assert.Equal(t,
		NewPunto(j2, NewTruco()), ronda.Punto(), "j2 gana truco")
}
