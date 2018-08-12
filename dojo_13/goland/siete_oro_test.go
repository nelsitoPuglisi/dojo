package Truco

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Siete_Oro_pierde_con_siete_espada(t *testing.T) {
	assert.Equal(t,
		NewSieteEspada(), NewSieteOro().Desafia(NewSieteEspada()), "pierde siete e")

}

func Test_Siete_Oro_pierde_con_ancho_basto(t *testing.T) {
	assert.Equal(t,
		NewAnchoBasto(), NewSieteOro().Desafia(NewAnchoBasto()), "pierde ancho b")

}

func Test_Siete_Oro_pierde_con_ancho_espada(t *testing.T) {
	assert.Equal(t,
		NewAnchoEspada(), NewSieteOro().Desafia(NewAnchoEspada()), "pierde ancho e")

}

func Test_Siete_Oro_le_gana_al_resto(t *testing.T) {
	assert.Equal(t,
		NewSieteOro(), NewSieteOro().Desafia(NewTresEspada()), "gana tres e")

}
