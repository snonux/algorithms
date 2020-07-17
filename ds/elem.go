package ds

type Elem interface {
	Equal(a Elem) bool
	Lower(a Elem) bool
	LowerEqual(a Elem) bool
	Higher(a Elem) bool
	HigherEqual(a Elem) bool
	Int() int
}
