package Truco

type Tanto interface {
	Siguiente() Tanto
	NoQuiero() Tanto
}
