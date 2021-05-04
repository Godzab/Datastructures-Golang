package main

import "fmt"

type Stack struct{
	items []int
}

func (s *Stack) Push(i int){
	s.items = append(s.items, i)
}

func (s *Stack) Pop() int{
	popped := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return popped
}


func main(){
	myStack := Stack{}
	myStack.Push(3)
	myStack.Push(4)
	myStack.Push(15)
	fmt.Println(myStack)
	dt := myStack.Pop()
	fmt.Println(dt, myStack)
}
