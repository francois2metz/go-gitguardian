package incidents

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Gaardsholt/go-gitguardian/types"
)

type ListSecretMembersOptions struct {
	Cursor  string `json:"cursor"`   // Pagination cursor.
	PerPage *int   `json:"per_page"` // Number of items to list per page.	[ 1 .. 100 ]
}

type IncidentListSecretMembersResult struct {
	Result []IncidentListSecretMembersResponse `json:"result"`
	Error  *Error                              `json:"error"`
}

type IncidentListSecretMembersResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func (c *IncidentsClient) ListSecretMembers(IncidentId int, lo ListSecretMembersOptions) (*IncidentListSecretMembersResult, error) {
	endpoint := types.Endpoints["ListSecretMembers"]
	req, err := c.client.NewRequest(endpoint.Operation, fmt.Sprintf(endpoint.Path, IncidentId), nil)
	if err != nil {
		return nil, err
	}

	// Add query parameters
	q := req.URL.Query()

	if lo.PerPage != nil {
		if !(*lo.PerPage >= 1 && *lo.PerPage <= 100) {
			return nil, fmt.Errorf("PerPage must be between 1 and 100")
		}
		q.Add("per_page", strconv.Itoa(*lo.PerPage))
	}

	req.URL.RawQuery = q.Encode()

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
		return &IncidentListSecretMembersResult{Error: &target}, fmt.Errorf("%s", target.Detail)
	}

	var target []IncidentListSecretMembersResponse
	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&target)
	if err != nil {
		return nil, err
	}

	return &IncidentListSecretMembersResult{Result: target}, nil
}
