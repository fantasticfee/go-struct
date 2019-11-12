package main

type QueueArray struct {
	queue    []interface{}
	head     int
	tail     int
	capacity int
}

func NewQueueArray(size int) *QueueArray {
	return &QueueArray{head: 0, tail: 0, capacity: size, queue: make([]interface{}, size)}
}

func (this *QueueArray) EnQueue(v interface{}) bool {
	if this.tail == this.capacity {
		if this.head == 0 {
			return false
		}

		for i := this.head; i <= this.tail; i++ {
			this.queue[i-this.head] = this.queue[i]
		}
		tail -= head
		head = 0
	}

	this.queue[this.tail] = v
	this.tail++
	return true
}

func (this *QueueArray) DeQueue() interface{} {
	if this.tail == this.head {
		return nil
	}
	v := this.queue[this.head]
	this.head++
	return v
}
