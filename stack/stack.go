package stack

import "fmt"

type Stack struct {
	list []interface{}
	len  int
	size int
}

func New(maxSize int) Stack {
	return Stack{
		list: make([]interface{}, maxSize),
		len:  0,
		size: maxSize,
	}
}

func (s *Stack) IsFull() bool {
	return s.len == s.size
}

func (s *Stack) IsEmpty() bool {
	return s.len == 0
}

func (s *Stack) Push(element interface{}) bool {
	// Check if the stack has space or not
	if s.IsFull() {
		return false
	}
	// insert the new element to the top of stack
	s.list[s.len] = element
	s.len++
	return true
}

func (s *Stack) Pop() (interface{}, bool) {
	// check the stack not to be empty
	if s.IsEmpty() {
		return nil, false
	}

	// get the top element of stack and return it
	element := s.list[s.len-1]
	s.len--
	return element, true
}

func (s *Stack) Top() (interface{}, bool) {
	// check the stack not to be empty
	if s.IsEmpty() {
		return nil, false
	}

	// return the top element of stack
	return s.list[s.len-1], true
}

func (s *Stack) Clear() {
	s.len = 0
}

func (s *Stack) Print() {
	for i := 0; i < s.len; i++ {
		fmt.Print(s.list[i])
		if i != s.len-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
