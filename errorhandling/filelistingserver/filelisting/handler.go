package filelisting

import (
	"net/http"
	"os"
	"io/ioutil"
	"strings"
)

const prefix = "/list/"

type userError string

func (e userError) Error() string{
	return e.Message()
}

func (e userError) Message() string{
	return string(e)
}

func FileListHandler (writer http.ResponseWriter, request *http.Request) error{
	if strings.Index(request.URL.Path, prefix) != 0{
		return userError("path must start with " +  prefix)
	}
	path := request.URL.Path[len(prefix):] // get the relative path
	file, err := os.Open(path)
	if err != nil{
		return err
	}
	defer file.Close()

	// read bytes from the file
	all, err := ioutil.ReadAll(file)
	if err != nil{
		return err
	}

	// write contents to request
	writer.Write(all)
	return nil
}
