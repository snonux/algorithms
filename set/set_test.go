package set

import (
	"fmt"
	"testing"

	"github.com/snonux/algorithms/ds"
)

const factor int = 10
const minLength int = 1
const maxLength int = 10

func TestElementary(t *testing.T) {
	s := NewElementary()
	for i := minLength; i <= maxLength; i *= factor {
		test(s, i, t)
	}
}

func TestBST(t *testing.T) {
	s := NewBST()
	for i := minLength; i <= maxLength; i *= factor {
		test(s, i, t)
	}
}

func test(s Set, l int, t *testing.T) {
	cb := func(t *testing.T) {
		list := ds.NewRandomArrayList(l, -1)
		for i, a := range list {
			t.Log("Inserting", a, i, s)
			s.Set(a, i)
		}
		for i, a := range list {
			val, err := s.Get(a)
			if err != nil {
				t.Errorf("Element %v: %v", val, err)
			}
			if val != i {
				t.Errorf("Element is %v but expected %v", val, i)
			}
		}
	}
	t.Run(fmt.Sprintf("%d", l), cb)
}
