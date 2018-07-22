package Truco

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Ancho_Espada_son_iguales(t *testing.T) {
	assert.Equal(t,
		NewTresEspada(), NewTresEspada(), "iguales")

}

func Test_Ancho_Espada_pierde_contra_el_siete_espada(t *testing.T) {
	assert.Equal(t,
		NewTresEspada(), NewTresEspada().Desafia(NewCuatroCopa()), "gana tres e")

}
