package main

import (
	"fmt"
	"github.com/trendyol/smurfs/go/client"
	"log"
)

func main() {
	fmt.Println("Starting micro2")
	host := "localhost:50051"
	smurfs, err := client.InitializeClient(client.Options{
		HostAddress: &host,
	})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	smurfs.Logger.Info("Micro CLI 1")
	smurfs.Logger.Info("Micro CLI 1")
	smurfs.Logger.Info("Micro CLI 1")
	smurfs.Logger.Info("Micro CLI 1")
	smurfs.Logger.Info("Micro CLI 1")
	smurfs.Logger.Info("Micro CLI 1")

	fmt.Println("Complete it")

	err = smurfs.Close()
	if err != nil {
		log.Fatalln(err)
	}
}
