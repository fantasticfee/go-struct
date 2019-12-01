package singlinklist

import (
	"fmt"
	"testing"
)

var l *LinkList

func init() {
	n5 := &LinkNode{val: 5}
	n4 := &LinkNode{val: 4, next: n5}
	n3 := &LinkNode{val: 3, next: n4}
	n2 := &LinkNode{val: 2, next: n3}
	n1 := &LinkNode{val: 1, next: n2}
	l = &LinkList{head: &LinkNode{next: n1}}
}

func TestReverse(t *testing.T) {
	l.Print()
	ReverseLinklist(l)
	l.Print()
}

func TestHasCycle(t *testing.T) {
	fmt.Println(l.HasCycle())
	l.head.next.next.next.next.next.next = l.head.next.next.next
	fmt.Println(l.HasCycle())
}

func TestMergeSortedList(t *testing.T) {
	n5 := &LinkNode{val: 9}
	n4 := &LinkNode{val: 7, next: n5}
	n3 := &LinkNode{val: 5, next: n4}
	n2 := &LinkNode{val: 3, next: n3}
	n1 := &LinkNode{val: 1, next: n2}
	l1 := &LinkList{head: &LinkNode{next: n1}}

	n10 := &LinkNode{val: 10}
	n9 := &LinkNode{val: 8, next: n10}
	n8 := &LinkNode{val: 6, next: n9}
	n7 := &LinkNode{val: 4, next: n8}
	n6 := &LinkNode{val: 2, next: n7}
	l2 := &LinkList{head: &LinkNode{next: n6}}

	MergeSortedList(l1, l2).Print()
}

func TestDeleteBottomN(t *testing.T) {
	l.Print()
	l.DeleteBottomN(3)
	l.Print()
}

func TestFindMiddleNode(t *testing.T) {
	l.DeleteBottomN(1)
	l.DeleteBottomN(1)
	l.DeleteBottomN(1)
	l.DeleteBottomN(1)
	l.Print()
	t.Log(l.FindMiddleNode())
}
