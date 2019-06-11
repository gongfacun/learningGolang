package main

import (
	"bufio"
	"fmt"
	"gostudy/functional/fib"
	"os"
)

func Writefile(name string) {
	file, err := os.OpenFile(name, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {

			fmt.Printf("%s  %s %s \n", pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close()

	buf := bufio.NewWriter(file)
	defer buf.Flush()
	f := fib.Fibnacci()
	for i := 0; i < 30; i++ {

		fmt.Fprintln(buf, f())
	}

}
func main() {
	Writefile("fib.txt")
}
