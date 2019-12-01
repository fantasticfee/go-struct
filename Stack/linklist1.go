package Stack

import "fmt"

type node struct {
	next *node
	val  interface{}
}

type LinkListStack struct {
	topNode *node
}

func NewLinkListStack() *LinkListStack {
	return &LinkListStack{nil}
}

func (this *LinkListStack) IsEmpty() bool {
	return this.topNode == nil
}

func (this *LinkListStack) Push(v interface{}) {
	this.topNode = &node{next: this.topNode, val: v}
}

func (this *LinkListStack) Pop() interface{} {
	if this.IsEmpty() {
		return nil
	}

	v := this.topNode.val
	this.topNode = this.topNode.next
	return v
}

func (this *LinkListStack) Top() interface{} {
	if this.IsEmpty() {
		return nil
	}
	return this.topNode.val
}

func (this *LinkListStack) Flush() {
	this.topNode = nil
}

func (this *LinkListStack) Print() {
	if this.IsEmpty() {
		fmt.Println("empty stack")
	} else {
		cur := this.topNode
		for cur != nil {
			fmt.Println(cur.val)
			cur = cur.next
		}
	}
}
