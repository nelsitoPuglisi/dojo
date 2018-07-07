package Room

type AreaAtLeast struct {
	expectedArea      *Area
	keyValueToBeSence string
}

func NewAreaAtLeast(expectedArea *Area) Filter {
	af := new(AreaAtLeast)
	af.expectedArea = expectedArea
	af.keyValueToBeSence = "area"
	return *af
}

func (this AreaAtLeast) fulfilledBy(dataBag DataBag) bool {
	valueToBeSence := dataBag.value(this.keyValueToBeSence)
	areaToBeSence := valueToBeSence.(*Area)
	return this.expectedArea.lt(areaToBeSence)
}
