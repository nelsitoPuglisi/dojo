package Truco

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_punto_son_iguales(t *testing.T) {
	assert.Equal(t,
		NewPunto(nil, 2), NewPunto(nil, 2), "iguales")
	assert.NotEqual(t,
		NewPunto(nil, 1), NewPunto(nil, 2), "distintos")

}
