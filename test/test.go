package main

import (
	"dbcache/infra"
	"fmt"
)

func main() {
	db := infra.NewDBCache(":8080")
	db.Set("1", "10")
	fmt.Println(db.Get("1"))
}
