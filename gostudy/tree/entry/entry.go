package main

import (
	"fmt"
	"gostudy/tree"
)

type mytreenode struct {
	node *tree.Node
}

func (node *mytreenode) postOrder() {
	if node == nil || node.node == nil {
		return
	}
	leftNode := mytreenode{node.node.Left}
	rightNode := mytreenode{node.node.Right}
	leftNode.postOrder()
	rightNode.postOrder()
	node.node.Print()

}
func main() {

	var root tree.Node
	root = tree.Node{nil, nil, 0}
	fmt.Println(root)
	root.Left = &tree.Node{Value: 4}
	root.Right = &tree.Node{nil, nil, 1}
	root.Right.Left = new(tree.Node)
	root.Right.Left.Value = 5
	root.Left.Right = tree.CreateNode(2)
	root.Left.Value = 10
	root.Left.Print()
	proot := &root
	proot.Left.SetValue(100)
	proot.Right.SetValue(50)
	nodes := []tree.Node{
		{nil, nil, 1},
		{Value: 6},
		root,
	}
	fmt.Println(nodes)

	root.TreePrint()
	fmt.Println()
	tmpnode := mytreenode{&root}
	tmpnode.postOrder()
	fmt.Println()
	root.TreePointerCount()
}
