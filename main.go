package main

import (
	"fmt"
	"net/http"
)


func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello from OpenShift Dev Spaces!!!")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, header := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, header)
		}
	}
}

func main() {
	
	http.HandleFunc("hello", hello)
	http.HandleFunc("headers", headers)
	http.ListenAndServe(":8080", nil)
}
