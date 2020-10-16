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

func (t *Tree) set(n *node, key, val int) {
	switch {
	case key < n.key:
		if n.left == nil {
			n.left = &node{key, val, nil, nil}
			return
		}
		t.set(n.left, key, val)

	case key > n.key:
		if n.right == nil {
			n.right = &node{key, val, nil, nil}
			return
		}
		t.set(n.right, key, val)
	default:
		// Val is already in the tree
	}
}

func (t *Tree) Set(key, val int) {
	if t.Empty() {
		t.root = &node{key, val, nil, nil}
		return
	}
	t.set(t.root, key, val)
}

func (t *Tree) get(n *node, key int) (int, error) {
	if n == nil {
		return 0, NotFound
	}

	switch {
	case key < n.key:
		return t.get(n.left, key)
	case key > n.key:
		return t.get(n.right, key)
	default:
		return n.val, nil
	}
}

func (t *Tree) Get(key int) (int, error) {
	return t.get(t.root, key)
}

/*
	Get(key int) (int, error)
	Del(key int) (int, error)
*/
