package avltree

type AvlTree struct {
	root *AvlNode
	size int
}

type AvlNode struct {
	element	Comparable
	rank	uint8 // height in AVL trees, not exactly height in RAVL/WAVL trees
	parent 	*AvlNode
	left   	*AvlNode
	right  	*AvlNode
}

type Comparable interface {
	CompareTo(o Comparable) int // A la .net IComparable & java.lang.Comparable
}

func NewAvlTree() *AvlTree {
	tree := new(AvlTree)
	tree.size = 0
	tree.root = nil
	return tree
}

func (tree *AvlTree) Size() int {
	return tree.size
}

func (tree *AvlTree) Find(e Comparable) *AvlNode {
	for node := tree.root; node != nil; {
		comparison := e.CompareTo(node.element)
		if comparison < 0 {
			node = node.left
		} else if comparison > 0 {
			node = node.right
		} else {
			return node
		}
	}
	return nil
}

func (tree *AvlTree) Insert(e Comparable) {
	if tree.root == nil {
		// todo assert element is valid for later compares
		tree.root = &AvlNode{e, 0, nil, nil, nil}
		tree.size = 1
		return
	}

	var parent *AvlNode
	var comparison int
	for x := tree.root; x != nil; {
		parent = x
		comparison = e.CompareTo(x.element)
		if comparison < 0 {
			x = x.left
		} else if comparison > 0 {
			x = x.right
		} else {
			x.element = e
			return
		}
	}

	node := &AvlNode{e, 0, parent, nil, nil}

	if comparison < 0 {
		parent.left = node
	} else {
		parent.right = node
	}

	if parent.rank == 0 {
		parent.rank++
		tree.retrace(parent)
	}

	tree.size++
}

func (tree *AvlTree) retrace(x *AvlNode) {
	for parent := x.parent;
		parent != nil && x.rank + 1 != parent.rank; x.rank++ {
		if parent.left == x { // new node added on the left
			if needToRotateRight(parent) {
				if x.left == nil || x.rank >= x.left.rank + 2 {
					x.rank--
					x.right.rank++
					tree.rotateLeft(x)
				}
				parent.rank--
				tree.rotateRight(parent)
				break
			}
		} else {
			if needToRotateLeft(parent) {
				if x.right == nil || x.rank >= x.right.rank + 2 {
					x.rank--
					x.left.rank++
					tree.rotateRight(x)
				}
				parent.rank--
				tree.rotateLeft(parent)
				break
			}
		}
		x = parent
		parent = x.parent
	}
}

func needToRotateRight(p *AvlNode) bool{
	if p.right == nil {
		if p.rank == 1 {
			return true
		}
		return false
	} else if p.rank >= p.right.rank + 2 {
		return true
	}
	return false
}

func needToRotateLeft(p *AvlNode) bool {
	if p.left == nil {
		if p.rank == 1 {
			return true
		}
		return false
	} else if p.rank >= p.left.rank + 2 {
		return true
	}
	return false
}

/** from CLR */
func (tree *AvlTree)rotateRight(p *AvlNode) {
	l := p.left
	p.left = l.right
	if l.right != nil {
		l.right.parent = p
	}
	l.parent = p.parent
	if p.parent == nil {
		tree.root = l
	} else if p.parent.right == p {
		p.parent.right = l
	} else {
		p.parent.left = l
	}
	l.right = p
	p.parent = l
}

/** from CLR */
func (tree *AvlTree)rotateLeft(p *AvlNode) {
	r := p.right
	p.right = r.left
	if r.left != nil {
		r.left.parent = p
	}
	r.parent = p.parent
	if p.parent == nil {
		tree.root = r
	} else if p.parent.left == p {
		p.parent.left = r
	} else {
		p.parent.right = r
	}
	r.left = p
	p.parent = r
}

func (tree *AvlTree) Delete(e Comparable) bool {
	if tree.root == nil {
		return false
	}

	node := tree.Find(e)
	if node == nil {
		return false
	}

	// if node has 2 children we need the successor or predecessor node
	if node.left != nil && node.right != nil {
		s := tree.Successor(node)
		node.element = s.element  // overwrite node's value with successor's
		node = s // then setup to delete the successor
	}

	replacement := node.right
	if node.left != nil {
		replacement = node.left
	}

	if replacement != nil {
		replacement.parent = node.parent
		var sibling *AvlNode

		if node.parent == nil {
			tree.root = replacement
			return true
		} else if node == node.parent.left {
			node.parent.left = replacement
			sibling = node.parent.right
		} else {
			node.parent.right = replacement
			sibling = node.parent.left
		}

		node.left = nil; node.right = nil; node.parent = nil

		tree.retraceDelete(replacement.parent, sibling, replacement)
	} else if node.parent == nil {
		tree.root = nil
	} else { // no children
		parent := node.parent
		var sibling *AvlNode

		if parent != nil {
			if node == parent.left {
				node.parent.left = nil
				sibling = parent.right
			} else if node == parent.right {
				node.parent.right = nil
				sibling = parent.left
			}
			node.parent = nil
		}

		tree.retraceDelete(node, sibling, parent)
	}

	tree.size--
	return true
}

func (tree *AvlTree) retraceDelete(parent *AvlNode, sibling *AvlNode, node *AvlNode) {
	var balance int
	if sibling == nil {
		balance = - 1 - int(node.rank)
	} else {
		balance = int(sibling.rank) - int(node.rank)
	}

	for ; balance != 1 ; balance = int(sibling.rank) - int(node.rank) {
		if balance == 0 {
			parent.rank--;
		} else if parent.left == sibling {
			parent.rank -= 2
			siblingBalance := rank(sibling.right) - rank(sibling.left)
			if siblingBalance == 0 {
				sibling.rank++
				parent.rank++
				tree.rotateRight(parent)
				break
			} else if siblingBalance > 0 {
				sibling.right.rank++
				sibling.rank--
				tree.rotateLeft(sibling)
			}
			tree.rotateRight(parent)
			parent = parent.parent
		} else {
			parent.rank -= 2
			siblingBalance := rank(sibling.right) - rank(sibling.left)
			if siblingBalance == 0 {
				sibling.rank++
				parent.rank++
				tree.rotateLeft(parent)
				break
			} else if siblingBalance < 0 {
				sibling.left.rank++
				sibling.rank--
				tree.rotateRight(sibling)
			}
			tree.rotateLeft(parent)
			parent = parent.parent
		}

		if parent.parent == nil {
			return
		}
		node = parent
		parent = parent.parent
		if parent.left == node {
			sibling = parent.right
		} else {
			sibling = parent.left
		}
	}
}

func rank(node *AvlNode) int {
	if node == nil {
		return -1
	} else {
		return int(node.rank)
	}
}

// Next larger element if it exists
func (tree *AvlTree) Successor(node *AvlNode) *AvlNode {
	if node == nil {
		return nil
	} else if node.right != nil {
		for node = node.right; node.left != nil; node = node.left {
		}
		return node
	} else {
		//todo
		return nil
	}
}