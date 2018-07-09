package Truco

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_envido_de_mano_uno_pierde_contra_envido_de_mano_dos(t *testing.T) {
	envido1 := NewEnvido(NewTresOro(), NewCincoOro(), NewAnchoEspada())
	envido2 := NewEnvido(NewSieteEspada(), NewCincoEspada(), NewAnchoBasto())

	assert.Equal(t, 28, envido1.Valor(), "envido1 28")
	assert.Equal(t, 32, envido2.Valor(), "envido2 32")
	assert.True(t, envido1.Valor() < envido2.Valor(), "envido1 gana")
}

func Test_envido_valor_debe_ser_igual_para_cualquier_orden_de_cartas(t *testing.T) {
	envido1 := NewEnvido(NewSieteEspada(), NewCincoEspada(), NewAnchoBasto())
	envido2 := NewEnvido(NewSieteEspada(), NewAnchoBasto(), NewCincoEspada())
	envido3 := NewEnvido(NewAnchoBasto(), NewSieteEspada(), NewCincoEspada())
	envido4 := NewEnvido(NewCincoEspada(), NewAnchoBasto(), NewSieteEspada())

	valor1 := envido1.Valor()
	valor2 := envido2.Valor()
	valor3 := envido3.Valor()
	valor4 := envido4.Valor()

	assert.Equal(t, 32, valor1)
	assert.Equal(t, valor1, valor2)
	assert.Equal(t, valor2, valor3)
	assert.Equal(t, valor3, valor4)
}
