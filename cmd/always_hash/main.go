package main

import (
	"fmt"
	"log"

	"wallet-op/cache"
)

func main() {
	// init cache
	if err := cache.InitWordCache("./bip-39/word.list"); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(cache.ListWord())
}
