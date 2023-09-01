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

func FetchExecutionReports(address string) (*fixstruct.OrdersExecutionReportResponse, error) {
	// Read API_URL from .env file
	executionReportURL := os.Getenv("EXECUTIONREPORT_URL")
	if executionReportURL == "" {
		return nil, errors.New("ORDERS_URL not found in .env file")
	}

	url := fmt.Sprintf("%s/%s", executionReportURL, address)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ordersExecutionReports fixstruct.OrdersExecutionReportResponse
	err = json.Unmarshal(body, &ordersExecutionReports)
	if err != nil {
		return nil, err
	}

	return &ordersExecutionReports, nil
}

func FetchAllExecutionReports() (*fixstruct.OrdersExecutionReportResponse, error) {
	// Read API_URL from .env file
	executionReportURL := os.Getenv("EXECUTIONREPORT_URL")
	if executionReportURL == "" {
		return nil, errors.New("ORDERS_URL not found in .env file")
	}

	resp, err := http.Get(executionReportURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ordersExecutionReports fixstruct.OrdersExecutionReportResponse
	err = json.Unmarshal(body, &ordersExecutionReports)
	if err != nil {
		return nil, err
	}

	return &ordersExecutionReports, nil
}
