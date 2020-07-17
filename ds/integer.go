package ds

import (
	"fmt"
	"math/rand"
)

type Integer struct {
	val int
}

func RandomIntegers(length, max int) ArrayList {
	a := make(ArrayList, length)
	for i := 0; i < length; i++ {
		a[i] = Integer{rand.Intn(max)}
	}
	return a
}

func SortedIntegers(length int) ArrayList {
	a := make(ArrayList, length)
	for i := 0; i < length; i++ {
		a[i] = Integer{i}
	}
	return a
}

func ReverseSortedIntegers(length int) ArrayList {
	a := make(ArrayList, length)
	j := length
	for i := 0; i < length; i++ {
		a[i] = Integer{j}
		j--
	}
	return a
}

func (i Integer) String() string {
	return fmt.Sprintf("%d", i.val)
}

func (i Integer) Int() int {
	return i.val
}

func (i Integer) Equal(j Comparer) bool {
	return i.val == j.Int()
}

func (i Integer) Lower(j Comparer) bool {
	return i.val < j.Int()
}

func (i Integer) LowerEqual(j Comparer) bool {
	return i.val <= j.Int()
}

func (i Integer) Higher(j Comparer) bool {
	return i.val > j.Int()
}

func (i Integer) HigherEqual(j Comparer) bool {
	return i.val >= j.Int()
}
