package Room

type ExactCapacity struct {
	expectedCapacity  *Capacity
	keyValueToBeSence string
}

func NewExactCapacity(expectedCapacity *Capacity) Filter {
	af := new(ExactCapacity)
	af.expectedCapacity = expectedCapacity
	af.keyValueToBeSence = "capacity"
	return *af
}

func (this ExactCapacity) fulfilledBy(dataBag DataBag) bool {
	valueToBeSence := dataBag.value(this.keyValueToBeSence)
	capacityToBeSence := valueToBeSence.(*Capacity)
	return this.expectedCapacity.eq(capacityToBeSence)
}
