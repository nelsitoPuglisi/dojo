package Truco

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_son_iguales(t *testing.T) {
	assert.Equal(t,
		NewSieteOro(), NewSieteOro(), "iguales")

}

func Test_pierde_contra_el_siete_espada(t *testing.T) {
	assert.Equal(t,
		NewSieteEspada(), NewSieteOro().Desafia(NewSieteEspada()), "gana siete e")

}
