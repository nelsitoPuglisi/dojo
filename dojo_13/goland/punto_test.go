package Truco

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_punto_son_iguales(t *testing.T) {
	assert.Equal(t,
		NewPunto(nil, NewTruco()), NewPunto(nil, NewTruco()), "iguales")
	assert.NotEqual(t,
		NewPunto(nil, NewNoCantaron()), NewPunto(nil, NewTruco()), "distintos")

}
