package main

import (
	"awesomeProject/cachepractice"
	"fmt"
	"time"
)

func main() {
	cache := cachepractice.New(5*time.Minute, 10*time.Minute)

	mapa := make(map[string]interface{})
	mapa["name"] = "Ivan"
	mapa["age"] = 18
	mapa["email"] = "testcache@g.co"

	cache.Set("consumer1", mapa, time.Minute)

	i, err := cache.Get("consumer1")
	if err != nil {
		fmt.Printf("failed to get consumer1: %w \n", err)
	}
	fmt.Println(i)

	_, err = cache.Get("consumer2")
	if err != nil {
		fmt.Printf("failed to get consumer2: %w", err)
	}
}
