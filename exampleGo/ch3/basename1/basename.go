package main

import "fmt"
import "strings"

func main() {

	fmt.Println("input: a/b/c/d.sob")
	fmt.Println(basename("a/b/c/d.sob"))

}

func basename1(s string) string {

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '/' {

			s = s[i+1:]
			break
		}

	}

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {

			s = s[:i]
			break
		}

	}

	return s

}

func basename(s string) string {
	salsh := strings.LastIndex(s, "/")
	s = s[salsh+1:]

	if dot := strings.LastIndex(s, "."); dot >= 0 {

		s = s[:dot]
	}

	return s
}
