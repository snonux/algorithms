package ds

import (
	"fmt"
	"math/rand"
)

type Integer struct {
	Val int
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
	return fmt.Sprintf("%d", i.Val)
}

func (i Integer) Int() int {
	return i.Val
}

func (i Integer) Equal(j Elem) bool {
	return i.Val == j.Int()
}

func (i Integer) Lower(j Elem) bool {
	return i.Val < j.Int()
}

func (i Integer) LowerEqual(j Elem) bool {
	return i.Val <= j.Int()
}

func (i Integer) Higher(j Elem) bool {
	return i.Val > j.Int()
}

func (i Integer) HigherEqual(j Elem) bool {
	return i.Val >= j.Int()
}
