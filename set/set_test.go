package set

import (
	"fmt"
	"testing"

	"github.com/snonux/algorithms/ds"
	"github.com/snonux/algorithms/sort"
)

const factor int = 10
const minLength int = 1
const maxLength int = 10

// Store results here to avoid compiler optimizations
var benchResult int

func TestElementary(t *testing.T) {
	s := NewElementary()
	for i := minLength; i <= maxLength; i *= factor {
		test(s, i, t)
	}
}

func TestBST(t *testing.T) {
	for i := minLength; i <= maxLength; i *= factor {
		test(NewBST(), i, t)
	}
}

func TestBalancedBST(t *testing.T) {
	for i := minLength; i <= maxLength; i *= factor {
		test(NewBalancedBST(), i, t)
	}
}

func test(s Set, l int, t *testing.T) {
	cb := func(t *testing.T) {
		vals := ds.NewRandomArrayList(l, -1)
		keys := ds.NewRandomArrayList(l, -1)
		mapping := make(map[int]int, l)

		for i, key := range keys {
			val := vals[i]
			mapping[key] = val

			t.Log("Inserting", key, val)
			s.Set(key, val)
		}

		for _, key := range sort.Shuffle(keys) {
			val, err := s.Get(key)
			if err != nil {
				t.Errorf("Element %v->%v: %v\n", key, val, err)
			}

			val2 := mapping[key]
			if val2 != val {
				t.Errorf("Element is %v->%v but expected %v\n", key, val, val2)
			}
			t.Log("Got", key, val, s)
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

func BenchmarkBalancedBST(t *testing.B) {
	s := NewBalancedBST()
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
