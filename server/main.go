package main

import (
	"dbcache/client"
	"sync"
)

func main() {
	port := ":8080"
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		client.NewServer(port)
		wg.Done()
	}()
	wg.Wait()
}
