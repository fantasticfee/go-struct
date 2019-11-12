package queue

type LinkNode struct {
	v    interface{}
	next *LinkNode
}

type QueueLinklist struct {
	head   *LinkNode
	tail   *LinkNode
	length int
}

func NewQueueLinklist() *QueueLinklist {
	return &QueueLinklist{nil, nil, 0}
}

func (this *QueueLinklist) EnQueue(v interface{}) {
	newNode := &LinkNode{v, nil}
	if this.tail == nil {
		this.head = newNode
		this.tail = newNode
	} else {
		this.tail.next = newNode
		this.tail = this.tail.next
	}
	this.length++
}

func (this *QueueLinklist) DeQueue() interface{} {
	if this.head == nil {
		return nil
	}
	v := this.head.val
	this.head = this.head.next
	this.length--
	return v
}
