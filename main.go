package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"

	"github.com/ignite/cli/ignite/pkg/cosmosclient"
	"github.com/joho/godotenv"

	// Importing the types package of Re Prtotocol
	"github.com/jim380/Re/x/fix/types"
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

	// set your home dir for your chain manually
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	homePath := filepath.Join(home, ".red")

	cosmosOptions := []cosmosclient.Option{
		cosmosclient.WithHome(homePath),
		cosmosclient.WithKeyringBackend("test"),
		// This keyring will request a password each time it is accessed
		//cosmosclient.WithKeyringBackend("file"),
		cosmosclient.WithKeyringDir(homePath),
		cosmosclient.WithAddressPrefix(addressPrefix),
		cosmosclient.WithNodeAddress(nodeAddress),
	}

	// Create a Re Protocol client instance
	client, err := cosmosclient.New(ctx, cosmosOptions...)
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

	// Define a message to create Account
	msg := &types.MsgRegisterAccount{
		Creator:          address,
		Address:          address,
		CompanyName:      "dddd",
		Website:          "pokkdndnd",
		SocialMediaLinks: "suyyyyy",
	}

	// Broadcast a transaction from account `alice` with the message
	// to create a MsgRegisterAccount store response in txResp
	txResp, err := client.BroadcastTx(ctx, account, msg)
	if err != nil {
		log.Fatal(err)
	}

	// Print response from broadcasting a transaction
	fmt.Print("MsgRegisterAccount:\n\n")
	fmt.Println(txResp)

	// Instantiate a query client for your `Re` blockchain
	queryClient := types.NewQueryClient(client.Context())

	// Query the blockchain using the client's `AccountRegistrationAll` method
	// to get all accounts store all accounts in queryResp
	queryResp, err := queryClient.AccountRegistrationAll(ctx, &types.QueryAllAccountRegistrationRequest{})
	if err != nil {
		log.Fatal(err)
	}

	// Print response from querying all the accounts
	fmt.Print("\n\nAccountRegistrationAll:\n\n")
	fmt.Println(queryResp)

}