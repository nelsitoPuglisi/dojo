package Room

type DataBag interface {
	value(property string) interface{}
}
