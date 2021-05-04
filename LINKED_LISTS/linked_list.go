package main

import "fmt"

type node struct{
	data int
	next *node
}

type LinkedList struct{
	head *node
	length int
}

func (l *LinkedList) Prepend(n *node)*LinkedList{
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
	return l
}

func main(){
	mylist := LinkedList{}
	node1 := &node{data: 48}
	node2 := &node{data: 64}
	node3 := &node{data: 14}

	mylist.
		Prepend(node1).
		Prepend(node2).
		Prepend(node3)

	fmt.Println(mylist)
}
