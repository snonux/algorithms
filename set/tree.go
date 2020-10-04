package tree

type Tree interface {
	Empty() bool
	Insert(key int, val interface{})
	Search(key int) interface{}
}
