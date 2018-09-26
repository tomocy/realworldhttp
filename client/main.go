package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	urllib "net/url"
)

func main() {
	url := flag.String("url", "http://localhost:8080", "a url for request")
	flag.Parse()

	jar, _ := cookiejar.New(nil)
	cookieURL, err := urllib.Parse(*url)
	if err != nil {
		log.Println(err)
		return
	}
	jar.SetCookies(cookieURL, []*http.Cookie{
		{Name: "name", Value: "tomocy"},
	})

	client := http.Client{
		Jar: jar,
	}

	resp, err := client.Get(*url)
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
