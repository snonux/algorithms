package queue

import "algorithms/ds"

type ElementaryPQ struct {
	a    ds.ArrayList
	size int
}

func NewElementaryPQ(capacity int) ElementaryPQ {
	return ElementaryPQ{make(ds.ArrayList, 0, capacity), 0}
}

func (q ElementaryPQ) Insert(a int) {
	q.a = append(q.a, a)
	q.size++
}

func (q ElementaryPQ) Max() (max int) {
	_, max = q.max()
	return
}

func (q ElementaryPQ) DeleteMax() int {
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

func (q ElementaryPQ) Empty() bool {
	return q.Size() == 0
}

func (q ElementaryPQ) Size() int {
	return q.size
}

func (q ElementaryPQ) max() (ind, max int) {
	for i, a := range q.a {
		if a > max {
			ind, max = i, a
		}
	}
	return ind, max
}
