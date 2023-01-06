package main

import "fmt"

type Stack struct {
	List []string
}

func main() {
	var stack Stack
	fmt.Println(stack.IsEmpty())
	fmt.Println(stack.Pop())
	stack.Append("James")
	stack.Append("Confidence")
	stack.Append("Victory")
	stack.Append("Mesoma")
	stack.Append("Marvelous")
	fmt.Println(stack.Pop())
	stack.Print()
}

func (s *Stack) Append(value string) {
	s.List = append(s.List, value)
}

func (s *Stack) Pop() string {
	if s.IsEmpty() {
		return "nil"
	}
	poppedValue := s.List[len(s.List)-1]
	s.List = s.List[:len(s.List)-1]
	return poppedValue
}

func (s *Stack) Print() {
	fmt.Println(s.List)
}

func (s *Stack) IsEmpty() bool {
	return len(s.List) == 0
}
