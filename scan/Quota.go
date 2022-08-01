package scan

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Gaardsholt/go-gitguardian/types"
)

func (c *ScanClient) Quota() (*QuotaResult, error) {
	endpoint := types.Endpoints["ScanQuotas"]
	req, err := c.client.NewRequest(endpoint.Operation, endpoint.Path, nil)
	if err != nil {
		return nil, err
	}

	r, err := c.client.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		var target Error
		decode := json.NewDecoder(r.Body)
		err = decode.Decode(&target)
		if err != nil {
			return nil, err
		}
		return &QuotaResult{Error: &target}, fmt.Errorf("%s", target.Detail)
	}

	var target QuotaResponse
	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&target)
	if err != nil {
		return nil, err
	}

	return &QuotaResult{Result: target}, nil
}
