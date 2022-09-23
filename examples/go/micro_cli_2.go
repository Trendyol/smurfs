package main

import "github.com/trendyol/smurfs/go/client"

func main() {
	host := "http://localhost:8080"
	smurfs, err := client.InitializeClient(client.Options{
		HostAddress: &host,
	})
	if err != nil {
		panic(err)
	}

	smurfs.Logger.Info("Micro CLI 1")
}
