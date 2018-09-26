package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	urllib "net/url"
)

func main() {
	url := flag.String("url", "http://localhost:8080", "a url for request")
	flag.Parse()

	values := urllib.Values{
		"query": {"hello world"},
	}

	resp, err := http.Get(*url + "?" + values.Encode())
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(body))
}
