package filelist

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type UserError interface {
	error
	Message() string
}

type Usererror string

func (err Usererror) Message() string {

	return string(err)
}
func (err Usererror) Error() string {
	return err.Message()

}

const prefix = "/list/"

func FileHandle(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {
		return Usererror(" 必须以/list开始")
	}
	path := request.URL.Path[len(prefix):]

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)

	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}
