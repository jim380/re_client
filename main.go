package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/ignite/cli/ignite/pkg/cosmosclient"
	"github.com/joho/godotenv"
)

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	addressPrefix := "re"

	// Get node address from environment variable or default to localhost:26657
	nodeAddress := os.Getenv("RE_NODE_ADDRESS")
	if nodeAddress == "" {
		nodeAddress = "http://0.0.0.0:26657"
	}

	_, err = url.Parse(nodeAddress)
	if err != nil {
		log.Fatalf("Invalid node address: %s", nodeAddress)
	}

	// Create a Re Protocol client instance
	client, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix(addressPrefix), cosmosclient.WithNodeAddress(nodeAddress))
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

	address, err := account.Address(addressPrefix)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Account address:", address)

}
