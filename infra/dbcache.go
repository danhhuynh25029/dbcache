package infra

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type DBCache interface {
	Set(key string, value string) error
	Get(key string) (string, error)
}
type dbCache struct {
	url  string
	port string
}

func NewDBCache(port string) DBCache {
	return &dbCache{
		url:  fmt.Sprintf("http://localhost%s", port),
		port: port,
	}
}
func (d *dbCache) Set(key, value string) error {
	postBody, _ := json.Marshal(map[string]string{
		"key":   key,
		"value": value,
	})
	requestBody := bytes.NewBuffer(postBody)
	fmt.Println(d.url + "/set")
	resp, err := http.Post(d.url+"/set", "application/json", requestBody)
	if err != nil {
		log.Fatalf("cannot set value : %v", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Print(string(body))
	return err
}
func (d *dbCache) Get(key string) (string, error) {
	res, err := http.Get(d.url + "/get" + "?key=" + key)
	if err != nil {
		log.Fatalf("cannot get value : %v", err)
		return "", err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(body), nil
}
