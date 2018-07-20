package main

import (
	"learngo/tree"
)

// extending tree methods with struct
type myTreeNode struct{
	node *tree.Node
}

func (myNode *myTreeNode) postOrder(){
	if myNode == nil || myNode.node == nil{
		return
	}
	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node

	// construct tree
	root = tree.Node{Val: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node) // always use dot(.) to reference(address and struct is the same)
	root.Right.Left.SetValue(10)
	root.Right.Right = tree.CreateNode(3)

	//// traverse
	//root.InOrder()
	//fmt.Println()
	//myRoot := myTreeNode{&root}
	//myRoot.postOrder()

	//root.print()
	//root.setValue(5)
	//root.print()
	//
	//// always use dot to reference
	//pRoot := &root
	//pRoot.print()
	//pRoot.setValue(10)
	//pRoot.print

	// new functional style
	root.InOrderPrint()
}

