package sort

import (
	"algorithms/ds"
	"fmt"
	"testing"
)

func test(name string, length int, sort func([]ds.Comparer) []ds.Comparer, t *testing.T) {
	cb := func(t *testing.T) {
		t.Parallel()
		a := ds.RandomIntegers(length, length)
		a = sort(a)
		if !Sorted(a) {
			t.Errorf("Array not sorted: %v", a)
		}
	}
	t.Run(fmt.Sprintf("Test%s%d", name, length), cb)
}
