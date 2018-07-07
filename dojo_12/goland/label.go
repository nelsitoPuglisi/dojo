package Room

type Label struct {
	value string
}

func NewLabel(value string) *Label {
	l := new(Label)
	l.value = value
	return l
}
