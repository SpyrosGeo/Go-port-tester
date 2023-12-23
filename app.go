package main

import (
	"flag"
	"net/http"
)

func main() {
	var port string
	flag.StringVar(&port, "p", "22", "specify port")
	flag.Parse()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	})
	// err:=http.ListenAndServe()
}
