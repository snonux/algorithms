package queue

import (
	"algorithms/ds"
	"fmt"
	"testing"
)

const minLength int = 1
const maxLength int = 1000
const factor int = 10

func TestElementaryPQ(t *testing.T) {
	q := NewElementaryPQ(1)
	for i := minLength; i <= maxLength; i *= factor {
		test(q, i, t)
	}
}

func test(q PQ, l int, t *testing.T) {
	cb := func(t *testing.T) {
		t.Parallel()
		for _, a := range ds.NewRandomArrayList(l, -1) {
			q.Insert(a)
		}
		prev, started := 0, false
		for !q.Empty() {
			next := q.DeleteMax()
			if started {
				if next < prev {
					t.Errorf("Expected element '%v' to be lower than previous '%v': %v",
						next, prev, q)
				}
				prev = next
				continue
			}
			started = true
			prev = next
		}
	}
	t.Run(fmt.Sprintf("%d", l), cb)
}
