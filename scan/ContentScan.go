package scan

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Gaardsholt/go-gitguardian/types"
)

func (c *ScanClient) ContentScan(payload ContentScanPayload) (*ContentScanResult, error) {

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(payload)
	if err != nil {
		return nil, err
	}

	endpoint := types.Endpoints["Scan"]
	req, err := c.client.NewRequest(endpoint.Operation, endpoint.Path, b)
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
		return &ContentScanResult{Error: &target}, fmt.Errorf("%s", target.Detail)
	}

	var target ContentScanResponse
	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&target)
	if err != nil {
		return nil, err
	}

	return &ContentScanResult{Result: target}, nil
}
