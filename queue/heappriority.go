package queue

import (
	"github.com/snonux/algorithms/ds"
)

type HeapPriority struct {
	ElementaryPriority
}

func NewHeapPriority(capacity int) *HeapPriority {
	q := HeapPriority{
		ElementaryPriority: ElementaryPriority{make(ds.ArrayList, 0, capacity), capacity},
	}

	// Index 0 unused
	q.a = append(q.a, 0)
	return &q
}

func (q *HeapPriority) Empty() bool {
	return q.Size() == 0
}

func (q *HeapPriority) Size() int {
	return len(q.a) - 1
}

func (q *HeapPriority) Insert(a int) {
	q.a = append(q.a, a)
	q.swim(len(q.a) - 1)
}

func (q *HeapPriority) Max() (max int) {
	if q.Empty() {
		return 0
	}

	return q.a[1]
}

func (q *HeapPriority) DeleteMax() int {
	switch q.Size() {
	case 0:
		return 0
	case 1:
		max := q.a[1]
		q.a = q.a[0:1]
		return max
	default:
		a := q.a[1]
		// Last index
		max := len(q.a) - 1

		q.a.Swap(1, max)
		q.a = q.a[0:max]
		q.sink(1)

		return a
	}
}

func (q *HeapPriority) swim(k int) {
	for k > 1 && q.a[k/2] < q.a[k] {
		q.a.Swap(k/2, k)
		k = k / 2
	}
}

func (q *HeapPriority) sink(k int) {
	// Last index
	max := len(q.a) - 1

	for 2*k <= max {
		// Left child
		j := 2 * k

		// Go to right child?
		if j < max && q.a[j] < q.a[j+1] {
			j++
		}

		// Found my spot in the heap
		if q.a[k] >= q.a[j] {
			break
		}

		q.a.Swap(k, j)
		k = j
	}
}
