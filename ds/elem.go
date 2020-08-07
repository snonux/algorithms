package ds

type Elem interface {
	Equal(e Elem) bool
	Lower(e Elem) bool
	LowerEqual(e Elem) bool
	Higher(e Elem) bool
	HigherEqual(e Elem) bool
	Compare(e Elem) int
	CompareCB(e Elem, lowerCB, higherCB, equalsCB func())
	Int() int
}
