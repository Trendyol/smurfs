package main

import (
	"fmt"
	"github.com/trendyol/smurfs/go/client"
	"log"
)

func main() {
	smurfs, err := client.InitializeClient(client.Options{})

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	token, err := smurfs.Auth.GetToken()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Printf("Token: %s\n", token.AccessToken)

	smurfs.Logger.Info("Micro CLI 2")
	smurfs.Logger.Error("Micro CLI 2")
	smurfs.Logger.Debug("Micro CLI 2")
	smurfs.Logger.Warn("Micro CLI 2")
	smurfs.Logger.Fatal("Micro CLI 2")

	fmt.Println("Complete it")

	err = smurfs.Close()
	if err != nil {
		log.Fatalln(err)
	}
}
