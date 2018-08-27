package tree

import "github.com/chunliu/algo/basic/ds"

type BinaryNode struct {
	Value      interface{}
	LeftChild  *BinaryNode
	RightChild *BinaryNode
}

type IProcessNode interface {
	ProcessNode(node *BinaryNode)
}

/**
** Approach: Recursion
** Complexity: the 4 traverse methods all take O(n) time.
**	- Visit each node once, so time complexity is O(n)
**/

// TraversePreorder traverse the tree by taking the order: Current node -> left child -> righ child
func TraversePreorder(root *BinaryNode, p IProcessNode) {
	p.ProcessNode(root)
	if root.LeftChild != nil {
		TraversePreorder(root.LeftChild, p)
	}
	if root.RightChild != nil {
		TraversePreorder(root.RightChild, p)
	}
}

// TraverseInorder traverse the tree by taking the order: left child -> Current node -> righ child
func TraverseInorder(root *BinaryNode, p IProcessNode) {
	if root.LeftChild != nil {
		TraverseInorder(root.LeftChild, p)
	}
	p.ProcessNode(root)
	if root.RightChild != nil {
		TraverseInorder(root.RightChild, p)
	}
}

// TraversePostorder traverse the tree by taking the order: left child -> righ child -> Current node
func TraversePostorder(root *BinaryNode, p IProcessNode) {
	if root.LeftChild != nil {
		TraversePostorder(root.LeftChild, p)
	}
	if root.RightChild != nil {
		TraversePostorder(root.RightChild, p)
	}
	p.ProcessNode(root)
}

// TraverseDepthFirst traverse the tree by visiting all nodes on the same level first, and then moving to the next level.
func TraverseDepthFirst(root *BinaryNode, p IProcessNode) {
	q := ds.NewQueue()
	q.Enqueue(root)
	for q.Len() > 0 {
		n := q.Dequeue().(*BinaryNode)
		p.ProcessNode(n)
		if n.LeftChild != nil {
			q.Enqueue(n.LeftChild)
		}
		if n.RightChild != nil {
			q.Enqueue(n.RightChild)
		}
	}
}
