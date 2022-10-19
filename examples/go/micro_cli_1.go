package main

import (
	"fmt"
	"github.com/trendyol/smurfs/go/client"
	"log"
)

func main() {
	fmt.Println("Starting micro1")
	host := "localhost:50051"
	smurfs, err := client.InitializeClient(client.Options{
		HostAddress: &host,
	})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	i := 0
	for {
		smurfs.Logger.Info(fmt.Sprintf("Micro CLI %d", i))
		i++

		if i > 100 {
			break
		}
	}

	err = smurfs.Close()
	if err != nil {
		log.Fatalln(err)
	}
}
