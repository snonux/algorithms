package search

type HashElem struct {
	key  int
	val  int
	next *HashElem
}

type Hash struct {
	root *HashElem
	size int
}

func NewHash() *Hash {
	return &Hash{}
}

func (h *Hash) Empty() bool {
	return s.root == nil
}

func (h *Hash) Size() int {
	return s.size
}

func (h *Hash) Put(key int, val int) {
	if s.Empty() {
		s.root = &HashElem{key, val, nil}
		s.size++
		return
	}

	elem := s.root

	for {
		if elem.key == key {
			elem.val = val
			return
		}
		if elem.next == nil {
			elem.next = &HashElem{key, val, nil}
			s.size++
			return
		}
		elem = elem.next
	}
}

func (h *Hash) Get(key int) (int, error) {
	elem := s.root

	for elem != nil {
		if elem.key == key {
			return elem.val, nil
		}
		elem = elem.next
	}

	return 0, NotFound
}

func (h *Hash) Del(key int) (int, error) {
	if s.Empty() {
		return 0, NotFound
	}

	if s.root.key == key {
		defer func() {
			s.root = s.root.next
		}()
		s.size--
		return s.root.val, nil
	}

	elem := s.root
	for elem.next != nil {
		if elem.next.key == key {
			defer func() {
				elem.next = elem.next.next
			}()
			s.size--
			return elem.next.val, nil
		}
		elem = elem.next
	}

	return 0, NotFound
}
