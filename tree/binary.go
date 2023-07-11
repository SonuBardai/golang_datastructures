package tree

import (
	"fmt"

	"github.com/SonuBardai/golang_datastructures/stackqueue"
)

type BinaryTreeNode struct {
	Val                   int
	LeftChild, RightChild *BinaryTreeNode
}

func NewBinaryTreeNode(val *int) *BinaryTreeNode {
	return &BinaryTreeNode{Val: *val, LeftChild: nil, RightChild: nil}
}

type BinaryTree struct {
	Root *BinaryTreeNode
}

func NewBinaryTree(val *int) BinaryTree {
	return BinaryTree{Root: NewBinaryTreeNode(val)}
}

func (t *BinaryTree) InsertBinaryTree(val *int) {
	if t.Root == nil {
		t.Root = NewBinaryTreeNode(val)
		return
	}
	queue := stackqueue.NewQueueVec(make([]interface{}, 0), 0, false)
	queue.EnVec(t.Root)
	for len(queue.Data) > 0 {
		current := queue.DeVec().(*BinaryTreeNode)
		if current.LeftChild == nil {
			current.LeftChild = NewBinaryTreeNode(val)
			return
		} else {
			queue.EnVec(current.LeftChild)
		}
		if current.RightChild == nil {
			current.RightChild = NewBinaryTreeNode(val)
			return
		} else {
			queue.EnVec(current.RightChild)
		}
	}
}

func (t *BinaryTree) DeleteBinaryTree() *BinaryTreeNode {
	if t.Root == nil {
		fmt.Println("Tree is empty")
		return nil
	}
	if t.Root.LeftChild == nil && t.Root.RightChild == nil {
		deleted := t.Root
		t.Root = nil
		return deleted
	}
	queue := stackqueue.NewQueueVec(make([]interface{}, 0), 0, false)
	queue.EnVec(t.Root)
	var prev *BinaryTreeNode
	for len(queue.Data) > 0 {
		current := queue.DeVec().(*BinaryTreeNode)
		if current.LeftChild != nil {
			queue.EnVec(current.LeftChild)
			prev = current
		}
		if current.RightChild != nil {
			queue.EnVec(current.RightChild)
			prev = current
		}
	}
	if prev.RightChild != nil {
		deleted := prev.RightChild
		prev.RightChild = nil
		return deleted
	}
	if prev.LeftChild != nil {
		deleted := prev.LeftChild
		prev.LeftChild = nil
		return deleted
	}
	return nil
}

type BinaryTraversalType int

const (
	PreOrder BinaryTraversalType = iota
	InOrder
	PostOrder
)

func preOrderTraversal(node *BinaryTreeNode) {
	if node == nil {
		return
	}
	fmt.Printf(" %d ", node.Val)
	preOrderTraversal(node.LeftChild)
	preOrderTraversal(node.RightChild)
}
func inOrderTraversal(node *BinaryTreeNode) {
	if node == nil {
		return
	}
	inOrderTraversal(node.LeftChild)
	fmt.Printf(" %d ", node.Val)
	inOrderTraversal(node.RightChild)
}
func postOrderTraversal(node *BinaryTreeNode) {
	if node == nil {
		return
	}
	postOrderTraversal(node.LeftChild)
	postOrderTraversal(node.RightChild)
	fmt.Printf(" %d ", node.Val)
}

func (t *BinaryTree) TraverseBinaryTree(traversalType BinaryTraversalType) {
	switch traversalType {
	case PreOrder:
		fmt.Println("PreOrder")
		preOrderTraversal(t.Root)
		fmt.Println()
		return
	case InOrder:
		fmt.Println("InOrder")
		inOrderTraversal(t.Root)
		fmt.Println()
		return
	case PostOrder:
		fmt.Println("PostOrder")
		postOrderTraversal(t.Root)
		fmt.Println()
		return
	}
}

func (t *BinaryTree) PrintBinaryTree() {
	if t.Root == nil {
		fmt.Println("Tree is empty")
		return
	}
	queue := stackqueue.NewQueueVec(make([]interface{}, 0), 0, false)
	queue.EnVec(t.Root)
	for len(queue.Data) > 0 {
		levelSize := len(queue.Data)
		for i := 0; i < levelSize; i++ {
			current := queue.DeVec().(*BinaryTreeNode)
			fmt.Printf(" %d ", current.Val)
			if current.LeftChild != nil {
				queue.EnVec(current.LeftChild)
			}
			if current.RightChild != nil {
				queue.EnVec(current.RightChild)
			}
		}
		fmt.Println()
	}
	fmt.Print("\n\n")
}

func TestBinaryTree() {
	rootValue := 1
	binaryTree := NewBinaryTree(&rootValue)
	binaryTree.PrintBinaryTree()

	items := []int{2, 3, 4, 5, 6, 7, 8}
	for _, val := range items {
		binaryTree.InsertBinaryTree(&val)
		binaryTree.PrintBinaryTree()
	}

	fmt.Println("Traversal")
	binaryTree.TraverseBinaryTree(PreOrder)
	binaryTree.TraverseBinaryTree(InOrder)
	binaryTree.TraverseBinaryTree(PostOrder)
	fmt.Print("\n\n")

	fmt.Println("Deleting")
	for i := 0; i <= len(items); i++ {
		binaryTree.DeleteBinaryTree()
		binaryTree.PrintBinaryTree()
	}
}
