package main

import (
	"context"
	"net/http"
	"time"
)

const timeout = 5

func main() {
	http.HandleFunc("/thanos", thanosHandler)
}

func thanosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ctx, _ := context.WithTimeout(context.Background(), time.Second*timeout)

		go func() {

		}()

		select {
		case <-ctx.Done():

		}
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	return
}

func timoutHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Timeout", http.StatusRequestTimeout)
}
