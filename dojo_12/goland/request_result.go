package Room

type RequestResult struct {
	rooms   []*ClassRoom
	filters []Filter
}

func NewRequestResult(rooms []*ClassRoom, filters ...Filter) *RequestResult {
	rr := new(RequestResult)

	rr.rooms = rooms
	rr.filters = filters

	return rr
}

func (this *RequestResult) havingAreaAtLeast(area *Area) *RequestResult {
	this.filters = append(this.filters, NewAreaAtLeast(area))
	return this
}

func (this *RequestResult) havingExactArea(area *Area) *RequestResult {
	this.filters = append(this.filters, NewExactArea(area))
	return this
}

func (this RequestResult) chooseOne() *ClassRoom {
	for _, singleRoom := range this.rooms {
		if singleRoom.fitsAll(this.filters) {
			return singleRoom
		}
	}

	return nil
}
