package set

import (
	"fmt"
	"testing"

	"github.com/snonux/algorithms/ds"
)

const factor int = 10
const minLength int = 1
const maxLength int = 10000

// Store results here to avoid compiler optimizations
var benchResult int

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

func BenchmarkElementary(t *testing.B) {
	s := NewElementary()
	for i := minLength; i <= maxLength; i *= factor {
		benchmark(s, i, t)
	}
}

func BenchmarkBST(t *testing.B) {
	s := NewBST()
	for i := minLength; i <= maxLength; i *= factor {
		benchmark(s, i, t)
	}
}

func benchmark(s Set, l int, b *testing.B) {
	list := ds.NewRandomArrayList(l, -1)

	b.Run(fmt.Sprintf("random(%d)", l), func(b *testing.B) {
		b.ResetTimer()
		for i, a := range list {
			s.Set(a, i)
		}
		for _, a := range list {
			benchResult, _ = s.Get(a)
		}
	})
}
