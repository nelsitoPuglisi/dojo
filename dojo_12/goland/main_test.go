package Room

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUsuarioPideAulaParaAlbergar10EstudiantesComoMinimo(t *testing.T) {
	labA := NewClassRoom(NewLabel("Lab A"), NewCapacity(5), nil)
	labB := NewClassRoom(NewLabel("Lab B"), NewCapacity(10), nil)

	requestResult := NewRequest(labA, labB).
		havingCapacityAtLeast(NewCapacity(9))

	actual := requestResult.chooseOne()

	assert.Equal(t, labB, actual, "labB")
	assert.Equal(t, NewLabel("Lab B"), actual.getLabel(), "labB")
}

func TestUsuarioPideAulaParaAlbergar10EstudiantesComoMinimoNoHay(t *testing.T) {
	labA := NewClassRoom(NewLabel("Lab A"), NewCapacity(5), nil)
	labB := NewClassRoom(NewLabel("Lab B"), NewCapacity(10), nil)

	requestResult := NewRequest(labA, labB).
		havingCapacityAtLeast(NewCapacity(30))

	var expected *ClassRoom = nil
	actual := requestResult.chooseOne()

	assert.Equal(t, expected, actual, "no room")
	assert.Equal(t, NewLabel("No Room"), actual.getLabel(), "no room label")
}

func TestUsuarioPideAulaCon30m2ComoMinimoObtieneUna(t *testing.T) {
	labA := NewClassRoom(NewLabel("Lab A"), NewCapacity(5), NewArea(30))
	labB := NewClassRoom(NewLabel("Lab B"), NewCapacity(10), nil)

	requestResult := NewRequest(labA, labB).
		havingAreaAtLeast(NewArea(29))

	actual := requestResult.chooseOne()

	assert.Equal(t, labA, actual, "lab a")
}

func TestUsuarioPideAulaCon30m2ComoMinimoConCapacidad4ComoMinimoObtieneUna(t *testing.T) {
	labA := NewClassRoom(NewLabel("Lab A"), NewCapacity(5), NewArea(30))
	labB := NewClassRoom(NewLabel("Lab B"), NewCapacity(10), nil)

	requestResult := NewRequest(labA, labB).
		havingCapacityAtLeast(NewCapacity(4)).
		havingAreaAtLeast(NewArea(29))

	actual := requestResult.chooseOne()

	assert.Equal(t, labA, actual, "lab a")
}

func TestUsuarioPideAulaCon30m2ComoMinimoConCapacidad4ComoMinimoObtieneNinguna(t *testing.T) {
	labA := NewClassRoom(NewLabel("Lab A"), NewCapacity(5), NewArea(30))
	labB := NewClassRoom(NewLabel("Lab B"), NewCapacity(10), nil)

	requestResult := NewRequest(labA, labB).
		havingCapacityAtLeast(NewCapacity(4)).
		havingAreaAtLeast(NewArea(40))

	var expected *ClassRoom = nil
	actual := requestResult.chooseOne()

	assert.Equal(t, expected, actual, "no room")
	assert.Equal(t, NewLabel("No Room"), actual.getLabel(), "no room label")
}

func TestUsuarioPideAulaCon30m2ConCapacidad4ExactamenteObtieneUna(t *testing.T) {
	labA := NewClassRoom(NewLabel("Lab A"), NewCapacity(5), NewArea(30))
	labB := NewClassRoom(NewLabel("Lab B"), NewCapacity(10), nil)
	labC := NewClassRoom(NewLabel("Lab C"), NewCapacity(15), NewArea(15))

	requestResult := NewRequest(labA, labB, labC).
		havingExactCapacity(NewCapacity(15)).
		havingExactArea(NewArea(15))

	actual := requestResult.chooseOne()

	assert.Equal(t, labC, actual, "lab c")
}
