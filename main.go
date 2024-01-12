package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var maxTimeout *int

func main() {
	port := flag.String("p", "8000", "[thttp -p 8000]")
	maxTimeout = flag.Int("t", 5000, "[thttp -t 5000]")
	flag.Parse()
	log.Printf("Listening on port %s", *port)

	// handle /timeout/* with random timeout
	http.HandleFunc("/timeout/", handleTimeout)
	// handle all other requests
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func handleTimeout(w http.ResponseWriter, r *http.Request) {
	// random timeout between 0 and maxTimeout
	time.Sleep(time.Duration(rand.Intn(*maxTimeout)) * time.Millisecond)

	handle(w, r)
}

func handle(w http.ResponseWriter, r *http.Request) {
	var b []byte
	var err error
	var reqStr string

	if r.Body != nil {
		b, err = io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			return
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Println(err)
			}
		}(r.Body)
	}

	reqStr = fmt.Sprintf("%s %s %s\n", r.Method, r.RequestURI, r.Proto)
	reqStr += fmt.Sprintf("Host: %s\n", r.Host)
	for k, v := range r.Header {
		reqStr += fmt.Sprintf("%s: %s\n", k, strings.Join(v[:], ","))
	}
	reqStr += fmt.Sprintf("Body:\n%s", b)

	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintf(w, reqStr)
	if err != nil {
		return
	}

	log.Println(reqStr)
}
