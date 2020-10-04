package set

type ElementaryElem struct {
	key  int
	val  interface{}
	next *ElementaryElem
}

type Elementary struct {
	root *ElementaryElem
}

func NewElementary() *Elementary {
	return &Elementary{}
}

func (s Elementary) Empty() bool {
	return s.root == nil
}

func (s Elementary) Set(key int, val interface{}) {
	if s.Empty() {
		s.root = &ElementaryElem{key, val, nil}
		return
	}

	elem := s.root

	for {
		if elem.key == key {
			elem.val = val
			return
		}
		if elem.next == nil {
			elem.next = &ElementaryElem{key, val, nil}
			return
		}
		elem = elem.next
	}
}

func (s Elementary) Get(key int) interface{} {
	elem := s.root

	for elem != nil {
		if elem.key == key {
			return elem.val
		}
		elem = elem.next
	}

	return nil
}

func (s Elementary) Del(key int) interface{} {
	if s.Empty() {
		return nil
	}

	if s.root.key == key {
		defer func() { s.root = nil }()
		return s.root.val
	}

	elem := s.root

	for elem.next != nil {
		if elem.next.key == key {
			defer func() { elem.next = elem.next.next }()
			return elem.next.val
		}
		elem = elem.next
	}

	return nil
}
