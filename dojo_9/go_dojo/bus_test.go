package trips

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBusSingleSection(t *testing.T) {
	b := NewBus(
		NewSection(
			NewPlaceZero("BA"), NewPlace("Trelew", 1600)))

	assert.Equal(t, b.Cost(), 3200.0, "distance must be 1600")

}

func TestBusFromBsAsToBahiaToTrellew(t *testing.T) {
	bsas := NewPlaceZero("BA")
	trellew := NewPlace("Trelew", 1600)
	bahia := NewPlace("Bahía", 600)
	b := NewBus(
		NewSection(bsas, bahia), NewSection(bahia, trellew))

	assert.Equal(t, b.Cost(), 3200.0, "distance must be 3200")

}

func TestBusFlightAndCostIts(test *testing.T) {
	bsas := NewPlaceZero("BA")
	trellew := NewPlace("Trelew", 1600)
	bahia := NewPlace("Bahía", 600)

	var b interface{}

	b = NewBus(
		NewSection(bsas, bahia), NewSection(bahia, trellew))

	var t Transport = b.(Transport)
	assert.Equal(test, t.Cost(), 3200.0, "distance must be 3200")

	var w Winged = t.(Winged)
	assert.Equal(test, w.Flight(), 100.0, "Flight")

}
