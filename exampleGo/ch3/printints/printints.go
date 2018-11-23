package main

import "fmt"
import "bytes"

func main() {
	fmt.Println("input: 1234567")

	fmt.Println(intsToString([]int{1, 2, 3, 4, 5, 6, 7}))
}

func intsToString(values []int) string {

	var buf bytes.Buffer
	buf.WriteByte('[')

	for i, v := range values {

		if i > 0 {

			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)

	}
	buf.WriteByte(']')

	return buf.String()

}
