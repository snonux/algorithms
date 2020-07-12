package ds

import "math/rand"

type Integer int

func RandomIntegers(length, max int) []Comparer {
	a := make([]Comparer, length)
	for i := 0; i < length; i++ {
		a[i] = Integer(rand.Intn(max))
	}
	return a
}

func (i Integer) LowerThan(j Comparer) bool {
	return i < j.(Integer)
}

func (i Integer) HigherThan(j Comparer) bool {
	return i > j.(Integer)
}

func (i Integer) Equals(j Comparer) bool {
	return i == j.(Integer)
}
