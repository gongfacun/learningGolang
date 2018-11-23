package main

import "fmt"

func main() {
	naturals := make(chan int)
	squarels := make(chan int)
	go counter(naturals)
	go squarer(squarels, naturals)
	printer(squarels)

}

func counter(out chan<- int) {

	for x := 0; x < 100; x++ {

		out <- x

	}
	close(out)

}

func squarer(out chan<- int, in <-chan int) {

	for x := range in {

		out <- x * x
	}
	close(out)
}

func printer(in <-chan int) {

	for x := range in {

		fmt.Println(x)

	}

}
