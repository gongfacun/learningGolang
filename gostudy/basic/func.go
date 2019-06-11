package basic

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) int {

	switch {

	case op == "+":
		return a + b
	case op == "-":
		return a - b
	case op == "*":
		return a * b
	case op == "/":
		return a / b
	default:
		panic("unsurpose operation:" + op)

	}
}
func div(a, b int) (q, r int) {

	return a / b, a % b
}
func pow(a int, b int) int {

	return int(math.Pow(float64(a), float64(b)))
}
func apply(op func(int, int) int, a, b int) int {
	pointer := reflect.ValueOf(op).Pointer()
	opname := runtime.FuncForPC(pointer).Name()
	fmt.Println("funcname:", opname)

	return op(a, b)
}
func swap(a, b int) (int, int) {

	return b, a
}

//func main() {
//fmt.Println(eval(5,4,"*"))
//fmt.Println(div(13,3))
//q,r:=div(13,3)
//fmt.Println(q,r)
//fmt.Println(apply(pow,3,5))
//fmt.Println(apply(func(i int, i2 int) int {
//	return int(math.Pow(float64(i),float64(i2)))
//},3,4))
//a,b :=4,7
//a,b=swap(a,b)
//fmt.Println(a,b)
//}
