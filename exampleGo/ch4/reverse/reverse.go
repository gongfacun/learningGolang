package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	reverse(&a)
	fmt.Println(a)
	e := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	rotate(e)
	fmt.Println(e)

	str := []string{"sd", "ds", "ds", "dsd", "dsd", "dsd", "ddd", "ccc", "dsd", "dsds"}
	str = deduplication(str)
	fmt.Println(str)

}

func reverse(s *[10]int) {

	for i, j := 0, len((*s))-1; i <= j; i, j = i+1, j-1 {

		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]

	}

}

func rotate(b []int) {

	lenss := len(b) - 1
	o := 1
	for i := lenss; i >= 0; i = i - 1 {
		b = append(b, b[i])
		o++

	}

	copy(b[0:lenss+1], b[lenss+1:])

}

func deduplication(dupl []string) []string {

	newstr := []string{}
	for i, s := range dupl {

		if i >= len(dupl)-1 {

		} else {
			if s == dupl[i+1] {

			} else {

				newstr = append(newstr, s)
			}

		}
	}

	fmt.Println(newstr)
	fmt.Println(".......")
	return newstr
}
