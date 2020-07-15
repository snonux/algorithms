package ds

type Comparer interface {
	Equal(a Comparer) bool
	Lower(a Comparer) bool
	LowerEqual(a Comparer) bool
	Higher(a Comparer) bool
	HigherEqual(a Comparer) bool
	Int() int
}
