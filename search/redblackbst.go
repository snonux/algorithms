package search

import "fmt"

type linkColor bool

const (
	red   linkColor = true
	black linkColor = false
)

type rbNode struct {
	key   int
	val   int
	left  *rbNode
	right *rbNode
	color linkColor
}

func (n *rbNode) String() string {
	recurse := func(n *rbNode) string {
		if n == nil {
			return ""
		}
		return n.String()
	}

	return fmt.Sprintf("rbNode{%d:%d,%v,%s,%s}",
		n.key,
		n.val,
		n.color,
		recurse(n.left),
		recurse(n.right),
	)
}

type RedBlackBST struct {
	root *rbNode
}

func NewRedBlackBST() *RedBlackBST {
	return &RedBlackBST{}
}

func (t *RedBlackBST) String() string {
	if t.Empty() {
		return "RedBlackBST{}"
	}

	return fmt.Sprintf("RedBlackBST{%s}", t.root)
}

func (t *RedBlackBST) Empty() bool {
	return t.root == nil
}

func (t *RedBlackBST) Put(key, val int) {
	if t.Empty() {
		t.root = &rbNode{key, val, nil, nil, black}
		return
	}
	parent, ptr, _, err := t.search(nil, &t.root, key)
	switch err {
	case nil:
		// key already in the tree
		return
	case NotFound:
		*ptr = &rbNode{key, val, nil, nil, red}
		return
	default:
		panic(err)
	}

	t.balance(parent)
}

func (t *RedBlackBST) balance(n *rbNode) {
	switch {
	case n == nil, n.right == nil:
		return
	case n.right.color == black:
		return
	case n.left.color == black:
		// Left is black and right is red
	case n.left.color == red:
		// Left and right are red
	}
}

func (t *RedBlackBST) Get(key int) (int, error) {
	_, _, n, err := t.search(nil, &t.root, key)
	if err != nil {
		return 0, err
	}
	return n.val, nil
}

func (t *RedBlackBST) Del(key int) (int, error) {
	_, ptr, n, err := t.search(nil, &t.root, key)
	if err != nil {
		return 0, err
	}

	// Case 1: n is leaf rbNode
	// Case 2: n has one child
	// Case 3: n has two childs

	switch {
	case n.left == nil:
		if n.right == nil {
			// I am a leaf rbNode
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

func (t *RedBlackBST) search(parent *rbNode, ptr **rbNode, key int) (*rbNode, **rbNode, *rbNode, error) {
	n := *ptr
	if n == nil {
		return parent, ptr, nil, NotFound
	}

	switch {
	case key < n.key:
		return t.search(n, &n.left, key)
	case n.key < key:
		return t.search(n, &n.right, key)
	default:
		return parent, ptr, n, nil
	}
}

func (t *RedBlackBST) deleteMin(ptr **rbNode) (*rbNode, error) {
	ptr, n, err := t.min(ptr)
	if err != nil {
		return nil, err
	}

	*ptr = n.right
	n.right = nil
	return n, nil
}

func (t *RedBlackBST) min(ptr **rbNode) (**rbNode, *rbNode, error) {
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

func (t *RedBlackBST) rotateLeft(ptr **rbNode) {
	x := *ptr
	y := x.right

	x.right = y.left
	x.left = y

	*ptr = y
	x.color = red
	y.color = black
}

func (t *RedBlackBST) rotateRight(ptr **rbNode) {
	y := *ptr
	x := y.right

	y.left = x.right
	x.right = y

	*ptr = x
	x.color = black
	y.color = red
}

func (t *RedBlackBST) flipColors(n *rbNode) {
	n.left.color = black
	n.right.color = black
	n.color = red
}
