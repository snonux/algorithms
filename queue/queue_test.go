package queue

import (
	"algorithms/ds"
	"fmt"
	"testing"
)

const minLength int = 1
const maxLength int = 10
const factor int = 10

// Store results here to avoid compiler optimizations
var benchResult ds.ArrayList

func TestElementaryPriority(t *testing.T) {
	q := NewElementaryPriority(1)
	for i := minLength; i <= maxLength; i *= factor {
		test(q, i, t)
	}
}

func TestHeapPriority(t *testing.T) {
	q := NewHeapPriority(1)
	for i := minLength; i <= maxLength; i *= factor {
		test(q, i, t)
	}
}

func test(q PriorityQueue, l int, t *testing.T) {
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

func BenchmarkElementaryPriority(b *testing.B) {
	q := NewElementaryPriority(1)
	for i := minLength; i <= maxLength; i *= factor {
		benchmark(q, i, b)
	}
}

func BenchmarkHeapPriority(b *testing.B) {
	q := NewHeapPriority(1)
	for i := minLength; i <= maxLength; i *= factor {
		benchmark(q, i, b)
	}
}

func benchmark(q PriorityQueue, l int, b *testing.B) {
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
