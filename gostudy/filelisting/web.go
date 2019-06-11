package main

import (
	"gostudy/filelisting/filelist"
	"net/http"
	"os"
)

type apphandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handle apphandler) func(writer http.ResponseWriter, request *http.Request) {

	return func(writer http.ResponseWriter, request *http.Request) {

		defer func() {
			if r := recover(); r != nil {
				http.Error(writer, "inner error", http.StatusInternalServerError)

			}

		}()

		err := handle(writer, request)

		if err != nil {

			if uerr, ok := err.(filelist.UserError); ok {

				http.Error(writer, uerr.Error(), http.StatusBadRequest)
				return

			}

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsTimeout(err):
				code = http.StatusRequestTimeout
			default:
				code = http.StatusServiceUnavailable

			}
			http.Error(writer, http.StatusText(code), code)

		}
	}
}

func main() {
	http.HandleFunc("/", errWrapper(filelist.FileHandle))
	err := http.ListenAndServe(":8888", nil)

	if err != nil {

		panic(err)
	}

}
