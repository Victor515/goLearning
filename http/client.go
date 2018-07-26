package main

import (
	"net/http"
	"net/http/httputil"
	"fmt"
)

func main() {
	// visit mobile site
	req, err := http.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
	//req.Header.Add("User-Agent",
	//	"Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")

	// send custom header
	//resp, err := http.DefaultClient.Do(req)

	//check whether we have redirected
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect: ", req)
			return nil
		},
	}
	resp, err := client.Do(req)

	// send get request
	//resp, err := http.Get("http://www.imooc.com")
	if err != nil{
		panic(err)
	}

	// close the response body
	defer resp.Body.Close()

	// dump the responses
	content, err := httputil.DumpResponse(resp, true)

	if err != nil{
		panic(err)
	}
	fmt.Printf("%s", content)
}
