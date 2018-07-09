package Truco

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_le_gana_a_siete_oro(t *testing.T) {
	assert.Equal(t,
		NewSieteEspada(), NewSieteEspada().Desafia(NewSieteOro()), "gana siete e")

}
