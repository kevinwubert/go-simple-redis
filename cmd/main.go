package main

import (
	"fmt"

	"github.com/kevinwubert/go-simple-redis/pkg/redisclient"
)

func main() {
	err := redisclient.Main()
	if err != nil {
		fmt.Println(err)
	}
}
