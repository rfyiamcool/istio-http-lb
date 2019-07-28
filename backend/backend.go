package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Hello struct {
	Version  string
	HostName string
}

var (
	hostname, _ = os.Hostname()
	version     = os.Getenv("VERSION")
)

func main() {
	hello := Hello{
		Version:  version,
		HostName: hostname,
	}
	js, _ := json.Marshal(hello)

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Println("new request", r.RemoteAddr, r.Header)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})

	http.ListenAndServe(":3000", nil)
}
