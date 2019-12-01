package singlinklist

import (
	"errors"
	"fmt"
	"strconv"
)

type LinkNode struct {
	val  interface{}
	next *LinkNode
}

type LinkList struct {
	head   *LinkNode
	length int
}

func NewLinkNode(v interface{}) *LinkNode {
	return &LinkNode{v, nil}
}

func (this *LinkNode) GetNext() *LinkNode {
	return this.next
}

func (this *LinkNode) GetValue() interface{} {
	return this.val
}

func NewLinkList() *LinkList {
	return &LinkList{
		head:   new(LinkNode),
		length: 0,
	}
}

func (this *LinkList) InsertAfter(p *LinkNode, v interface{}) bool {
	if p == nil {
		return false
	}

	newNode := new(LinkNode)
	newNode.val = v
	newNode.next = p.next
	p.next = newNode
	this.length++

	return true
}

func (this *LinkList) InsertBefore(p *LinkNode, v interface{}) bool {
	if p == nil || this.head.next == nil {
		return false
	}
	pre := this.head
	cur := this.head.next

	for nil == cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	newNode := new(LinkNode)
	newNode.val = v
	newNode.next = cur
	pre.next = newNode
	this.length++

	return true
}

func (this *LinkList) InsertToHead(v interface{}) bool {
	return this.InsertAfter(this.head, v)
}

func (this *LinkList) InsertToTail(v interface{}) bool {
	cur := this.head
	for cur.next != nil {
		cur = cur.next
	}
	return this.InsertAfter(cur, v)
}

func (this *LinkList) FindByIndex(n int) *LinkNode {
	if n > this.length {
		return nil
	}

	cur := this.head
	for i := 1; i <= n && nil != cur; i++ {
		cur = cur.next
	}

	return cur
}

func (this *LinkList) DeleteNode(p *LinkNode) bool {
	if nil == p || nil == this.head.next {
		return false
	}
	pre := this.head
	cur := pre.next
	for nil != cur {
		if p == cur {
			pre.next = cur.next
			cur.next = nil
			this.length--
			return true
		}
		pre = cur
		cur = cur.next
	}

	return false
}

func (this *LinkList) Print() {
	cur := this.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if cur != nil {
			format += "->"
		}
	}
	fmt.Println(format)
}

func (this *LinkList) HasCycle() bool {
	if nil == this.head || this.length <= 0 {
		return false
	}
	cur := this.head.next
	faster := this.head.next
	for nil != faster && nil != faster.next {
		cur = cur.next
		faster = faster.next.next
		if cur == faster {
			return false
		}
	}

	return false
}

func (this *LinkList) DeleteBottomN(n int) {
	if n >= this.length || nil == this.head || nil == this.head.next {
		return
	}

	deleteNode := this.head
	for i := 1; i <= n && nil != deleteNode; i++ {
		deleteNode = deleteNode.next
	}

	if nil == deleteNode {
		return
	}

	cur := this.head
	for deleteNode.next != nil {
		cur = cur.next
		deleteNode = deleteNode.next
	}
	cur.next = cur.next.next
}

func (this *LinkList) FindMiddleNode() *LinkNode {
	if nil == this.head || nil == this.head.next {
		return nil
	}

	if nil == this.head.next.next {
		return this.head.next
	}

	slow, faster := this.head, this.head
	for nil != faster && nil != faster.next {
		slow = slow.next
		faster = faster.next.next
	}

	return slow
}

func parseInt(i interface{}) (int, error) {
	s, ok := i.(string)
	if !ok {
		return 0, errors.New("not string")
	}
	return strconv.Atoi(s)
}

func MergeSortedList(l1, l2 *LinkList) *LinkList {
	if nil == l1 || nil == l1.head || nil == l1.head.next {
		return l2
	}

	if nil == l2 || nil == l2.head || nil == l2.head.next {
		return l1
	}

	l := NewLinkList()
	cur := l.head
	cur1 := l1.head.next
	cur2 := l2.head.next
	for nil != cur1 && nil != cur2 {
		a, _ := parseInt(cur1.val)
		b, _ := parseInt(cur2.val)
		if a < b {
			cur.next = cur1
			cur1 = cur1.next
		} else {
			cur.next = cur2
			cur2 = cur2.next
		}
		cur = cur.next
	}

	for nil != cur1 {
		cur.next = cur1
		cur1 = cur1.next
		cur = cur.next
	}

	for nil != cur2 {
		cur.next = cur2
		cur2 = cur2.next
		cur = cur.next
	}

	return l
}

func ReverseLinklist(list *LinkList) {
	if nil == list || list.length <= 1 {
		return
	}

	pre := list.head.next
	pre.next = nil
	cur := list.head.next.next
	for nil != cur {
		tmp := cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}
	list.head.next = cur
	return
}
