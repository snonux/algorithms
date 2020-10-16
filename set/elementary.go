package set

type ElementaryElem struct {
	key  int
	val  int
	next *ElementaryElem
}

type Elementary struct {
	root *ElementaryElem
}

func NewElementary() *Elementary {
	return &Elementary{}
}

func (s *Elementary) Empty() bool {
	return s.root == nil
}

func (s *Elementary) Set(key int, val int) {
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

func (s *Elementary) Get(key int) (int, error) {
	elem := s.root

	for elem != nil {
		if elem.key == key {
			return elem.val, nil
		}
		elem = elem.next
	}

	return 0, NotFound
}

func (s *Elementary) Del(key int) (int, error) {
	if s.Empty() {
		return 0, NotFound
	}

	if s.root.key == key {
		defer func() {
			s.root = s.root.next
		}()
		return s.root.val, nil
	}

	elem := s.root
	for elem.next != nil {
		if elem.next.key == key {
			defer func() {
				elem.next = elem.next.next
			}()
			return elem.next.val, nil
		}
		elem = elem.next
	}

	return 0, NotFound
}
