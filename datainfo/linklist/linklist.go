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

func PrintNode(n *Node) {
	n.Print()
}

func (n *Node) Print() {
	cur := n
	for cur != nil {
		fmt.Print(cur.Val, " ")
		cur = cur.Next
	}
	fmt.Println()
}

/**
 * @description: 比较2个链表值是否相等
 * @param {*Node} n1
 * @param {*Node} n2
 */
func Equal(n1, n2 *Node) bool {
	if n1 == nil && n2 == nil {
		return true
	}
	if n1 == nil || n2 == nil {
		return false
	}
	return n1.Equal(n2)
}

func (n *Node) Equal(other *Node) bool {
	cur1 := n
	cur2 := other

	for cur1 != nil && cur2 != nil {
		if cur1.Val != cur2.Val {
			return false
		}

		cur1 = cur1.Next
		cur2 = cur2.Next
	}

	return cur1 == nil && cur2 == nil
}

/**
 * @description: 插入节点
 * @param {*Node} node
 * @param {...int} vals
 * @return {*Node} last 插入的最后一格节点
 */
func InsertNode(node *Node, vals ...int) (last *Node) {
	return node.InsertNode(vals...)
}

func (n *Node) InsertNode(vals ...int) *Node {

	for n.Next != nil {
		n = n.Next
	}

	for _, v := range vals {
		n.Next = &Node{Val: v}
		n = n.Next
	}

	return n
}

/**
 * @description: 获取指定索引的节点
 * @param {*Node} head
 * @param {int} index
 * @return {*Node} pre 前一个节点
 * @return {*Node} node 获取的节点
 */
func GetIndexNode(head *Node, index int) (pre, node *Node) {
	return head.GetIndexNode(index)
}

func (n *Node) GetIndexNode(index int) (pre, node *Node) {
	if index == 0 {
		return nil, n
	}

	var i int
	for n.Next != nil {
		i++
		pre = n
		n = n.Next
		if i == index {
			return pre, n
		}
	}

	return nil, nil
}

/**
 * @description: 删除节点
 * @param {*Node} node
 * @return {int} deletedVal 删除的值
 * @return {bool} ok 是否成功
 */
func RemoveNext(node *Node) (deletedVal int, ok bool) {
	return node.RemoveNext()
}

func (n *Node) RemoveNext() (int, bool) {
	if n.Next == nil {
		return 0, false
	}

	deleteNode := n.Next
	n.Next = deleteNode.Next
	deleteNode.Next = nil
	return deleteNode.Val, true
}

/**
 * @description: 删除指定索引的节点
 * @param {*Node} node
 * @param {int} index
 * @return {*Node} head 头结点
 * @return {int} deletedVal 删除的值
 * @return {bool} ok 是否成功
 */
func RemoveFromIndex(node *Node, index int) (head *Node, deletedVal int, ok bool) {
	return node.RemoveFromIndex(index)
}

func (n *Node) RemoveFromIndex(index int) (head *Node, deletedVal int, ok bool) {
	pre, node := n.GetIndexNode(index)
	if node == nil {
		return nil, 0, false
	}

	next := node.Next
	node.Next = nil
	if pre == nil {
		return next, node.Val, true
	}

	pre.Next = next
	return n, node.Val, true
}

/**
 * @description: 查找节点
 * @param {*Node} n
 * @param {int} val 查找的值
 * @return {int} index 索引， -1表示未找到
 */
func FindNode(n *Node, val int) (index int) {
	return n.FindNode(val)
}

func (n *Node) FindNode(val int) (index int) {
	cur := n
	for cur != nil {
		if cur.Val == val {
			return index
		}
		cur = cur.Next
		index++
	}
	return -1
}

/**
 * @description: 查找多个节点
 * @param {*Node} n
 * @param {int} val 查找的值
 * @return {[]int} indexes
 */
func FindNodes(n *Node, val int) (indexes []int) {
	return n.FindNodes(val)
}

func (n *Node) FindNodes(val int) (indexes []int) {
	cur := n
	index := 0
	for cur != nil {
		if cur.Val == val {
			indexes = append(indexes, index)
		}
		cur = cur.Next
		index++
	}
	return indexes
}
