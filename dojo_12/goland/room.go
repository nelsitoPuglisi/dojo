package Room

type Room interface {
	fitsAll(filters []Filter) bool
}
