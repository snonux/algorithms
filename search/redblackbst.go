package search

import (
	"fmt"
)

type linkColor bool

const (
	red   linkColor = true
	black linkColor = false
)

type rbNode struct {
	key      int
	val      int
	color    linkColor
	capacity int
	left     *rbNode
	right    *rbNode
	// Just mark a node as deleted if deleted. Not fully implemented in lecture.
	deleted bool
}

func (n *rbNode) String() string {
	recurse := func(n *rbNode) string {
		if n == nil {
			return ""
		}
		return fmt.Sprintf("\n%s", n.String())
	}

	color := "red"
	if n.color == black {
		color = "black"
	}
	return fmt.Sprintf("rbNode{%v;%d:%d,%s,%d,%s,%s}",
		n.deleted,
		n.key,
		n.val,
		color,
		n.capacity,
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

func (n *rbNode) Capacity() int {
	if n == nil {
		return 0
	}
	return n.capacity
}

type RedBlackBST struct {
	root *rbNode
	size int
}

func NewRedBlackBST() *RedBlackBST {
	return &RedBlackBST{}
}

func (t *RedBlackBST) String() string {
	if t.Empty() {
		return fmt.Sprintf("RedBlackBST{%d:%d}", t.Size(), t.Capacity())
	}

	return fmt.Sprintf("RedBlackBST{%d:%d;%s}", t.Size(), t.Capacity(), t.root)
}

func (t *RedBlackBST) Size() int {
	return t.size
}

func (t *RedBlackBST) Empty() bool {
	return t.Size() == 0
}

func (t *RedBlackBST) Capacity() int {
	if t.root == nil {
		return 0
	}
	return t.root.capacity
}

func (t *RedBlackBST) Put(key, val int) {
	t.root = t.put(t.root, key, val)
	t.root.color = black
}

func (t *RedBlackBST) put(n *rbNode, key, val int) *rbNode {
	if n == nil {
		t.size++
		return &rbNode{key, val, red, 1, nil, nil, false}
	}

	switch {
	case key < n.key:
		n.left = t.put(n.left, key, val)
	case key > n.key:
		n.right = t.put(n.right, key, val)
	default:
		if n.deleted {
			n.deleted = false
		}
		t.size++
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

	n.capacity = 1 + n.left.Capacity() + n.right.Capacity()
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
		if n.deleted {
			return 0, NotFound
		}
		return n.val, nil
	}
}

func (t *RedBlackBST) Del(key int) (int, error) {
	return t.del(t.root, key)
}

func (t *RedBlackBST) del(n *rbNode, key int) (int, error) {
	if n == nil {
		return 0, NotFound
	}

	switch {
	case key < n.key:
		return t.del(n.left, key)
	case key > n.key:
		return t.del(n.right, key)
	default:
		if n.deleted {
			return 0, NotFound
		}
		t.size--
		n.deleted = true
		val := n.val
		n.val = -1
		return val, nil
	}
}

func (t *RedBlackBST) rotateLeft(n *rbNode) *rbNode {
	x := n.right
	n.right = x.left
	x.left = n

	x.color = n.color
	n.color = red

	x.capacity = n.capacity
	n.capacity = 1 + n.left.Capacity() + n.right.Capacity()

	return x
}

func (t *RedBlackBST) rotateRight(n *rbNode) *rbNode {
	x := n.left
	n.left = x.right
	x.right = n

	x.color = n.color
	n.color = red

	x.capacity = n.capacity
	n.capacity = 1 + n.left.Capacity() + n.right.Capacity()

	return x
}

func (t *RedBlackBST) flipColors(n *rbNode) {
	n.color = red
	n.left.color = black
	n.right.color = black
}
