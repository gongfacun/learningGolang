package basic

import (
	"fmt"
	"math"
	"math/cmplx"
)

func main() {

	fmt.Println("hello world")

	variablezeroone()
	variableinit()
	variableTypeDocument()
	eular()
	triangle()
	consts()
}

func variablezeroone() {
	var a int
	var b string

	fmt.Printf("%d is a %q is b", a, b)

}

func variableinit() {

	var aa int = 12
	var bb string = "my first"
	fmt.Println(aa, bb)
}

func variableTypeDocument() {
	var a, b, c int = 1, 2, 3
	d, e, f := 1, "two", true

	var (
		g = 1
		h = "hello"
		m = false
	)

	fmt.Println(a, b, c, d, e, f, g, h, m)
}

func eular() {

	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
}
func triangle() {

	var a, b int = 3, 4

	fmt.Println(colcTriangle(a, b))
}
func colcTriangle(a int, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}
func consts() {

	const (
		a  = 0
		bb = 1
		c  = 2
		d  = 3
		m  = 4
	)
	const (
		e = iota
		f
		g
		h
		l
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(a, bb, c, d, m)
	fmt.Println(e, f, g, h, l)
	fmt.Println(b, kb, mb, gb, tb, pb)

}
