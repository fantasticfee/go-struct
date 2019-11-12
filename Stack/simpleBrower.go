package Stack

type Brower struct {
	forwardStack  StackArray
	backwardStack StackArray
}

func NewBrower() *Brower {
	return &Brower{
		forwardStack:  NewStackArray(uint(32)),
		backwardStack: NewStackArray(uint(32)),
	}
}

func (b *Brower) CanForward() bool {
	if b.forwardStack.IsEmpty() {
		return false
	}

	return true
}

func (b *Brower) CanBackward() bool {
	if b.backwardStack.IsEmpty() {
		return false
	}
	return true
}

func (b *Brower) PushBack(v interface{}) {
	b.backwardStack.Push(v)
}

//func (b *Brower) PopBack() {
//	b.backwardStack.Pop()
//}

func (b *Brower) Back() {
	if !b.CanBackward() {
		return
	}
	top := b.backwardStack.Pop()
	b.forwardStack.Push(top)
	return
}

func (b *Brower) Forward() {
	if !b.CanForward() {
		return
	}
	top := b.forwardStack.Pop()
	b.backwardStack.Push(top)
	return
}

func (b *Brower) PushForward(v interface{}) {
	b.forwardStack.Push(v)
}

//func (b *Brower) PopForward(v interface{}) {
//	b.forwardStack.Pop(v)
//}
