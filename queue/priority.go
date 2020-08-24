package queue

type PriorityQueue interface {
	Insert(a int)
	Max() (max int)
	DeleteMax() int
	Empty() bool
	Size() int
	Clear()
}
