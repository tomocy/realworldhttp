package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	http.HandleFunc("/", echoHandler)
	http.HandleFunc("/digest", digestAuthenticationHandler)
	addr := ":8080"
	log.Printf("start to listen and serve at %s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Println(err)
	}
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	dumpedReq, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	fmt.Println(string(dumpedReq))
}

func digestAuthenticationHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := r.Header["Authorization"]; ok {
		fmt.Fprintln(w, "You are authenticated!")
		return
	}

	addHeaderForClientToTryDigestAuthentication(w)
}

func addHeaderForClientToTryDigestAuthentication(w http.ResponseWriter) {
	w.Header().Add("WWW-Authenticate", generateHeaderValueForDigestAuthentication())
	w.WriteHeader(http.StatusUnauthorized)
}

func generateHeaderValueForDigestAuthentication() string {
	realm := "Need to be authorized"
	algorithm := "MD5"
	qop := "auth"
	nonce := "aaaaaaaaaaaaaaa"

	return fmt.Sprintf("Digest realm=%s, algorithm=%s, qop=%s, nonce=%s", realm, algorithm, qop, nonce)
}
