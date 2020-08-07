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

func (i Integer) Compare(j Elem) int {
    jVal := j.Int()
    switch {
      case i.Val < jVal:
      return -1
      case i.Val > jVal:
      return 1
    }
	return 0
}

func (i Integer) CompareCB(j Elem, lowerCB, higherCB, equalsCB func()) {
    jVal := j.Int()
    switch {
      case i.Val < jVal:
        lowerCB()
      case i.Val > jVal:
        higherCB()
      default:
        equalsCB()
    }
}
