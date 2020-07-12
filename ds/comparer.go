package ds

type Comparer interface {
	LowerThan(a Comparer) bool
	HigherThan(a Comparer) bool
	Equals(a Comparer) bool
}

type CompareList []Comparer
