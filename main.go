package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, current time is %v\n", time.Now().Local())
}

var reqCount = 0

func printRequest(w http.ResponseWriter, req *http.Request) {
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println(err)
	}

	reqCount++
	dumpStr := fmt.Sprintf("[Request #%d - %v]\n%s", reqCount, time.Now().Local(), string(requestDump))
	fmt.Printf("\n%s\n", dumpStr)
	fmt.Fprintf(w, dumpStr)
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/print", printRequest)
	http.ListenAndServe(":8090", nil)
}
