package Truco

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Cinco_Oro_pierde_con_siete_espada(t *testing.T) {
	assert.Equal(t,
		NewSieteEspada(), NewCincoOro().Desafia(NewSieteEspada()), "pierde siete e")

}

func Test_Cinco_Oro_pierde_con_ancho_basto(t *testing.T) {
	assert.Equal(t,
		NewAnchoBasto(), NewCincoOro().Desafia(NewAnchoBasto()), "pierde ancho b")

}

func Test_Cinco_Oro_pierde_con_ancho_espada(t *testing.T) {
	assert.Equal(t,
		NewAnchoEspada(), NewCincoOro().Desafia(NewAnchoEspada()), "pierde ancho e")

}

func Test_Cinco_Oro_le_gana_al_resto(t *testing.T) {
	assert.Equal(t,
		NewTresEspada(), NewCincoOro().Desafia(NewTresEspada()), "gana tres e")

}
