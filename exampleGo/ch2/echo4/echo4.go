package main

import "fmt"
import "flag"
import "strings"

var n = flag.Bool("n", false, "omit trailing newline")
var s = flag.String("s", "", "separater")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *s))

	if !*n {

		fmt.Println()
	}

}
