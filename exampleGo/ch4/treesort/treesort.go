package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	mint := []int{3, 7, 9, 6, 5, 4, 2, 8, 1}
	fmt.Println(mint)
	Sort(mint)
}

type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	sortint := appendValues(values[:0], root)

	fmt.Println(sortint)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		fmt.Println(t.value)
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		fmt.Println("save_value:", value)
		return t
	}
	fmt.Println("t_value:", t.value)
	fmt.Println("value:", value)
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
