package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"
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
	port := 8880
	if len(os.Args) >= 2 {
		rawPort := os.Args[1]
		if num, err := strconv.Atoi(rawPort); err != nil {
			fmt.Printf("%q is not a port number: %v\n", rawPort, err)
			os.Exit(1)
		} else {
			port = num
		}
	}
	fmt.Println("listening on port:", port)

	http.HandleFunc("/", hello)
	http.HandleFunc("/print", printRequest)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		fmt.Printf("failed to listen on port %d, got error: %v\n", port, err)
		os.Exit(2)
	}
}
