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

func FetchOrders(address string) (*fixstruct.OrderResponse, error) {
	// Read API_URL from .env file
	ordersURL := os.Getenv("ORDERS_URL")
	if ordersURL == "" {
		return nil, errors.New("ORDERS_URL not found in .env file")
	}

	url := fmt.Sprintf("%s/%s", ordersURL, address)
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
