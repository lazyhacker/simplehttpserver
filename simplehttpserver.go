package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

func LogRequestHandler(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s", r.URL.String())
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Credentials", "false")
		h.ServeHTTP(w, r)
	})
}

func main() {
	var port int
	var host string
	flag.IntVar(&port, "port", 8000, "default port for http server")
	flag.StringVar(&host, "host", "", "default will listen to all interfaces")
	flag.Parse()
	p := strconv.Itoa(port)

	http.Handle("/", LogRequestHandler(http.FileServer(http.Dir("."))))

	log.Printf("Starting HTTP on %s:%s ...\n", host, p)
	log.Fatal(http.ListenAndServe(host+":"+p, nil))
}
