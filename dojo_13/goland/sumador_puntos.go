package Truco

type SumadorPuntos interface {
	/***
	* Incrementa 'valor' en 'inc' y retorna 'valor + inc'.
	 */
	IncrementarEn(valor int, inc int) int
}
