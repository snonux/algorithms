package ds

import (
	"math/rand"
)

type Integer int

func RandomIntegers(length, max int) ArrayList {
	a := make(ArrayList, length)
	for i := 0; i < length; i++ {
		if max > 0 {
			a[i] = Integer(rand.Intn(max))
			continue
		}
		a[i] = Integer(rand.Int())
	}
	return a
}

func AscendingIntegers(length int) ArrayList {
	a := make(ArrayList, length)
	for i := 0; i < length; i++ {
		a[i] = Integer(i)
	}
	return a
}

func DescendingIntegers(length int) ArrayList {
	a := make(ArrayList, length)
	j := length
	for i := 0; i < length; i++ {
		a[i] = Integer(j)
		j--
	}
	return a
}

func (i Integer) Equal(j Integer) bool {
	return i == j
}

func (i Integer) Lower(j Integer) bool {
	return i < j
}

func (i Integer) LowerEqual(j Integer) bool {
	return i <= j
}

func (i Integer) Higher(j Integer) bool {
	return i > j
}

func (i Integer) HigherEqual(j Integer) bool {
	return i >= j
}

func (i Integer) Compare(j Integer) int {
	switch {
	case i < j:
		return -1
	case i > j:
		return 1
	}
	return 0
}

func (i Integer) CompareCB(j Integer, lowerCB, higherCB, equalsCB func()) {
	switch {
	case i < j:
		lowerCB()
	case i > j:
		higherCB()
	default:
		equalsCB()
	}
}
