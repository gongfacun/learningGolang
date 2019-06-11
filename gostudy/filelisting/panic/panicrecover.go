package main

import (
	"errors"
	"fmt"
)

func tryfunc() {

	defer func() {
		err := recover()

		if herr, ok := err.(error); !ok {

			panic(herr)

		} else {

			fmt.Println("this is error")

		}

	}()
	panic(errors.New("this is error,can you catch it"))
}
func main() {
	tryfunc()
}
