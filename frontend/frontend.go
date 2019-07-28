package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Hello struct {
	Version  string
	HostName string
}

var (
	backend = os.Getenv("BACKEND")

	backendClient = http.Client{
		Timeout: time.Second * 360,
	}

	hostname, _ = os.Hostname()
)

func main() {
	if backend == "" {
		backend = "http://127.0.0.1:3000"
	}

	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		log.Println("new request", r.RemoteAddr, r.Header)
		w.Write([]byte(hostname))
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		body, status, err := httpRequest(w, r)
		if err != nil {
			return
		}

		hello := Hello{}
		jsonErr := json.Unmarshal(body, &hello)
		if jsonErr != nil {
			w.Write([]byte(err.Error()))
			return
		}

		log.Printf("recv backend json: %v", string(body))
		w.WriteHeader(status)
		w.Write(body)
	})

	http.HandleFunc("/timeout", func(w http.ResponseWriter, r *http.Request) {
		body, status, err := httpRequest(w, r)
		if err != nil {
			return
		}

		log.Printf("recv backend body: %v", string(body))
		w.WriteHeader(status)
		w.Write(body)
	})

	http.HandleFunc("/abort", func(w http.ResponseWriter, r *http.Request) {
		body, status, err := httpRequest(w, r)
		if err != nil {
			return
		}

		log.Printf("recv backend json: %v", string(body))
		w.WriteHeader(status)
		w.Write(body)
	})

	http.HandleFunc("/retry", func(w http.ResponseWriter, r *http.Request) {
		body, status, err := httpRequest(w, r)
		if err != nil {
			return
		}

		log.Printf("recv backend json: %v", string(body))
		w.WriteHeader(status)
		w.Write(body)
	})

	http.ListenAndServe(":3001", nil)
}

func httpRequest(w http.ResponseWriter, r *http.Request) ([]byte, int, error) {
	log.Println("new request", r.RemoteAddr, r.Header)

	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", backend, r.RequestURI), nil)
	req.Header.Set("User-Agent", r.UserAgent())
	req.Header.Set("Content", "Application/json")

	resp, err := backendClient.Do(req)
	if err != nil {
		w.Write([]byte(err.Error()))
		return nil, 0, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return nil, 0, err
	}

	return body, resp.StatusCode, nil
}
