package main

import (
	"dbcache/cache"
	"fmt"
)

func main() {
	cache := cache.NewCache()
	cache.Set("1", "124")
	value, err := cache.Get("1")
	if err != nil {
		fmt.Println("error get value from redis")
	}
	fmt.Println("value is ", string(value))
}
