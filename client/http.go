package client

import (
	"dbcache/cache"
	"encoding/json"

	"log"
	"net/http"
)

var c = cache.NewCache()

type data struct {
	Key   string
	Value string
}

func NewServer(port string) {
	http.HandleFunc("/set", set)
	http.HandleFunc("/get", get)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("cannot start server : %v", err)
	} else {
		log.Fatalf("server is running with port : %s", port)
	}
}
func set(w http.ResponseWriter, r *http.Request) {
	var d data
	if r.Method == "POST" {
		err := json.NewDecoder(r.Body).Decode(&d)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = c.Set(d.Key, d.Value)
		if err != nil {
			log.Fatalf("cannot set value into cache : %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}
func get(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		query := r.URL.Query()
		res, err := c.Get(query.Get("key"))
		if err != nil {
			log.Fatalf("cannot")
		} else {
			w.Write(res)
		}
	}
}
