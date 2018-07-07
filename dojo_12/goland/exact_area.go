package Room

type ExactArea struct {
	expectedArea      *Area
	keyValueToBeSence string
}

func NewExactArea(expectedArea *Area) Filter {
	af := new(ExactArea)
	af.expectedArea = expectedArea
	af.keyValueToBeSence = "area"
	return *af
}

func (this ExactArea) fulfilledBy(dataBag DataBag) bool {
	valueToBeSence := dataBag.value(this.keyValueToBeSence)
	areaToBeSence := valueToBeSence.(*Area)
	return this.expectedArea.eq(areaToBeSence)
}
