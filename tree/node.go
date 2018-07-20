package tree

import "fmt"

type Node struct {
	Val int
	Left, Right *Node
}

// write methods for treeNode
func (node Node) Print(){
	fmt.Print(node.Val, " ")
}

// go is pass by value, so func (node treeNode) setValue(val int) won't work
func (node *Node) SetValue(val int){
	node.Val = val
}

// no constructor, and we could build factory method
func CreateNode(val int) *Node {
	// note: we return the address of a local variable
	// question: where is this variable allocated? heap or stack?
	// answer: no need to know, depends on compilers and runtime
	return &Node{Val: val}
}


