package nautobot_v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Device : defines a device as represented in Nautobot
type Device struct {
	ID               string                 `json:"id"`
	AssetTag         string                 `json:"asset_tag"`
	Comments         string                 `json:"comments" datastore:",noindex"`
	ConfigContext    map[string]interface{} `json:"config_context"`
	Created          string                 `json:"created"`
	CustomFields     map[string]interface{} `json:"custom_fields"`
	DeviceRole       *DeviceRole            `json:"device_role"`
	DeviceType       *NestedDeviceType      `json:"device_type"`
	Display          string                 `json:"display"`
	Face             *LabelValue            `json:"face"`
	LastUpdated      string                 `json:"last_updated"`
	LocalContextData map[string]interface{} `json:"local_context_data"`
	Name             string                 `json:"name"`
	NotesURL         string                 `json:"notes_url"`
	ParentDevice     *struct {
		ID      string `json:"id"`
		Display string `json:"display"`
		Name    string `json:"name"`
		URL     string `json:"url"`
	} `json:"parent_device"`
	Platform   *Platform   `json:"platform"`
	Position   *int        `json:"position"`
	PrimaryIP  *IPAddress  `json:"primary_ip"`
	PrimaryIP4 *IPAddress  `json:"primary_ip4"`
	PrimaryIP6 *IPAddress  `json:"primary_ip6"`
	Rack       *Rack       `json:"rack"`
	Serial     string      `json:"serial"`
	Site       *Site       `json:"site"`
	Status     *LabelValue `json:"status"`
	Tags       []Tag       `json:"tags"`
	Tenant     *Tenant     `json:"tenant"`
	URL        string      `json:"url"`
	VCPosition *int        `json:"vc_position"`
	VCPriority *int        `json:"vc_priority"`
}

// GetDevice : method for API endpoint /api/dcim/devices/:id/
//
// https://demo.nautobot.com/api/docs/#/dcim/dcim_devices_retrieve
func (c *Client) GetDevice(uuid string, query *url.Values) (*Device, error) {
	req, err := c.Request(
		http.MethodGet,
		fmt.Sprintf("api/dcim/devices/%s/", url.PathEscape(uuid)),
		nil,
		query,
	)
	if err != nil {
		return nil, err
	}

	ret := new(Device)
	err = c.Do(req, ret)
	return ret, err
}

// GetDevices : method for API endpoint /api/dcim/devices/
//
// https://demo.nautobot.com/api/docs/#/dcim/dcim_devices_list
func (c *Client) GetDevices(query *url.Values) ([]Device, error) {
	devices := make([]Device, 0)
	offset := 0
	if query == nil {
		query = &url.Values{}
	}
	for {
		query.Set("offset", fmt.Sprintf("%d", offset))
		req, err := c.Request(http.MethodGet, "api/dcim/devices/", nil, query)
		if err != nil {
			return nil, err
		}

		resp := new(rawListResponse)
		ret := make([]Device, 0)
		if err = c.Do(req, resp); err != nil {
			return ret, err
		}

		if err = json.Unmarshal(resp.Results, &ret); err != nil {
			return devices, fmt.Errorf("error decoding results field: %w", err)
		}
		devices = append(devices, ret...)
		if resp.Count <= len(devices) {
			break
		}
		offset += 50
	}

	return devices, nil
}
