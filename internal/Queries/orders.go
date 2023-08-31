package queries

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	fixstruct "github.com/jim380/re_client/internal/FIXStruct"
)

func fetchOrders(address string) (*fixstruct.Response, error) {
	url := fmt.Sprintf("http://localhost:1317/jim380/Re/fix/orders/%s", address)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var orders fixstruct.Response
	err = json.Unmarshal(body, &orders)
	if err != nil {
		return nil, err
	}

	return &orders, nil
}
