package gen

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/DagmarC/introtoalgo/chapter10/stack"
)

type BinTree struct {
	root *Node
}

type Node struct {
	key    int
	left   *Node
	right  *Node
	parent *Node
}

// ============= BINARY TREE FN =============

func NewBinTree() *BinTree {
	return &BinTree{}
}

func (bt *BinTree) String() string {
	var buf strings.Builder
	buf.WriteString("Binary Tree: ")

	if bt.root == nil {
		buf.WriteString("Empty\n")
		return buf.String()
	}

	recInorder(bt.root, &buf)

	return buf.String()
}

func recInorder(n *Node, buf *strings.Builder) {
	if n == nil {
		return
	}
	recInorder(n.left, buf)
	buf.Write([]byte(strconv.Itoa(n.key)))
	buf.WriteString(" ")
	recInorder(n.right, buf)
}

// =============NODE FN =============

func NewNode(k int) *Node {
	return &Node{key: k}
}

func (n *Node) Left() (*Node, error) {
	if n == nil {
		return nil, errors.New("node n is nil")
	}
	return n.left, nil
}

func (n *Node) Right() (*Node, error) {
	if n == nil {
		return nil, errors.New("node n is nil")
	}
	return n.right, nil
}

func (n *Node) Parent() (*Node, error) {
	if n == nil {
		return nil, errors.New("node n is nil")
	}
	return n.parent, nil
}

func (n *Node) String() string {
	return strconv.Itoa(n.key)
}

// ============= non-recursive BT Print via stack =============

func (bt *BinTree) printBTstack() error {
	var err error
	n := bt.root
	if n == nil {
		return errors.New("binary tree is empty")
	}
	s := stack.Stack[*Node]{} // generics stack, if type (*Node) implements String() string method then it can be used

	for {
		s.Push(n)
		n = n.left
		if n == nil {
			n, err = s.Pull(&Node{})
			if err != nil {
				return nil // Quit, Stack is already empty
			}
			fmt.Printf("%d ", n.key)

			// Pull from stack until there exists n.right child
			for n.right == nil {
				n, err = s.Pull(&Node{})
				if err != nil {
					return nil // Quit, Stack is already empty
				}
				fmt.Printf("%d ", n.key)
			}
			n = n.right // n.right is not nil anymore
		}
	}
}
