package queue

import (
	"algorithms/ds"
	"fmt"
	"testing"
)

const minLength int = 1
const maxLength int = 10000
const factor int = 100

// Store results here to avoid compiler optimizations
var benchResult ds.ArrayList

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
		t.Log("Size", q.Size(), q.Empty())
		for !q.Empty() {
			next := q.DeleteMax()
			if started {
				if next > prev {
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

func BenchmarkElementaryPQ(b *testing.B) {
	q := NewElementaryPQ(1)
	for i := minLength; i <= maxLength; i *= factor {
		benchmark(q, i, b)
	}
}

func benchmark(q PQ, l int, b *testing.B) {
	benchResult = ds.NewRandomArrayList(l, -1)

	b.Run(fmt.Sprintf("randomInsert(%d)", l), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			q.Clear()
			for _, a := range benchResult {
				q.Insert(a)
			}
		}
	})

	b.Run(fmt.Sprintf("randomInsertAndDeleteMax(%d)", l), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			q.Clear()
			for _, a := range benchResult {
				q.Insert(a)
			}
			for i := 0; !q.Empty(); i++ {
				benchResult[i] = q.DeleteMax()
			}
		}
	})
}
