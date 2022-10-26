package client

import (
	"log"
	"net/http"
)

func NewServer(port string) {
	http.HandleFunc("/set", set)
	http.HandleFunc("/get", get)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("cannot start server : %v", err)
	}
}
func set(w http.ResponseWriter, r *http.Request) {

}
func get(w http.ResponseWriter, r *http.Request) {

}
