package queue

import "algorithms/ds"

type ElementaryPriority struct {
	a    ds.ArrayList
	size int
	// Initial capacity
	capacity int
}

func NewElementaryPriority(capacity int) *ElementaryPriority {
	return &ElementaryPriority{make(ds.ArrayList, 0, capacity), 0, capacity}
}

func (q *ElementaryPriority) Insert(a int) {
	q.a = append(q.a, a)
	q.size++
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
	q.size--

	return max
}

func (q *ElementaryPriority) Empty() bool {
	return q.Size() == 0
}

func (q *ElementaryPriority) Size() int {
	return q.size
}

func (q *ElementaryPriority) Clear() {
	q.size = 0
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
