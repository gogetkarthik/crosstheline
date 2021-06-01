package queuestack

import (
	"github.com/crosstheline/queue"
	"github.com/crosstheline/stack"
)

type QueueWithStack struct {
	stack1 stack.Stack
	stack2 stack.Stack
}

func NewQueueWithStack() queue.Queue {
	return &QueueWithStack{
		stack1: stack.NewStack(),
		stack2: stack.NewStack(),
	}
}

func (qws *QueueWithStack) Dequeue() int {
	qws.moveStackValue()
	return qws.stack2.Pop()
}

func (qws *QueueWithStack) Enqueue(val int) {
	qws.stack1.Push(val)

}

func (qws *QueueWithStack) Peek() int {
	qws.moveStackValue()
	return qws.stack2.Peek()
}

func (qws *QueueWithStack) moveStackValue() {
	if qws.stack2.Size() < 1 {
		if qws.stack1.Size() > 0 {
			for qws.stack1.Size() > 0 {
				qws.stack2.Push(qws.stack1.Pop())
			}
		}
	}
}
