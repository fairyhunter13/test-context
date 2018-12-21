package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

const timeout = 5

func main() {
	http.HandleFunc("/thanos", thanosHandler)
	log.Println("Server running: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func thanosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*timeout)
		defer cancel()

		go func() {
			time.Sleep(5 * time.Second)
			w.Write([]byte("Done doing work!"))
		}()

		select {
		case <-ctx.Done():
			timeoutHandler(w, r)
		}
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	return
}

func timeoutHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Timeout", http.StatusRequestTimeout)
}
