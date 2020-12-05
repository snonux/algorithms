package set

import (
	"fmt"
	"testing"

	"github.com/snonux/algorithms/ds"
	"github.com/snonux/algorithms/sort"
)

const factor int = 10
const minLength int = 1
const maxLength int = 10000

// Store results here to avoid compiler optimizations
var benchResult int

func TestElementary(t *testing.T) {
	for i := minLength; i <= maxLength; i *= factor {
		test(NewElementary(), i, t)
	}
}

func TestBST(t *testing.T) {
	for i := minLength; i <= maxLength; i *= factor {
		test(NewBST(), i, t)
	}
}

func test(s Set, l int, t *testing.T) {
	keys := ds.NewRandomArrayList(l, l)
	randoms := ds.NewRandomArrayList(l, -1)
	mapping := make(map[int]int, l)

	get := func(key int, del bool) int {
		var val int
		var err error
		switch {
		case del:
			defer delete(mapping, key)
			val, err = s.Del(key)
			//t.Log("Del", key, val, err)
		default:
			val, err = s.Get(key)
			//t.Log("Get", key, val, err)
		}

		if mVal, ok := mapping[key]; ok {
			if err != nil {
				t.Error("Could not get element", key, val, mVal, err)
			}
			if mVal != val {
				t.Error("Got wrong value for element", key, val, mVal)
			}
			return val
		}

		if err == nil {
			t.Error("Got element but expected not to", key, val)
		}
		return val
	}
	testGet := func(key int) int { return get(key, false) }
	testDel := func(key int) int { return get(key, true) }

	testSet := func(key, val int) {
		s.Set(key, val)
		mapping[key] = val
		//t.Log("Set", key, val)
		testGet(key)
	}

	t.Log("Set random key-values", l)
	var prevKey int
	for _, key := range sort.Shuffle(keys) {
		testSet(key, randoms[key])
		testGet(prevKey)
		prevKey = key
	}
	t.Log("Del random key-values", l)
	for _, key := range sort.Shuffle(keys) {
		testDel(key)
		testGet(prevKey)
		prevKey = key
	}
	if !s.Empty() {
		t.Error("Expected set to be empty", l)
	}
}

func TestBalancedBST(t *testing.T) {
	for i := minLength; i <= maxLength; i *= factor {
		test(NewBalancedBST(), i, t)
	}
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
