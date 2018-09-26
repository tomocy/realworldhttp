package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	transport := http.Transport{}
	transport.RegisterProtocol("file", http.NewFileTransport(http.Dir(".")))
	client := http.Client{Transport: &transport}

	resp, err := client.Get("file://./main.go")
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
