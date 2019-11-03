package StackArray

type StackArray struct {
	items []string
	count uint
	size  uint
}

func NewStackArray(n uint) *StackArray {
	newItem := &StackArray{
		items: make([]string, n),
		count: 0,
		size:  n,
	}
	return newItem
}

func (this *StackArray) Push(item string) bool {
	if this.count == this.size {
		newStack := &StackArray{}
		items := make([]string, this.size*2)
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

func (this *StackArray) Pop(item string) string {
	if this.count == 0 {
		return false
	}
	tmp := this.items[count]
	this.count--
	return tmp
}
