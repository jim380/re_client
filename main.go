package main

import (
	// Importing the general purpose Cosmos blockchain client
	"context"
	"log"

	"github.com/ignite/cli/ignite/pkg/cosmosclient"
)

func main() {
	ctx := context.Background()
	addressPrefix := "re"

	// Create a Re Protocol client instance
	client, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix(addressPrefix))
	if err != nil {
		log.Fatal(err)
	}

	// Account `alice` was initialized
	accountName := "alice"

	// Get account from the keyring
	account, err := client.Account(accountName)
	if err != nil {
		log.Fatal(err)
	}

	addr, err := account.Address(addressPrefix)
	if err != nil {
		log.Fatal(err)
	}
	println(addr)
}
