package singlinklist

import "fmt"

type LinkNode struct {
	next  *LinkNode
	value interface{}
}

type LinkList struct {
	head   *LinkNode
	length uint
}

func NewLinkNode(v interface{}) *ListNode {
	return &ListNode{nil, v}
}

func (this *ListNode) GetNext() *ListNode {
	return this.next
}

func (this *ListNode) GetValue() *ListNode {
	return this.value
}

func NewLinkList() *LinkList {
	return &LinkList{NewLinkNode(0), 0}
}

//在某个节点后面插入节点
func (this *LinkList) InsertAfter(p *LinkNode, v interface{}) bool {
	if nil == p {
		return false
	}

	newNode := NewLinkNode(v)
	var tmp *LinkNode
	tmp = p.next
	p.next = newNode
	newNode.next = tmp
	this.length++
	return true
}

//在某个节点前面插入节点
func (this *LinkList) InsertBefore(p *LinkNode, v interface{}) bool {
	if p == nil {
		return false
	}

	cur := this.head
	after := this.head.next
	for after != nil {
		if after == p {
			break
		}
		cur = after
		after = after.next
	}
	if after == nil {
		return false
	}

	newNode := NewLinkNode(v)
	newNode.next = cur.next
	cur.next = newNode.next
	this.length++

	return true
}

//在链表头部插入节点
func (this *LinkedList) InsertToHead(v interface{}) bool {
	return this.InsertAfter(this.head, v)
}

//在链表尾部插入节点
func (this *LinkedList) InsertToTail(v interface{}) bool {
	cur := this.head
	for nil != cur.next {
		cur = cur.next
	}
	return this.InsertAfter(cur, v)
}

//通过索引查找节点
func (this *LinkedList) FindByIndex(index uint) *ListNode {
	cur := this.head.next

	if index >= this.length {
		return nil
	}

	for i = 0; i < index; i++ {
		cur = cur.next
	}
	return cur
}

//删除传入的节点
func (this *LinkedList) DeleteNode(p *ListNode) bool {
	if nil == p {
		return false
	}

	cur := this.head.next
	pre := this.head
	for cur != nil {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if cur == nil {
		false
	}
	pre.next = cur.next
	cur.next = nil
	this.length--
	return true
}

//delete the nth node of linklist from tail
func (this *LinkedList) DeleteTailNthNode(l *LinkList, num uint) bool {
	if l == nil {
		return false
	}

	iLen := l.length
	if iLen == 0 || num > iLen {
		return false
	}

	ReverseLinkList(l)
	if num == 1 {
		tmp := l.head.next
		l.head.next = tmp.next
		tmp.next = nil
		return true
	}

	var pre *LinkNode
	cur := l.head
	for i := 1; i <= num; i++ {
		pre = cur
		cur = cur.next
	}
	pre.next = cur.next
	cur.next = nil

	return true
}

func (this *LinkedList) DeleteTailNthNode2(l *LinkList, num uint) bool {
	if l == nil {
		return false
	}

	iLen := l.length
	if iLen == 0 || num > iLen {
		return false
	}
	preNum = iLen - num
	var pre *LinkNode
	cur := l.head.next
	for i := 1; i <= preNum && cur != nil; i++ {
		pre = cur
		cur = cur.next
	}
	pre.next = cur.next
	cur.next = nil
	return true
}

//
func (this *LinkedList) CombineOrderedLinklist(a *LinkList, b *LinkList) *LinkList {
	if a == nil || a.length == 0 {
		return b
	}
	if b == nil || b.length == 0 {
		return a
	}
	c := NewLinkList()
	cur1 := a.head.next
	cur2 := b.head.next
	res := c.head
	for cur1 != nil || cur2 != nil {
		if cur2 == nil || cur1.GetValue() >= cur2.GetValue() {
			res.next = cur1
			res = res.next
			cur1 = cur1.next
		} else if cur1 == nil || cur2.GetValue() > cur1.GetValue() {
			res.next = cur2
			res = res.next
			cur2 = cur2.next
		}
	}
	return c
}

//
func (this *LinkedList) FindMediumLinklist(l *LinkList) *LinkNode {
	if l == nil {
		return nil
	}
	iLen := l.length
	if iLen == 0 {
		return nil
	}
	if iLen == 1 {
		return l.head.next
	}
	var midNum uint

	if iLen%2 != 0 {
		midNum = iLen/2 + 1
	} else {
		midNum = iLen / 2
	}
	cur := l.head
	for i := 1; i <= midNum; i++ {
		cur = cur.next
	}
	return cur
}

//打印链表
func (this *LinkedList) Print() {
	cur := this.head.next
	format := ""
	for cur != nil {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if cur != nil {
			format += "->"
		}
	}
	fmt.Println(format)
}
