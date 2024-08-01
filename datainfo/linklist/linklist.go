package linklist

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func NewLinkNode(vals ...int) *Node {
	if len(vals) == 0 {
		return nil
	}

	res := &Node{Val: vals[0]}

	res.InsertNode(vals[1:]...)

	return res
}

func (n *Node) InsertNode(vals ...int) *Node {

	node := n
	for _, v := range vals {
		node.Next = &Node{Val: v}
		node = node.Next
	}

	return node
}

func (n *Node) Print() {
	cur := n
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}
