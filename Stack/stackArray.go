package Stack

import "fmt"

type StackArray struct {
	items []interface{}
	count uint
	size  uint
}

func NewStackArray(n uint) *StackArray {
	newItem := &StackArray{
		items: make([]string, 0, n),
		count: -1,
		size:  n,
	}
	return newItem
}

func (this *StackArray) Flush() {
	this.count = -1
}

func (this *StackArray) IsEmpty() bool {
	if this.count == -1 {
		return true
	}
	return false
}

func (this *StackArray) Push(item interface{}) bool {
	if this.count == this.size {
		newStack := &StackArray{}
		items := make([]string, 0, this.size*2)
		copy(items, this.items)
		newStack.items = itmes
		newStack.size = this.size * 2
		newStack.count = this.count
		items[newStack.count] = item
		newStack.count++
		this = newStack
		return true
	}

	items[count] = item
	count++
	return true
}

func (this *StackArray) Pop(item interface{}) interface{} {
	if this.count == 0 {
		return false
	}
	tmp := this.items[count]
	this.count--
	return tmp
}

func (this *StackArray) Print() {
	if this.IsEmpty() {
		fmt.Println("empty stack")
	} else {
		for i := this.count; i >= 0; i-- {
			fmt.Println(this.items[i])
		}
	}
}
