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
	color linkColor
	size  int
	left  *rbNode
	right *rbNode
}

func (n *rbNode) String() string {
	recurse := func(n *rbNode) string {
		if n == nil {
			return ""
		}
		return n.String()
	}

	return fmt.Sprintf("rbNode{%d:%d,%v,%d,%s,%s}",
		n.key,
		n.val,
		n.color,
		n.size,
		recurse(n.left),
		recurse(n.right),
	)
}

func (n *rbNode) isRed() bool {
	if n == nil {
		return false
	}
	return n.color == red
}

func (n *rbNode) Size() int {
	if n == nil {
		return 0
	}
	return n.size
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

func (t *RedBlackBST) Size() int {
	if t.Empty() {
		return 0
	}
	return t.root.size
}

func (t *RedBlackBST) Put(key, val int) {
	t.root = t.put(t.root, key, val)
	t.root.color = black
}

func (t *RedBlackBST) put(n *rbNode, key, val int) *rbNode {
	if n == nil {
		return &rbNode{key, val, red, 1, nil, nil}
	}

	switch {
	case key < n.key:
		n.left = t.put(n.left, key, val)
	case key > n.key:
		n.right = t.put(n.right, key, val)
	default:
		n.val = val
	}

	switch {
	case n.right.isRed() && !n.left.isRed():
		n = t.rotateLeft(n)
	case n.left != nil && n.left.isRed() && n.left.left.isRed():
		n = t.rotateRight(n)
	case n.left.isRed() && n.right.isRed():
		t.flipColors(n)
	}

	n.size = 1 + n.left.Size() + n.right.Size()
	return n
}

func (t *RedBlackBST) Get(key int) (int, error) {
	return t.get(t.root, key)
}

func (t *RedBlackBST) get(n *rbNode, key int) (int, error) {
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

func (t *RedBlackBST) Del(key int) (int, error) {
	panic("Not yet implemented")
}

func (t *RedBlackBST) rotateLeft(n *rbNode) *rbNode {
	x := n.right
	n.right = x.left
	x.left = n

	x.color = n.color
	n.color = red

	x.size = n.size
	n.size = 1 + n.left.Size() + n.right.Size()

	return x
}

func (t *RedBlackBST) rotateRight(n *rbNode) *rbNode {
	x := n.left
	n.left = x.right
	x.right = n

	x.color = n.color
	n.color = red

	x.size = n.size
	n.size = 1 + n.left.Size() + n.right.Size()

	return x
}

func (t *RedBlackBST) flipColors(n *rbNode) {
	n.color = red
	n.left.color = black
	n.right.color = black
}
