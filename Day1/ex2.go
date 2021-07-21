package main

import "fmt"

type tree struct{ // struct for tree
	left *tree
	right *tree
	value string
}

func newnode(value string) *tree { // initialize new node
	node := tree{nil, nil,value}
	return &node
}

func (node *tree) inorder(){ // inorder traversal
	if node != nil {
		node.left.inorder()
		fmt.Println(node.value)
		node.right.inorder()
	}
}

func (node *tree) preorder() { // preorder traversal
	if node != nil {
		fmt.Println(node.value)
		node.left.preorder()
		node.right.preorder()
	}
}

func (node *tree) postorder() { // postorder traversal
	if node != nil {
		node.left.postorder()
		node.right.postorder()
		fmt.Println(node.value)
	}
}

func main(){
	root := newnode("+")
	root.left = newnode("a")
	root.right = newnode("-")
	root.right.left = newnode("b")
	root.right.right = newnode("c")

	root.inorder()
	fmt.Println()
	root.preorder()
	fmt.Println()
	root.postorder()
}