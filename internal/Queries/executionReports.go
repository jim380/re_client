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

// FetchExecutionReports retrieves execution reports for a given chainID and address.
// It sends a GET request to the EXECUTIONREPORT_URL with the specified chainID and address,
// reads the response body, and unmarshals it into the OrdersExecutionReportResponse struct.
// The function returns the execution reports and any error encountered.
func FetchExecutionReports(chainID, address string) (*fixstruct.OrdersExecutionReportResponse, error) {
	// Read API_URL from .env file
	executionReportURL := os.Getenv("EXECUTIONREPORT_URL")
	if executionReportURL == "" {
		return nil, errors.New("EXECUTIONREPORT_URL not found in .env file")
	}

	// Construct the URL with the chainID and address
	url := fmt.Sprintf("%s/%s/%s", executionReportURL, chainID, address)

	// Send a GET request to the URL
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal the response body into the OrdersExecutionReportResponse struct
	var ordersExecutionReports fixstruct.OrdersExecutionReportResponse
	err = json.Unmarshal(body, &ordersExecutionReports)
	if err != nil {
		return nil, err
	}

	return &ordersExecutionReports, nil
}

// FetchAllExecutionReports retrieves all execution reports for a given chainID.
// It sends a GET request to the EXECUTIONREPORT_URL with the specified chainID,
// reads the response body, and unmarshals it into the OrdersExecutionReportResponse struct.
// The function returns the execution reports and any error encountered.
func FetchAllExecutionReports(chainID string) (*fixstruct.OrdersExecutionReportResponse, error) {
	// Read API_URL from .env file
	executionReportURL := os.Getenv("EXECUTIONREPORT_URL")
	if executionReportURL == "" {
		return nil, errors.New("EXECUTIONREPORT_URL not found in .env file")
	}

	// Construct the URL with the chainID
	url := fmt.Sprintf("%s/%s", executionReportURL, chainID)

	// Send a GET request to the URL
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal the response body into the OrdersExecutionReportResponse struct
	var ordersExecutionReports fixstruct.OrdersExecutionReportResponse
	err = json.Unmarshal(body, &ordersExecutionReports)
	if err != nil {
		return nil, err
	}

	return &ordersExecutionReports, nil
}
