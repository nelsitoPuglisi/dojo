package Truco

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Ancho_Basto_son_iguales(t *testing.T) {
	assert.Equal(t,
		NewTresEspada(), NewTresEspada(), "iguales")

}

func Test_Ancho_Basto_pierde_contra_el_ancho_espada(t *testing.T) {
	assert.Equal(t,
		NewAnchoEspada(), NewAnchoBasto().Desafia(NewAnchoEspada()), "gana ancho e")

}

func Test_Ancho_Basto_gana_al_resto(t *testing.T) {
	assert.Equal(t,
		NewAnchoBasto(), NewAnchoBasto().Desafia(NewSieteEspada()), "gana ancho e")

}
