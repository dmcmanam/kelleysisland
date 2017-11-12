package avltree

import (
	"fmt"
	"testing"
)

type Element struct {
	value int
}

func (a Element) CompareTo(o Comparable) int {
	e := o.(Element)
	if a.value < e.value {
		return -1
	} else if a.value == e.value {
		return 0
	} else {
		return 1
	}
}

func printInOrderAVL(root *AvlNode) {
	if root != nil {
		printInOrderAVL(root.left)
		fmt.Print(root.element, ",")
		printInOrderAVL(root.right)
	}
}

func TestInsertNoRotations(t *testing.T) {
	x := NewAvlTree()
	x.Insert(Element{5})
	x.Insert(Element{6})
	x.Insert(Element{4})
	if x.Size() != 3 {
		t.Fail()
	}
	if x.root.rank != 1 {
		t.Fail()
	}
}

func TestInsertRightRotation(t *testing.T) {
	x := NewAvlTree()
	x.Insert(Element{5})
	x.Insert(Element{4})
	x.Insert(Element{3})

	if x.root.element.CompareTo(Element{4}) != 0 {
		t.Fail()
	}
}

func TestInsertLeftRotation(t *testing.T) {
	x := NewAvlTree()
	x.Insert(Element{1})
	x.Insert(Element{2})
	x.Insert(Element{3})

	if x.root.element.CompareTo(Element{2}) != 0 {
		t.Fail()
	}
	if x.root.rank != 1 {
		t.Fail()
	}
	if x.root.right.rank != 0 {
		t.Fail()
	}
}

func TestInsertThreeLeftRotations(t *testing.T) {
	x := NewAvlTree()
	x.Insert(Element{1})
	x.Insert(Element{2})
	x.Insert(Element{3})
	x.Insert(Element{4})
	x.Insert(Element{5})
	x.Insert(Element{6})

	if x.root.element.CompareTo(Element{4}) != 0 {
		t.Fail()
	}
	if x.root.rank != 2 {
		t.Fail()
	}
	if x.root.right.rank != 1 {
		t.Fail()
	}
}

func TestInsertLeftRightRotation(t *testing.T) {
	x := NewAvlTree()
	x.Insert(Element{3})
	x.Insert(Element{1})
	x.Insert(Element{2})

	if x.root.element.CompareTo(Element{2}) != 0 {
		t.Fail()
	}
	if x.root.rank != 1 {
		t.Fail()
	}
	if x.root.right.rank != 0 {
		t.Fail()
	}
}

func TestInsertRightLeftRotation(t *testing.T) {
	x := NewAvlTree()
	x.Insert(Element{3})
	x.Insert(Element{6})
	x.Insert(Element{4})

	if x.root.element.CompareTo(Element{4}) != 0 {
		t.Fail()
	}
	if x.root.rank != 1 {
		t.Fail()
	}
	if x.root.right.rank != 0 {
		t.Fail()
	}
}

func TestInsertBuildFibonacciTree(t *testing.T) {
	x := NewAvlTree()
	x.Insert(Element{8})
	x.Insert(Element{5});x.Insert(Element{11})
	x.Insert(Element{3});x.Insert(Element{7});x.Insert(Element{10});x.Insert(Element{12})
	x.Insert(Element{2});x.Insert(Element{4});x.Insert(Element{6});x.Insert(Element{9})
	x.Insert(Element{1})

	if x.root.element.CompareTo(Element{8}) != 0 {
		t.Fail()
	}
}

func TestFind(t *testing.T) {
	x := NewAvlTree()
	x.Insert(Element{3})
	x.Insert(Element{2})
	x.Insert(Element{4})
	x.Insert(Element{1})

	n := x.Find(Element{4})

	if (n.element.CompareTo(Element{4}) != 0) {
		t.Fail()
	}
}

func TestDeleteOneRightRotation(t *testing.T) {
	x := NewAvlTree()
	x.Insert(Element{3})
	x.Insert(Element{2})
	x.Insert(Element{4})
	x.Insert(Element{1})


	x.Delete(Element{4})
}

func TestDeleteFibonacciTree(t *testing.T) {
	x := NewAvlTree()
	x.Insert(Element{8})
	x.Insert(Element{5});x.Insert(Element{11})
	x.Insert(Element{3});x.Insert(Element{7});x.Insert(Element{10});x.Insert(Element{12})
	x.Insert(Element{2});x.Insert(Element{4});x.Insert(Element{6});x.Insert(Element{9})
	x.Insert(Element{1})

	if x.root.element.CompareTo(Element{8}) != 0 {
		t.Fail()
	}

	x.Delete(Element{12})

	if x.root.element.CompareTo(Element{5}) != 0 {
		t.Error("Root element should be 5 but instead: ",  x.root.element);
	}
	if x.root.right.rank != 2 {
		t.Fail()
	}
	printInOrderAVL(x.root)
}