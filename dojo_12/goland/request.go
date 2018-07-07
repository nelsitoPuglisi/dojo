package Room

type Request struct {
	rooms []*ClassRoom
}

func NewRequest(rooms ...*ClassRoom) *Request {
	r := new(Request)
	r.rooms = rooms
	return r
}

func (r Request) havingCapacityAtLeast(capacity *Capacity) *RequestResult {
	return NewRequestResult(r.rooms, NewCapacityAtLeast(capacity))
}

func (r Request) havingExactCapacity(capacity *Capacity) *RequestResult {
	return NewRequestResult(r.rooms, NewExactCapacity(capacity))
}

func (r Request) havingAreaAtLeast(area *Area) *RequestResult {
	return NewRequestResult(r.rooms, NewAreaAtLeast(area))
}
