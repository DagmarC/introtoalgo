package ntree

import (
	"fmt"
	"strings"
)

type NTree[T fmt.Stringer] struct {
	root *Nnode[T]
}

type Nnode[T fmt.Stringer] struct {
	key      T
	leftCh   *Nnode[T]
	rightSib *Nnode[T]
	visited  bool // for print usage
}

func NewNTree[T fmt.Stringer]() *NTree[T] {
	return &NTree[T]{}
}

func NewNnode[T fmt.Stringer](val T) *Nnode[T] {
	return &Nnode[T]{key: val}
}

func (nt *NTree[T]) String() string {
	var buf strings.Builder

	if nt.root == nil {
		buf.WriteString("Empty N-ary tree")
		return buf.String()
	}

	buf.WriteString("N-ary tree: ")
	recPrint(nt.root, &buf)

	return buf.String()
}

func recPrint[T fmt.Stringer](n *Nnode[T], buf *strings.Builder) error {
	if n == nil {
		return nil
	}
	if !n.visited {
		_, err := buf.WriteString(n.key.String() + " ")
		if err != nil {
			return err
		}
		n.visited = true
	}

	recPrint(n.leftCh, buf) // call on each leftmost child of n

	// traverse siblings
	for sb := n.rightSib; sb != nil; sb = sb.rightSib {
		recPrint(sb, buf) // call on each sibling
	}

	return nil
}
