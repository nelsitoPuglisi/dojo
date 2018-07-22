package Truco

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Cinco_Espada_son_iguales(t *testing.T) {
	assert.Equal(t,
		NewCincoEspada(), NewCincoEspada(), "iguales")

}

func Test_Cinco_Espada_pierde_contra_el_siete_espada(t *testing.T) {
	assert.Equal(t,
		NewCincoEspada(), NewCincoEspada().Desafia(NewCuatroCopa()), "gana cinco e")

}
