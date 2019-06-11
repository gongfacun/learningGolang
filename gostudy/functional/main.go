package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func adder() func(int) int {

	sum := 0
	return func(i int) int {
		sum = sum + i

		return sum
	}
}

func fibnacci() inGen {

	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type inGen func() int

func (g inGen) Read(p []byte) (n int, err error) {

	next := g()
	if next >= 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {

		fmt.Println(scanner.Text())
	}
}

func main() {

	a := adder()
	for i := 1; i < 10; i++ {

		fmt.Printf("0,1,.... %d = %d\n", i, a(i))
	}

	b := fibnacci()
	for i := 0; i < 10; i++ {
		fmt.Println(b())
	}

	printFileContents(fibnacci())

}
