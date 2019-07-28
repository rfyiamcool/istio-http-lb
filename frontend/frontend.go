package main

import (
	"encoding/json"
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

func main() {
	backend := os.Getenv("BACKEND")
	if backend == "" {
		backend = "http://127.0.0.1:3000/hello"
	}

	backendClient := http.Client{
		Timeout: time.Second * 2,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("new request", r.RemoteAddr, r.Header)
		req, _ := http.NewRequest(http.MethodGet, backend, nil)
		req.Header.Set("Content", "Application/json")
		resp, err := backendClient.Do(req)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		body, _ := ioutil.ReadAll(resp.Body)
		hello := Hello{}
		jsonErr := json.Unmarshal(body, &hello)
		if jsonErr != nil {
			w.Write([]byte(err.Error()))
			return
		}

		log.Printf("recv backend json: %+v", hello)
		w.Write(body)
	})

	http.ListenAndServe(":3001", nil)
}
