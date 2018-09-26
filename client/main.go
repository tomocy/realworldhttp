package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
)

func main() {
	url := flag.String("url", "http://localhost:8080", "a url for request")
	flag.Parse()

	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	h := make(textproto.MIMEHeader)
	h.Set("Content-Type", "text/plain")
	h.Set("Content-Disposition", `form-data; name="message": filename="contents.txt"`)
	fileWriter, err := w.CreatePart(h)
	if err != nil {
		log.Println(err)
		return
	}

	file, err := os.Open("contents.txt")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	io.Copy(fileWriter, file)
	w.Close()

	resp, err := http.Post(*url, w.FormDataContentType(), &b)
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
