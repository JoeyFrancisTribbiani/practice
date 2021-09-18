package queue

type MyCircularQueue struct {
	head int
	rear int
	len  int
	list []int
}

func ConstructorQueue(k int) MyCircularQueue {
	q := new(MyCircularQueue)
	list := []int{}
	for i := 0; i < k; i++ {
		list = append(list, 0)
	}
	q.list = list
	q.len = k
	q.head = -1
	q.rear = -1
	return *q
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.IsEmpty() {
		this.head = 0
		this.rear = 0
		this.list[0] = value
		return true
	}
	if this.rear == this.len-1 {
		this.rear = 0
	} else {
		this.rear = this.rear%this.len + 1
	}
	this.list[this.rear] = value
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	if this.head == this.rear {
		this.head = -1
		this.rear = -1
		return true
	}
	if this.head == this.len-1 {
		this.head = 0
	} else {
		this.head = this.head%this.len + 1
	}
	return true
}
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.list[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.list[this.rear]
}

func (this *MyCircularQueue) IsEmpty() bool {
	if this.head == -1 && this.rear == -1 {
		return true
	}
	return false
}

func (this *MyCircularQueue) IsFull() bool {
	if this.head > this.rear && this.head-this.rear == 1 {
		return true
	}
	if this.head == 0 && this.rear == this.len-1 {
		return true
	}
	return false
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
