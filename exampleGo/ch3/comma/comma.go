package main

import "fmt"
import "bytes"

func main() {
	fmt.Println("input 1234567")
	fmt.Println(comma1("1234567"))
}

func comma(s string) string {

	n := len(s)
	if n <= 3 {

		return s
	}

	s = comma(s[:n-3]) + "," + s[n-3:]
	return s

}

func comma1(s string) string {

	n := len(s)

	if n <= 3 {

		return s

	}
	fmt.Println(n)

	var buf bytes.Buffer
	for i := n % 3; i <= n; i += 3 {

		fmt.Println(i)
		if i > 3 {
			buf.WriteString(",")
			buf.WriteString(s[i-3 : i])
		} else {
			buf.WriteString(s[:i])
		}

	}

	return buf.String()

}
