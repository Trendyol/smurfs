package main

import (
	"fmt"
	"github.com/trendyol/smurfs/go/host"
)

func main() {
	_, err := host.InitializeHost(host.Options{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Host CLI")
}
