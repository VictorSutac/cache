package main

import (
	"fmt"

	"myproject/cache"
)

func main() {
	cache := cache.New()

	cache.Set("userId", 42)

	userId := cache.Get("userId")
	fmt.Println("userId", userId)

	cache.Delete("userId")

	userId = cache.Get("userId")
	fmt.Println("userId after delet: ", userId) // Output: userId after delete: <nil>
}
