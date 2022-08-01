package invitations

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Gaardsholt/go-gitguardian/types"
)

type InvitationsResendResult struct {
	Error *Error `json:"error"`
}

func (c *InvitationsClient) Resend(InvitationId int) (*InvitationsResendResult, error) {
	endpoint := types.Endpoints["InvitationsResend"]
	req, err := c.client.NewRequest(endpoint.Operation, fmt.Sprintf(endpoint.Path, InvitationId), nil)
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
		return &InvitationsResendResult{Error: &target}, fmt.Errorf("%s", target.Detail)
	}

	return nil, nil
}
