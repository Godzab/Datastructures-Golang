package main

import "fmt"

type MaxHeap struct{
	array []int
}

func (h *MaxHeap) Insert(key int){
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array)-1)
}

func (h *MaxHeap) Extract() int{
	extracted := h.array[0]
	l := len(h.array) -1

	if l <= 0{
		fmt.Println("Cannot extract because array length is 0.")
		return -1
	}

	h.array[0] = h.array[l]
	h.array = h.array[:l]
	h.maxHeapifyDown(0)
	return extracted
}

func (h *MaxHeap) maxHeapifyUp(index int){
	for h.array[parent(index)] < h.array[index]{
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) maxHeapifyDown(index int){
	lastIdx := len(h.array) -1
	l, r := left(index), right(index)
	childToCompare := 0
	for l <= lastIdx{
		if l == lastIdx{
			childToCompare = l
		} else if h.array[l] > h.array[r]{
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.array[index] < h.array[childToCompare]{
			h.swap(index,childToCompare)
			index = childToCompare
			l,r = left(index), right(index)
		} else {
			return
		}
	}
}

// parent Get parent index
func parent(idx int) int{
	return (idx - 1) / 2
}

func left(idx int) int{
	return (idx * 2) + 1
}

func right(idx int) int{
	return (idx * 2) + 2
}

func (h *MaxHeap) swap(idx1, idx2 int){
	h.array[idx1], h.array[idx2] = h.array[idx2], h.array[idx1]
}

func main(){
	a := []int{12, 15,30,25,66,8,5,9}
	m := MaxHeap{}
	for _, v := range a{
		m.Insert(v)
	}

	m.Extract()
	m.Extract()
	m.Extract()

	fmt.Println(m)
}
