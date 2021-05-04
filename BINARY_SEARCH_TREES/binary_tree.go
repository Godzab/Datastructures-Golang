package main

import "fmt"

//Node
type Node struct{
	Key int
	Left, Right *Node
}

//Insert Method
func (n *Node) Insert(k int){
	if n.Key < k{
		//Move right
		if n.Right == nil{
			n.Right = &Node{Key: k}
		} else{
			n.Right.Insert(k)
		}
	} else if n.Key > k{
		//Move left
		if n.Left == nil{
			n.Left = &Node{Key: k}
		} else{
			n.Left.Insert(k)
		}
	} else {
		return
	}
}


//Search Method
func (n *Node) Search(k int)bool{
	if n == nil{
		return false
	}
	if n.Key < k{
		return n.Right.Search(k)
	} else if n.Key > k{
		return n.Left.Search(k)
	}
	return true
}

func main(){
	tree := &Node{Key: 100}
	tree.Insert(50)
	tree.Insert(115)
	tree.Insert(300)
	tree.Insert(56)
	fmt.Println(tree.Search(300))
}
