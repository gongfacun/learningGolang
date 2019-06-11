package basic

import (
	"strconv"
)

func convertoB(n int) string {
	result := ""
	for ; n > 0; n /= 2 {

		lst := n % 2
		result = strconv.Itoa(lst) + result

	}
	return result
}

//func main() {
//
//	fmt.Println("应用开始")
//	fmt.Println(convertoB(3))
//	fmt.Println(convertoB(50))
//	fmt.Println(convertoB(128))
//}
