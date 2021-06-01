package queuestack

import (
	"fmt"
	"testing"

	"github.com/crosstheline/stack"
)

func TestNewQueueWithStack(t *testing.T) {

	stq1 := NewQueueWithStack()
	fmt.Println(stq1.Peek())
	stq1.Enqueue(1)
	fmt.Println(stq1.Peek())
	stq1.Enqueue(4)
	fmt.Println(stq1.Peek())
	stq1.Enqueue(6)
	fmt.Println(stq1.Peek())
	fmt.Println(stq1.Dequeue())
	fmt.Println(stq1.Peek())
	fmt.Println(stq1.Dequeue())
	fmt.Println(stq1.Peek())
	fmt.Println(stq1.Dequeue())
	fmt.Println(stq1.Peek())

	//tests := []struct {
	//	name string
	//	want queue.Queue
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if got := NewQueueWithStack(); !reflect.DeepEqual(got, tt.want) {
	//			t.Errorf("NewQueueWithStack() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
}

func TestQueueWithStack_Dequeue(t *testing.T) {
	type fields struct {
		stack1 stack.Stack
		stack2 stack.Stack
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qws := &QueueWithStack{
				stack1: tt.fields.stack1,
				stack2: tt.fields.stack2,
			}
			if got := qws.Dequeue(); got != tt.want {
				t.Errorf("Dequeue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueueWithStack_Enqueue(t *testing.T) {
	type fields struct {
		stack1 stack.Stack
		stack2 stack.Stack
	}
	type args struct {
		val int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//qws := &QueueWithStack{
			//	stack1: tt.fields.stack1,
			//	stack2: tt.fields.stack2,
			//}
		})
	}
}

func TestQueueWithStack_Peek(t *testing.T) {
	type fields struct {
		stack1 stack.Stack
		stack2 stack.Stack
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qws := &QueueWithStack{
				stack1: tt.fields.stack1,
				stack2: tt.fields.stack2,
			}
			if got := qws.Peek(); got != tt.want {
				t.Errorf("Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueueWithStack_moveStackValue(t *testing.T) {
	type fields struct {
		stack1 stack.Stack
		stack2 stack.Stack
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//qws := &QueueWithStack{
			//	stack1: tt.fields.stack1,
			//	stack2: tt.fields.stack2,
			//}
		})
	}
}
