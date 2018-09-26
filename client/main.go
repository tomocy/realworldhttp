package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	url := flag.String("url", "http://localhost:8080", "a url for request")
	flag.Parse()

	file, err := os.Open("contents.txt")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	resp, err := http.Post(*url, "text/plain", file)
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
