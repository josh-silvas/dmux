package circuits

import (
	"fmt"
	"github.com/josh-silvas/gonautobot/core"
	"github.com/josh-silvas/gonautobot/extras"
	"github.com/josh-silvas/gonautobot/shared"
	"github.com/josh-silvas/gonautobot/shared/nested"
	"net/http"
	"net/url"
	"time"
)

type (
	// Circuit : defines a circuit entry in Nautobot
	Circuit struct {
		ID           string             `json:"id"`
		Display      string             `json:"display"`
		URL          string             `json:"url"`
		CircuitID    string             `json:"cid"`
		Provider     nested.Provider    `json:"provider"`
		Type         nested.CircuitType `json:"type"`
		Status       shared.LabelValue  `json:"status"`
		Tenant       nested.Tenant      `json:"tenant"`
		InstallDate  string             `json:"install_date"`
		CommitRate   int                `json:"commit_rate"`
		Description  string             `json:"description"`
		TerminationA Termination        `json:"termination_a"`
		TerminationZ Termination        `json:"termination_z"`
		Comments     string             `json:"comments"`
		Created      string             `json:"created"`
		LastUpdated  time.Time          `json:"last_updated"`
		Tags         []extras.Tag       `json:"tags"`
		NotesURL     string             `json:"notes_url"`
		CustomFields map[string]any     `json:"custom_fields"`
	}

	// CircuitRequest : Models a new circuit entry in Nautobot.
	CircuitRequest struct {
		CircuitID    string         `json:"cid"`
		Provider     string         `json:"provider"`
		Type         string         `json:"circuit_type"`
		Status       string         `json:"status"`
		Tenant       string         `json:"tenant,omitempty"`
		InstallDate  string         `json:"install_date,omitempty"`
		CommitRate   int            `json:"commit_rate,omitempty"`
		Description  string         `json:"description,omitempty"`
		TerminationA Termination    `json:"termination_a,omitempty"`
		TerminationZ Termination    `json:"termination_z,omitempty"`
		Comments     string         `json:"comments,omitempty"`
		Created      string         `json:"created,omitempty"`
		Tags         []extras.Tag   `json:"tags,omitempty"`
		CustomFields map[string]any `json:"custom_fields,omitempty"`
	}
)

// CircuitGet : Go function to process requests for the endpoint: /api/circuits/circuits/
//
// https://demo.nautobot.com/api/docs/#/circuits/circuits_circuits_list
func (c *Client) CircuitGet(uuid string) (Circuit, error) {
	req, err := c.Request(http.MethodGet, fmt.Sprintf("circuits/circuits/%s/", uuid), nil, nil)
	if err != nil {
		return Circuit{}, err
	}

	ret := new(Circuit)
	return *ret, c.UnmarshalDo(req, ret)
}

// CircuitFilter : Go function to process requests for the endpoint: /api/circuits/circuits/
//
// https://demo.nautobot.com/api/docs/#/circuits/circuits_circuits_list
func (c *Client) CircuitFilter(q *url.Values) ([]Circuit, error) {
	resp := make([]Circuit, 0)
	return resp, core.Paginate[Circuit](c.Client, "circuits/circuits/", q, &resp)
}

// CircuitAll : Go function to process requests for the endpoint: /api/circuits/circuits/
//
// https://demo.nautobot.com/api/docs/#/circuits/circuits_circuits_list
func (c *Client) CircuitAll() ([]Circuit, error) {
	resp := make([]Circuit, 0)
	return resp, core.Paginate[Circuit](c.Client, "circuits/circuits/", nil, &resp)
}

// CircuitDelete : Go function to process requests for the endpoint: /api/circuits/circuits/
//
// https://demo.nautobot.com/api/docs/#/circuits/circuits_circuits_destroy
func (c *Client) CircuitDelete(uuid string) error {
	req, err := c.Request(http.MethodDelete, fmt.Sprintf("circuits/circuits/%s/", uuid), nil, nil)
	if err != nil {
		return err
	}
	return c.UnmarshalDo(req, nil)
}

// CircuitCreate : Go function to process requests for the endpoint: /api/circuits/circuits/
//
// https://demo.nautobot.com/api/docs/#/circuits/circuits_circuits_create
func (c *Client) CircuitCreate(r CircuitRequest) (Circuit, error) {
	req, err := c.Request(http.MethodPost, "circuits/circuits/", r, nil)
	if err != nil {
		return Circuit{}, err
	}
	var ret Circuit
	return ret, c.UnmarshalDo(req, &ret)
}
