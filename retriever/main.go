package main

import (
	"learngo/retriever/mock"
	real2 "learngo/retriever/real"
	"fmt"
)

// Features in Go: the user defines what the interface should look like
// Interfaces are implemented implicitly
type Retriever interface{
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func DownLoad(r Retriever) string{
	return r.Get("https://www.google.com")
}

func post(poster Poster){
	poster.Post("http://www.imooc.com", map[string]string{
		"name" : "victor",
		"course": "golang",
	})
}

func session(s RetrieverPoster) string{
	s.Post("http://www.imooc.com", map[string]string{
		"contents" : "A fake imooc site",
	})
	return s.Get("http://www.imooc.com")
}

func inspect(r Retriever){
	fmt.Printf("%T, %v\n", r, r)
	// type switch
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents", v.Contents)
	case real2.Retriever:
		fmt.Println("Timeout: ", v.TimeOut)
	}
}

func main() {
	var r RetrieverPoster
	contents := "this is immoc site"
	r = mock.Retriever{&contents}
	inspect(r)
	//r = real2.Retriever{TimeOut:time.Minute}
	//inspect(r)

	// type assertion
	//mockRetriever := r.(mock.Retriever) // this will cause panic!
	//fmt.Println(mockRetriever.Contents)

	// non-strict type assertion
	//if realRetriever, ok := r.(real2.Retriever); ok{
	//	fmt.Println(realRetriever.TimeOut)
	//}else{
	//	fmt.Println("Not a real retriever")
	//}


	fmt.Println("Try a session...")
	fmt.Println(session(r))

	// we use interface{} to represent "any type"
}
