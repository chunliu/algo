package tree

import (
	"testing"
)

type P struct {
	S []string
}

func (p *P) ProcessNode(node *BinaryNode) {
	p.S = append(p.S, node.Value.(string))
}

func CreateTree() *BinaryNode {
	tree := &BinaryNode{
		Value: "D",
	}
	tree.LeftChild = &BinaryNode{Value: "B"}
	tree.RightChild = &BinaryNode{Value: "E"}
	tree.LeftChild.LeftChild = &BinaryNode{Value: "A"}
	tree.LeftChild.RightChild = &BinaryNode{Value: "C"}

	return tree
}

func TestTraversalPreorder(t *testing.T) {
	p := &P{}
	tree := CreateTree()

	TraversePreorder(tree, p)
	expected := []string{"D", "B", "A", "C", "E"}
	if !strSliceEqual(expected, p.S) {
		t.Errorf("Failed. Expected: %v. Get: %v", expected, p.S)
	}
}

func strSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestTraverseInorder(t *testing.T) {
	p := &P{}
	tree := CreateTree()

	TraverseInorder(tree, p)
	expected := []string{"A", "B", "C", "D", "E"}
	if !strSliceEqual(expected, p.S) {
		t.Errorf("Failed. Expected: %v. Get: %v", expected, p.S)
	}
}

func TestTraversePostorder(t *testing.T) {
	p := &P{}
	tree := CreateTree()

	TraversePostorder(tree, p)
	expected := []string{"A", "C", "B", "E", "D"}
	if !strSliceEqual(expected, p.S) {
		t.Errorf("Failed. Expected: %v. Get: %v", expected, p.S)
	}
}

func TestTraverseDepthFirst(t *testing.T) {
	p := &P{}
	tree := CreateTree()

	TraverseDepthFirst(tree, p)
	expected := []string{"D", "B", "E", "A", "C"}
	if !strSliceEqual(expected, p.S) {
		t.Errorf("Failed. Expected: %v. Get: %v", expected, p.S)
	}
}
