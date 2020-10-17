package set

type node struct {
	key   int
	val   int
	left  *node
	right *node
}

type Tree struct {
	root *node
}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) Empty() bool {
	return t.root == nil
}

func (t *Tree) Set(key, val int) {
	if t.Empty() {
		t.root = &node{key, val, nil, nil}
		return
	}
	ptr, _, err := t.search(&t.root, key)
	if err != nil {
		// key already in the tree as node found
		return
	}

	*ptr = &node{key, val, nil, nil}
}

func (t *Tree) Get(key int) (int, error) {
	_, n, err := t.search(&t.root, key)
	return n.val, err
}

func (t *Tree) Del(key int) (int, error) {
	ptr, n, err := t.search(&t.root, key)
	if err != nil {
		return n.val, err
	}

	// Case 1: n is leaf node
	// Case 2: n has one child
	// Case 3: n has two childs

	switch {
	case n.left == nil:
		if n.right == nil {
			// I am a leaf node
			*ptr = nil
			return n.val, nil
		}
		// I have a right child
		*ptr = n.right
		return n.val, nil

	case n.right == nil:
		// I have a left child
		*ptr = n.left
		return n.val, nil
	default:
		// I have two childs!
		return 0, NotImplemented
	}
}

func (t *Tree) min(ptr **node) (**node, *node, error) {
	return nil, nil, NotImplemented
}

func (t *Tree) search(ptr **node, key int) (**node, *node, error) {
	n := *ptr
	if n == nil {
		return ptr, nil, NotFound
	}

	switch {
	case key < n.key:
		return t.search(&n.left, key)
	case n.key < key:
		return t.search(&n.right, key)
	default:
		return ptr, n, nil
	}
}
