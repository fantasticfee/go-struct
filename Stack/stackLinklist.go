package Stack

import "fmt"

type StackNode struct {
	next *StackNode
	val  interface{}
}

type StackLinklist struct {
	top *StackNode
}

func NewStackLinklist() *StackLinklist {
	return &StackLinklist{top: nil}
}

func (this *StackLinklist) IsEmpty() bool {
	if this.top == nil {
		return true
	}
	return false
}

func (this *StackLinklist) Push(v interface{}) {
	this.top = &StackNode{next: this.top, val: v}
}

func (this *StackLinklist) Pop() interface{} {
	if this.IsEmpty() {
		return nil
	}

	node := this.top.val
	this.top = this.top.next
	return node
}

func (this *StackLinklist) Top() interface{} {
	if this.IsEmpty() {
		return nil
	}
	return this.top.val
}

func (this *StackLinklist) Flush() {
	this.top = nil
}

func (this *StackLinklist) Print() {
	if this.IsEmpty() {
		fmt.Println("empty stack")
	} else {
		cur := this.top
		for cur != nil {
			fmt.Println(cur.val)
			cur = cur.next
		}
	}
}
