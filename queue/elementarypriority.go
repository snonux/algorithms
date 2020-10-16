package queue

import (
	"github.com/snonux/algorithms/ds"
)

type ElementaryPriority struct {
	a ds.ArrayList
	// Initial capacity
	capacity int
}

func NewElementaryPriority(capacity int) *ElementaryPriority {
	return &ElementaryPriority{make(ds.ArrayList, 0, capacity), capacity}
}

func (q *ElementaryPriority) Insert(a int) {
	q.a = append(q.a, a)
}

func (q *ElementaryPriority) Max() (max int) {
	_, max = q.max()
	return
}

func (q *ElementaryPriority) DeleteMax() int {
	if q.Empty() {
		return 0
	}

	ind, max := q.max()
	for i := ind + 1; i < q.Size(); i++ {
		q.a[i-1] = q.a[i]
	}
	q.a = q.a[0 : len(q.a)-1]

	return max
}

func (q *ElementaryPriority) Empty() bool {
	return q.Size() == 0
}

func (q *ElementaryPriority) Size() int {
	return len(q.a)
}

func (q *ElementaryPriority) Clear() {
	q.a = make(ds.ArrayList, 0, q.capacity)
}

func (q *ElementaryPriority) max() (ind, max int) {
	for i, a := range q.a {
		if a > max {
			ind, max = i, a
		}
	}
	return ind, max
}
