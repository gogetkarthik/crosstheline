package stack

type Stack interface {
	Peek() int
	Pop() int
	Push(val int)
	Size() int
}

type store []int

func NewStack() Stack {
	return &store{}
}

func (s *store) Peek() int {
	l := len(*s)

	if l < 1 {
		return 0
	}

	return (*s)[l-1]
}

func (s *store) Pop() int {
	l := len(*s)

	if l < 1 {
		return 0
	}

	rVal := (*s)[l-1]

	*s = (*s)[:l-1]
	return rVal
}

func (s *store) Push(val int) {
	*s = append(*s, val)
}

func (s *store) Size() int {
	return len(*s)
}
