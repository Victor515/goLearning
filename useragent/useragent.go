package main

import (
	"log"
	"github.com/EDDYCJY/fake-useragent"
	"net/http"
)

func main() {
	client := http.Client{}
	i := 0
	for i < 10 {
		useragent := browser.Random()
		//log.Printf("User Agent: %v\n", useragent)
		req, err := http.NewRequest(http.MethodGet, "http://www.zhenai.com/zhenghun", nil)
		req.Header.Set("User-Agent", useragent)
		if err != nil{
			panic(err)
			continue
		}
		response, err := client.Do(req)
		log.Printf("Header User Agent: %v, Response Status Code: %v\n", req.UserAgent(), response.StatusCode)
		i++
	}

}
