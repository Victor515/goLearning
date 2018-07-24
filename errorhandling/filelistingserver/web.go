package main

import (
	"net/http"
	"learngo/errorhandling/filelistingserver/filelisting"
	"os"
	"log"
)

type appHanlder func(writer http.ResponseWriter, request *http.Request) error

// define an interface in which error message can be shown to users
type userError interface {
	error
	Message() string
}

// wrap up our appHandler(functional programming)
func errWrapper(handler appHanlder) func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		// handle panic within this function
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

			}
		}()

		err := handler(w, r)
		
		if err != nil{
			log.Printf("Error occured: %s", err)

			// determine if the err is userError
			if err, ok := err.(userError); ok{
				http.Error(w, err.Message(), http.StatusBadRequest)
				return
			}

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(w, http.StatusText(code), code)
		}
	}
}


// hosting a web server listening request with path '/list/'
// error example:
// 1. "/abc"(path must start with list)
// 2. "/list/fib2.txt"(forbidden)

// implementing features:
// 1. Use functional programming to provide an uniform interface of error handling. errWrapper takes a hanlding function and returns
// the function type desired by http.HandleFunc
// 2. Use recover to handle unexpected error, so that panic is avoided
// 3. Create an userError interface to selectively choose to display error message
func main() {
	http.HandleFunc("/",
		errWrapper(filelisting.FileListHandler))

	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		panic(err)
	}
}
