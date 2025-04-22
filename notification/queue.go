package notification

type Queue struct {
	ch chan Notification
}

func NewQueue(bufferSize int) *Queue {
	return &Queue{
		ch: make(chan Notification, bufferSize),
	}
}

func (q *Queue) Enqueue(n Notification) {
	q.ch <- n
}

func (q *Queue) Dequeue() <-chan Notification {
	return q.ch
}
