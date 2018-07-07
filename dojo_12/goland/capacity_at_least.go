package Room

type CapacityAtLeast struct {
	expectedCapacity  *Capacity
	keyValueToBeSence string
}

func NewCapacityAtLeast(expectedCapacity *Capacity) Filter {
	af := new(CapacityAtLeast)
	af.expectedCapacity = expectedCapacity
	af.keyValueToBeSence = "capacity"
	return *af
}

func (this CapacityAtLeast) fulfilledBy(dataBag DataBag) bool {
	valueToBeSence := dataBag.value(this.keyValueToBeSence)
	capacityToBeSence := valueToBeSence.(*Capacity)
	return this.expectedCapacity.lt(capacityToBeSence)
}
