package stack

// Stack defines behavior of a stack data structure
type Stack interface {
	Pop() interface{}
	Push(interface{})
	Peek() interface{}
	IsEmpty() bool
}

type intStack []int

func (s *intStack) Pop() int {
	n := (*s)[len(*s)-1]
	(*s) = (*s)[0 : len(*s)-1]
	return n
}

func (s *intStack) Push(n int) {
	*s = append(*s, n)
}

func (s *intStack) Peek() int {
	return (*s)[len(*s)-1]
}

func (s *intStack) IsEmpty() bool {
	return len(*s) == 0
}
