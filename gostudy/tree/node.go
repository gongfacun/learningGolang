package tree

import "fmt"

type Node struct {
	Left, Right *Node
	Value       int
}

func (node Node) Print() {
	fmt.Println(node.Value)
}
func CreateNode(value int) *Node {

	return &Node{Value: value}
}
func (node *Node) SetValue(value int) {

	node.Value = value
}
func (node *Node) TreePrint() {
	node.TreePriontFunc(func(n *Node) {
		n.Print()
	})

	fmt.Println("end")
}

func (node *Node) TreePointerCount() {
	countnum := 0
	node.TreePriontFunc(func(n *Node) {
		countnum++
	})

	fmt.Println("pointer__count:", countnum)

}
func (node *Node) TreePriontFunc(f func(*Node)) {

	if node == nil {
		return
	}
	node.Left.TreePriontFunc(f)
	f(node)
	node.Right.TreePriontFunc(f)
}
