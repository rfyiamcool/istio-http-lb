package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Hello struct {
	Version  string
	HostName string
}

var (
	hostname, _ = os.Hostname()
	version     = os.Getenv("VERSION")
	timeout     = 60 * time.Second
)

func main() {
	hello := Hello{
		Version:  version,
		HostName: hostname,
	}
	js, _ := json.Marshal(hello)

	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		log.Println("new request", r.RemoteAddr, r.Header)
		w.Write([]byte(hostname))
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Println("new request", r.RemoteAddr, r.Header)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})

	http.HandleFunc("/timeout", func(w http.ResponseWriter, r *http.Request) {
		log.Println("new request", r.RemoteAddr, r.Header)
		time.Sleep(timeout)
		w.Write([]byte("ok"))
	})

	http.HandleFunc("/abort", func(w http.ResponseWriter, r *http.Request) {
		log.Println("new request", r.RemoteAddr, r.Header)
		w.WriteHeader(555)
		w.Write([]byte("active abort error"))
	})

	http.HandleFunc("/retry", func(w http.ResponseWriter, r *http.Request) {
		log.Println("new request", r.RemoteAddr, r.Header)
		n := rand.Intn(100)
		if n > 70 {
			w.WriteHeader(555)
			w.Write([]byte("active raise error"))
			return
		}

		w.Write([]byte("ok"))
	})

	http.ListenAndServe(":3000", nil)
}
