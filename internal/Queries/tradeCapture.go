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

// FetchTradeCapture retrieves all trade capture reports for a given chainID with the corresponding address.
// It sends a GET request to the TRADECAPTURE_URL with the specified chainID and address,
// reads the response body, and unmarshals it into the TradeCaptureResponse struct.
// The function returns the trade capture reports and any error encountered.
func FetchTradeCapture(chainID, address string) (*fixstruct.TradeCaptureResponse, error) {
	// Read API_URL from .env file
	tradeCaptureURL := os.Getenv("TRADECAPTURE_URL")
	if tradeCaptureURL == "" {
		return nil, errors.New("TRADECAPTURE_URL not found in .env file")
	}

	url := fmt.Sprintf("%s/%s/%s", tradeCaptureURL, chainID, address)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var tradeCaptures fixstruct.TradeCaptureResponse
	err = json.Unmarshal(body, &tradeCaptures)
	if err != nil {
		return nil, err
	}

	return &tradeCaptures, nil
}

// FetchAllTradeCapture retrieves all trade captures for a given chainID.
// It sends a GET request to the TRADECAPTURE_URL with the specified chainID,
// reads the response body, and unmarshals it into the TradeCaptureResponse struct.
// The function returns the trade capture reports and any error encountered.
func FetchAllTradeCapture(chainID string) (*fixstruct.TradeCaptureResponse, error) {
	// Read API_URL from .env file
	tradeCaptureURL := os.Getenv("TRADECAPTURE_URL")
	if tradeCaptureURL == "" {
		return nil, errors.New("TRADECAPTURE_URL not found in .env file")
	}

	url := fmt.Sprintf("%s/%s", tradeCaptureURL, chainID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var tradeCaptures fixstruct.TradeCaptureResponse
	err = json.Unmarshal(body, &tradeCaptures)
	if err != nil {
		return nil, err
	}

	return &tradeCaptures, nil
}
