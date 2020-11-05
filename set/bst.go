package set

type node struct {
	key   int
	val   int
	left  *node
	right *node
}

type BST struct {
	root *node
}

func NewBST() *BST {
	return &BST{}
}

func (t *BST) Empty() bool {
	return t.root == nil
}

func (t *BST) Set(key, val int) {
	if t.Empty() {
		t.root = &node{key, val, nil, nil}
		return
	}
	ptr, _, err := t.search(&t.root, key)
	switch err {
	case nil:
		// key already in the tree
		return
	case NotFound:
		*ptr = &node{key, val, nil, nil}
		return
	default:
		panic(err)
	}
}

func (t *BST) Get(key int) (int, error) {
	_, n, err := t.search(&t.root, key)
	return n.val, err
}

func (t *BST) Del(key int) (int, error) {
	ptr, n, err := t.search(&t.root, key)
	if err != nil {
		return 0, err
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
		// I have two children!

		o, err := t.deleteMin(&n.right)
		if err != nil {
			return 0, err
		}

		o.left = n.left
		o.right = n.right
		*ptr = o

		return n.val, nil
	}
}

func (t *BST) search(ptr **node, key int) (**node, *node, error) {
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

func (t *BST) deleteMin(ptr **node) (*node, error) {
	ptr, n, err := t.min(ptr)
	if err != nil {
		return nil, err
	}

	*ptr = n.right
	n.right = nil
	return n, nil
}

func (t *BST) min(ptr **node) (**node, *node, error) {
	n := *ptr
	if n == nil {
		return nil, nil, NotFound
	}

	for {
		if n.left == nil {
			return ptr, n, nil
		}
		ptr = &n.left
		n = *ptr
	}
}

/*
func (t *BST) max(ptr **node) (**node, *node, error) {
	n := *ptr
	if n == nil {
		return nil, nil, NotFound
	}

	for {
		if n.right == nil {
			return ptr, n, nil
		}
		ptr = &n.right
		n = *ptr
	}
}
*/
