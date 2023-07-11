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

func (t *BinaryTree) DeleteDeepestRightmost() {
	if t.Root == nil {
		fmt.Println("Tree is empty!")
		return
	}
	queue := stackqueue.NewQueueVec(make([]interface{}, 0), 0, false)
	queue.EnVec(t.Root)
	var current *BinaryTreeNode
	var prev *BinaryTreeNode
	for len(queue.Data) > 0 {
		current = queue.DeVec().(*BinaryTreeNode)
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
		prev.RightChild = nil
	} else {
		prev.LeftChild = nil
	}
}

func (t *BinaryTree) PrintBinaryTree() {
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
}
