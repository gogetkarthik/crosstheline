package queue

type queueStore []int

type Queue interface {
	Dequeue() int
	Enqueue(val int)
	Peek() int
}

func NewQueue() Queue {
	return &queueStore{}
}

func (q *queueStore) Dequeue() int {
	var rVal int
	qLen := len(*q)
	if qLen > 0 {
		rVal = (*q)[0]
		*q = (*q)[1:]
	}

	return rVal
}

func (q *queueStore) Enqueue(val int) {
	*q = append(*q, val)
}

func (q *queueStore) Peek() int {
	if len(*q) > 0 {
		return (*q)[0]
	}

	return 0
}
