package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

func handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s", r.URL.String())
		h.ServeHTTP(w, r)
	})
}

func main() {
	var port int
	flag.IntVar(&port, "port", 8000, "default port for http server")
	flag.Parse()
	p := strconv.Itoa(port)

	http.Handle("/", handler(http.FileServer(http.Dir("."))))
	log.Printf("Starting HTTP on 127.0.0.1 port %s ...\n", p)
	log.Fatal(http.ListenAndServe(":"+p, nil))
}
