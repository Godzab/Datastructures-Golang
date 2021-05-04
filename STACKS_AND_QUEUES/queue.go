package main

import "fmt"

type Queue struct{
	items []int
}

func (q *Queue) Enqueue(i int){
	q.items = append(q.items, i)
}

func (q *Queue) Dequeue()int{
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main(){
	myQueue := Queue{}
	myQueue.Enqueue(5)
	myQueue.Enqueue(25)
	myQueue.Enqueue(33)
	myQueue.Enqueue(18)
	fmt.Println(myQueue)
	i1 := myQueue.Dequeue()
	fmt.Println(i1, myQueue)
	i2 := myQueue.Dequeue()
	fmt.Println(i2, myQueue)
}