package basic

import "fmt"

func printArray(arr [5]int) {
	arr[0] = 100 //此更改在函数内起作用
	for _, v := range arr {

		fmt.Println(v)
	}
}
func printArray1(arr *[5]int) {
	arr[0] = 101
	for _, v := range arr {
		fmt.Println(v)
	}

}

//func main() {
//	var  a [3]int
//	b :=[4]int{3,4,5,6}
//	c :=[5]int{9,8,7,6,5}
//	d :=[...]int{5,0,3,5,4,6,9,1}
//	fmt.Println(a)
//	for i:=0;i<len(b);i++{
//		fmt.Println(b[i])
//	}
//	for i :=range c{
//
//		fmt.Println(c[i])
//
//	}
//	for _,v:=range d{
//		fmt.Println(v)
//	}
//
//	printArray(c)
//	fmt.Println(c)
//	printArray1(&c)
//	fmt.Println(c)
//}
