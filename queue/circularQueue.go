package queue

type CircularQueue struct {
	items []interface{}
	head int
	tail int
	size int
}

func NewCircularQueue(size int) *CircularQueue {
	return &CircularQueue{make([]interface{},size),0,0,size}
}

func (this *CircularQueue) EnQueue(v interface{}) bool {
	if (this.tail+1)%this.size == this.head {
		return false
	}
	this.items[tail] = v
	this.tail= (this.tail+1)%this.size
	return true
}

func (this *CircularQueue) DeQueue() interface{} {
	if this.tail == this.head {
		return nil
	}

	v := this.items[head]
	this.head = (this.head+1)%this.size
	return true
}
