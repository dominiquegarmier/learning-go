package main

type node struct {
	value int
	left  *node
	right *node
}

type btree struct {
	root *node
}

func (r *node) _invert() {
	r.left, r.right = r.right, r.left
	if r.left != nil {
		r.left._invert()
	}
	if r.right != nil {
		r.right._invert()
	}
}

func (r *btree) invert() {
	r.root._invert()
}

func main() {
	b := btree{}
	b.root = &node{1, nil, nil}
	b.root.left = &node{3, nil, nil}
	b.root.right = &node{2, nil, nil}

	b.invert()
}
