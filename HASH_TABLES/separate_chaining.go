package main

import "fmt"

const arraySize = 7

type HashTable struct{
	array [arraySize]*Bucket
}

type Bucket struct{
	head *BucketNode
	length int
}

type BucketNode struct{
	key string
	next *BucketNode
}

func (h *HashTable) Insert(val string){
	index := hash(val)
	h.array[index].insert(val)
}


func (h *HashTable) Search(val string)bool{
	index := hash(val)
	bucket := h.array[index]
	return bucket.Search(val)
}

func (b *Bucket) insert(key string){
	newNode := &BucketNode{key: key}
	second := b.head
	b.head = newNode
	b.head.next = second
	b.length++
}

func (b *Bucket) Search(val string)bool{
	current := b.head
	for current	!= nil{
		if current.key == val{
			return true
		}
		current = current.next
	}
	return false
}

func hash(key string)int{
	sum := 0
	for _, v := range key{
		sum += int(v)
	}
	return sum % arraySize
}

func Init()*HashTable{
	result := &HashTable{}
	for i := range result.array{
		result.array[i] = &Bucket{}
	}
	return result
}

func main(){
	 ht := Init()
	 ht.Insert("Godzab")
	 ht.Insert("Godzab")
	 fmt.Println(ht)
	 fmt.Println(ht.Search("Godzab"))
	 fmt.Println(ht.Search("Kedaman"))
}
