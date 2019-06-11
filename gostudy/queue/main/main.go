package main

import (
	"fmt"
	"gostudy/queue"
)

func main() {
	q := queue.Queue{3, 2, 1}
	q.Push(4)
	q.Push(5)
	q.Push(6)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
}
