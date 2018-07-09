package Truco

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ancho_de_espada_mata_siete_espada(t *testing.T) {
	anchoEspada := NewAnchoEspada()
	sieteEspada := NewSieteEspada()

	assert.Equal(t, anchoEspada,
		anchoEspada.Desafia(sieteEspada), "anchoEspada win")
}

func Test_ancho_de_basto_mata_siete_espada(t *testing.T) {
	anchoBasto := NewAnchoBasto()
	sieteEspada := NewSieteEspada()

	assert.Equal(t, anchoBasto,
		anchoBasto.Desafia(sieteEspada), "anchoBasto win")
}

func Test_ancho_de_basto_mata_siete_oro(t *testing.T) {
	anchoBasto := NewAnchoBasto()
	sieteOro := NewSieteOro()

	assert.Equal(t, anchoBasto,
		sieteOro.Desafia(anchoBasto), "siete oro is defeated")
}

func Test_siete_de_espada_pierde_con_ancho_espada(t *testing.T) {
	anchoEspada := NewAnchoEspada()
	sieteEspada := NewSieteEspada()

	assert.Equal(t, anchoEspada,
		sieteEspada.Desafia(anchoEspada), "siete is defeat")
}

func Test_siete_de_espada_gana_a_siete_oro(t *testing.T) {
	sieteEspada := NewSieteEspada()
	sieteOro := NewSieteOro()

	assert.Equal(t, sieteEspada,
		sieteEspada.Desafia(sieteOro), "siete win")
}

func Test_envido_triada_jugador_uno_veintiocho_pierde_contra_triada_jugador_dos_treintaidos(t *testing.T) {
	triada1 := NewTriada(NewTresOro(), NewCincoOro(), NewAnchoEspada())
	triada2 := NewTriada(NewSieteEspada(), NewCincoEspada(), NewAnchoBasto())

	ganador := triada1.DesafiaEnvido(triada2)
	esperada := NewTriada(NewSieteEspada(), NewCincoEspada(), NewAnchoBasto())

	assert.Equal(t, esperada, ganador)
}

func Test_envido_triada_jugador_uno_treintaitres_gana_contra_triada_jugador_dos_treintaidos(t *testing.T) {
	triada1 := NewTriada(NewAnchoEspada(), NewSieteOro(), NewSeisOro())
	triada2 := NewTriada(NewAnchoBasto(), NewSieteEspada(), NewCincoEspada())

	ganador := triada1.DesafiaEnvido(triada2)
	esperada := NewTriada(NewAnchoEspada(), NewSieteOro(), NewSeisOro())

	assert.Equal(t, esperada, ganador)
}
