package queries

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	fixstruct "github.com/jim380/re_client/internal/FIXStruct"
)

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
}

// FetchOrders retrieves all orders for a given chainID with the corresponding address.
// It sends a GET request to the ORDERS_URL with the specified chainID and address,
// reads the response body, and unmarshals it into the OrderResponse struct.
// The function returns the orders and any error encountered.
func FetchOrders(chainID, address string) (*fixstruct.OrderResponse, error) {
	// Read API_URL from .env file
	ordersURL := os.Getenv("ORDERS_URL")
	if ordersURL == "" {
		return nil, errors.New("ORDERS_URL not found in .env file")
	}

	url := fmt.Sprintf("%s/%s/%s", ordersURL, chainID, address)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var orders fixstruct.OrderResponse
	err = json.Unmarshal(body, &orders)
	if err != nil {
		return nil, err
	}

	return &orders, nil
}

// FetchAllOrders retrieves all orders for a given chainID.
// It sends a GET request to the ORDERS_URL with the specified chainID,
// reads the response body, and unmarshals it into the OrderResponse struct.
// The function returns the orders and any error encountered.
func FetchAllOrders(chainID string) (*fixstruct.OrderResponse, error) {
	// Read API_URL from .env file
	ordersURL := os.Getenv("ORDERS_URL")
	if ordersURL == "" {
		return nil, errors.New("ORDERS_URL not found in .env file")
	}

	url := fmt.Sprintf("%s/%s", ordersURL, chainID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var orders fixstruct.OrderResponse
	err = json.Unmarshal(body, &orders)
	if err != nil {
		return nil, err
	}

	return &orders, nil
}
