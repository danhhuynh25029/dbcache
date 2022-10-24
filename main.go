package main

import (
	"sync"
	"fmt"
)
type Cache struct{
	cache map[string]*result
	sync.Mutex
}
type result struct{
	value []byte
	err error
}
type Func func() ([]byte,error)

func NewCache() *Cache{
	return &Cache{cache : make(map[string]*result)}
}
func (c *Cache) Get(key string)([]byte,error){
	res,ok := c.cache[key]
	if !ok {
		res := &result{}
		res.value,res.err = []byte{},nil
		c.cache[key] = res
	}
	return res.value,res.err
}
func (c *Cache) Set(key string,value string)(error){
	r := result{[]byte(value),nil}
	c.cache[key] = &r
	return nil
}
func main(){
	cache := NewCache()
	cache.Set("1","124")
	value,err := cache.Get("1")
	if err != nil{
		fmt.Println("error get value from redis")
	}
	fmt.Println("value is ",string(value))
}
