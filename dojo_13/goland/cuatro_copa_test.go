package Truco

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Cuatro_Copa_Espada_son_iguales(t *testing.T) {
	assert.Equal(t,
		NewCuatroCopa(), NewCuatroCopa(), "iguales")

}

func Test_Cuatro_Copa_pierde_contra_el_siete_espada(t *testing.T) {
	assert.Equal(t,
		NewTresEspada(), NewCuatroCopa().Desafia(NewTresEspada()), "gana tres e")

}
