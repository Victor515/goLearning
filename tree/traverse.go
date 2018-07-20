package tree

import "fmt"

func (node *Node) InOrder(){
	if node == nil{
		return
	}

	node.Left.InOrder()
	node.Print()
	node.Right.InOrder()
}

// functional programming style
func (node *Node) InOrderFunc(f func(*Node)){
	if(node == nil){
		return
	}

	node.Left.InOrderFunc(f)
	f(node)
	node.Right.InOrderFunc(f)
}

func (node *Node) InOrderPrint(){
	node.InOrderFunc(func(n *Node){
		n.Print()
	})
	fmt.Println()
}