package Room

type Filter interface {
	fulfilledBy(dataBag DataBag) bool
}
