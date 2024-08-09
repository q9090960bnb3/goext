package linklist

import (
	"reflect"
	"slices"
	"testing"
)

func TestNewLinkNode(t *testing.T) {
	type args struct {
		vals []int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "valid",
			args: args{
				vals: []int{1, 3, 7, 9, 2, 6, 4},
			},
			want: NewLinkNode(1, 3, 7, 9, 2, 6, 4),
		},
		{
			name: "nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLinkNode(tt.args.vals...); !Equal(got, tt.want) {
				t.Errorf("NewLinkNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_Print(t *testing.T) {
	arr := []int{1, 3, 7, 9, 2, 6, 4}

	node := NewLinkNode(arr...)

	PrintNode(node)
}

func TestInsertNode(t *testing.T) {
	arr := []int{1, 3, 7, 9, 2, 6, 4}

	node := NewLinkNode(arr...)

	PrintNode(node)

	InsertNode(node, 5, 8)

	PrintNode(node)
}

func TestEqual(t *testing.T) {
	type args struct {
		n1 *Node
		n2 *Node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "equal",
			args: args{
				n1: NewLinkNode(1, 2, 3),
				n2: NewLinkNode(1, 2, 3),
			},
			want: true,
		},
		{
			name: "no equal",
			args: args{
				n1: NewLinkNode(1, 2, 3),
				n2: NewLinkNode(1, 2),
			},
		},
		{
			name: "no equal value",
			args: args{
				n1: NewLinkNode(1, 2, 3),
				n2: NewLinkNode(1, 2, 9),
			},
		},
		{
			name: "nil equal",
			want: true,
		},
		{
			name: "nil n2 no equal",
			args: args{
				n1: NewLinkNode(1, 2, 3),
			},
		},
		{
			name: "nil n1 no equal",
			args: args{
				n2: NewLinkNode(1, 2, 3),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.args.n1, tt.args.n2); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetIndexNode(t *testing.T) {
	type args struct {
		head  *Node
		index int
	}
	tests := []struct {
		name     string
		args     args
		wantPre  *Node
		wantNode *Node
	}{
		// TODO: Add test cases.
		{
			name: "valid",
			args: args{
				head:  NewLinkNode(1, 3, 7, 9, 2, 6, 4),
				index: 3,
			},
			wantPre: &Node{
				Val: 7,
			},
			wantNode: &Node{
				Val: 9,
			},
		},
		{
			name: "invalid",
			args: args{
				head:  NewLinkNode(1, 3, 7, 9, 2, 6, 4),
				index: 10,
			},
			wantPre:  nil,
			wantNode: nil,
		},
		{
			name: "head",
			args: args{
				head:  NewLinkNode(1, 3, 7, 9, 2, 6, 4),
				index: 0,
			},
			wantPre: nil,
			wantNode: &Node{
				Val: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPre, gotNode := GetIndexNode(tt.args.head, tt.args.index)
			if gotPre != nil && tt.wantPre != nil {
				if tt.wantPre.Val != gotPre.Val {
					t.Errorf("GetIndexNode() gotPreVal = %v, want %v", gotPre.Val, tt.wantPre.Val)
				}
			} else {
				if gotPre != tt.wantPre {
					t.Errorf("GetIndexNode() gotPre = %v, want %v", gotPre, tt.wantPre)
				}
			}

			if gotNode != nil && tt.wantNode != nil {
				if tt.wantNode.Val != gotNode.Val {
					t.Errorf("GetIndexNode() gotNodeVal = %v, want %v", gotNode.Val, tt.wantNode.Val)
				}
			} else {
				if gotNode != tt.wantNode {
					t.Errorf("GetIndexNode() gotNode = %v, want %v", gotNode, tt.wantNode)
				}
			}
		})
	}
}

func TestRemoveNext(t *testing.T) {

	type args struct {
		node *Node
	}
	tests := []struct {
		name           string
		args           args
		wantDeletedVal int
		wantOk         bool
	}{
		{
			name: "remove valid node",
			args: args{
				node: func() *Node {
					head := NewLinkNode(1, 3, 7, 9, 2, 6, 4)
					_, n := GetIndexNode(head, 3)
					return n
				}(),
			},
			wantDeletedVal: 2,
			wantOk:         true,
		},
		{
			name: "remove last next node",
			args: args{
				node: func() *Node {
					head := NewLinkNode(1, 3, 7, 9, 2, 6, 4)
					_, n := GetIndexNode(head, 6)
					return n
				}(),
			},
			wantDeletedVal: 0,
			wantOk:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDeletedVal, gotOk := RemoveNext(tt.args.node)
			if gotDeletedVal != tt.wantDeletedVal {
				t.Errorf("RemoveNext() gotDeletedVal = %v, want %v", gotDeletedVal, tt.wantDeletedVal)
			}
			if gotOk != tt.wantOk {
				t.Errorf("RemoveNext() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestRemoveFromIndex(t *testing.T) {
	type args struct {
		node  *Node
		index int
	}
	tests := []struct {
		name           string
		args           args
		wantHead       *Node
		wantDeletedVal int
		wantOk         bool
	}{
		{
			name: "remove valid node",
			args: args{
				node:  NewLinkNode(1, 3, 7, 9, 2, 6, 4),
				index: 3,
			},
			wantHead:       NewLinkNode(1, 3, 7, 2, 6, 4),
			wantDeletedVal: 9,
			wantOk:         true,
		},
		{
			name: "remove head node",
			args: args{
				node:  NewLinkNode(1, 3, 7, 9, 2, 6, 4),
				index: 0,
			},
			wantHead:       NewLinkNode(3, 7, 9, 2, 6, 4),
			wantDeletedVal: 1,
			wantOk:         true,
		},
		{
			name: "remove invalid node",
			args: args{
				node:  NewLinkNode(1, 3, 7, 9, 2, 6, 4),
				index: 18,
			},
			wantOk: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHead, gotDeletedVal, gotOk := RemoveFromIndex(tt.args.node, tt.args.index)
			if !reflect.DeepEqual(gotHead, tt.wantHead) {
				t.Errorf("RemoveFromIndex() gotHead = %v, want %v", gotHead, tt.wantHead)
			}
			if gotDeletedVal != tt.wantDeletedVal {
				t.Errorf("RemoveFromIndex() gotDeletedVal = %v, want %v", gotDeletedVal, tt.wantDeletedVal)
			}
			if gotOk != tt.wantOk {
				t.Errorf("RemoveFromIndex() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestFindNode(t *testing.T) {
	type args struct {
		n   *Node
		val int
	}
	tests := []struct {
		name      string
		args      args
		wantIndex int
	}{
		{
			name: "find valid node",
			args: args{
				n:   NewLinkNode(1, 3, 7, 9, 2, 6, 4),
				val: 3,
			},
			wantIndex: 1,
		},
		{
			name: "find invalid node",
			args: args{
				n:   NewLinkNode(1, 3, 7, 9, 2, 6, 4),
				val: 10,
			},
			wantIndex: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIndex := FindNode(tt.args.n, tt.args.val); gotIndex != tt.wantIndex {
				t.Errorf("FindNode() = %v, want %v", gotIndex, tt.wantIndex)
			}
		})
	}
}

func TestFindNodes(t *testing.T) {
	type args struct {
		n   *Node
		val int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "find valid node",
			args: args{
				n:   NewLinkNode(1, 3, 7, 9, 3, 6, 4),
				val: 3,
			},
			want: []int{1, 4},
		},
		{
			name: "find invalid node",
			args: args{
				n:   NewLinkNode(1, 3, 7, 9, 2, 6, 4),
				val: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindNodes(tt.args.n, tt.args.val); !slices.Equal(got, tt.want) {
				t.Errorf("FindNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
