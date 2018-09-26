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

	proxyURL, err := urllib.Parse(*url)
	if err != nil {
		log.Println(err)
		return
	}
	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}

	resp, err := client.Get("http://tomocy:tomocy@github.com")
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
