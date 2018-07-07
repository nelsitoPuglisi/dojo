package Room

type ClassRoom struct {
	properties map[string]interface{}
	label      *Label
}

func NewClassRoomFromMap(label *Label, properties map[string]interface{}) *ClassRoom {
	room := NewClassRoomLabeled(label)
	room.properties = properties
	return room
}

func NewClassRoom(label *Label, capacity *Capacity, area *Area) *ClassRoom {
	props := make(map[string]interface{})
	props["capacity"] = capacity
	props["area"] = area

	room := NewClassRoomFromMap(label, props)

	return room
}

func NewClassRoomLabeled(label *Label) *ClassRoom {
	room := new(ClassRoom)

	room.label = label

	return room
}

func (this ClassRoom) fitsAll(filters []Filter) bool {
	for _, f := range filters {
		if !f.fulfilledBy(this) {
			return false
		}
	}
	return true
}

func (this ClassRoom) value(propertyName string) interface{} {
	result, ok := this.properties[propertyName]

	if ok {
		return result
	}

	return nil

}

func (this *ClassRoom) getLabel() *Label {
	if this == nil {
		return NewLabel("No Room")
	}

	return this.label
}
