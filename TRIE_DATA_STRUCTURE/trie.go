package main

import "fmt"

// Number of possible tries characters
const alphabetSize = 26

//Node
type Node struct {
	children [alphabetSize]*Node
	isEnd bool
}

//Trie
type Trie struct{
	root *Node
}

func NewTrie() *Trie {
	return &Trie{root: &Node{}}
}

//Insert method
func (t *Trie) Insert(w string){
	currentNode := t.root
	for _, char := range w{
		//Changes the character to the character index
		charIndex := char - 'a'
		if currentNode.children[charIndex] == nil{
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}
//Search method
func (t *Trie) Search(w string) bool{
	currentNode := t.root
	for _, char := range w{
		//Changes the character to the character index
		charIndex := char - 'a'
		if currentNode.children[charIndex] == nil{
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	return currentNode.isEnd
}

func main(){
	tt := NewTrie()
	tt.Insert("aragorn")
	tt.Insert("oreo")
	tt.Insert("argorn")
	tt.Insert("eragorn")
	tt.Insert("oregon")
	tt.Insert("oregano")
	fmt.Println(tt.root)
}
