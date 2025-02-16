package dcim

import (
	"errors"
	"fmt"
	"github.com/josh-silvas/gonautobot/core"
	"github.com/josh-silvas/gonautobot/extras"
	"github.com/josh-silvas/gonautobot/shared"
	"github.com/josh-silvas/gonautobot/shared/nested"
	"net/http"
	"net/url"
)

type (
	// Device : Data structure to represent a device record in Nautobot.
	Device struct {
		ID                 string                        `json:"id"`
		AssetTag           string                        `json:"asset_tag"`
		Cluster            *nested.VirtualizationCluster `json:"cluster"`
		Comments           string                        `json:"comments" datastore:",noindex"`
		ConfigContext      map[string]interface{}        `json:"config_context"`
		Created            string                        `json:"created"`
		CustomFields       map[string]interface{}        `json:"custom_fields"`
		DeviceRole         *nested.DeviceRole            `json:"device_role"`
		DeviceType         *nested.DeviceType            `json:"device_type"`
		Display            string                        `json:"display"`
		Face               *shared.LabelValue            `json:"face"`
		LastUpdated        string                        `json:"last_updated"`
		LocalContextSchema *nested.ConfigContextSchema   `json:"local_context_schema"`
		LocalContextData   map[string]interface{}        `json:"local_context_data"`
		Location           *nested.Location              `json:"location"`
		Name               string                        `json:"name"`
		NotesURL           string                        `json:"notes_url"`
		ParentDevice       *Device                       `json:"parent_device"`
		Platform           *nested.Platform              `json:"platform"`
		Position           *int                          `json:"position"`
		PrimaryIP          *nested.IPAddress             `json:"primary_ip"`
		PrimaryIP4         *nested.IPAddress             `json:"primary_ip4"`
		PrimaryIP6         *nested.IPAddress             `json:"primary_ip6"`
		Rack               *nested.Rack                  `json:"rack"`
		SecretsGroup       *nested.SecretsGroup          `json:"secrets_group"`
		Serial             string                        `json:"serial"`
		Site               *Site                         `json:"site"`
		Status             *shared.LabelValue            `json:"status"`
		Tags               []extras.Tag                  `json:"tags"`
		Tenant             *nested.Tenant                `json:"tenant"`
		URL                string                        `json:"url"`
		VCPosition         *int                          `json:"vc_position"`
		VCPriority         *int                          `json:"vc_priority"`
		VirtualChassis     *nested.VirtualChassis        `json:"virtual_chassis"`
	}

	// NewDevice : Structured input for a new Device record in Nautobot.
	NewDevice struct {
		Name                         string         `json:"name"`
		Role                         string         `json:"role"`
		Status                       string         `json:"status"`
		DeviceType                   string         `json:"device_type"`
		Location                     string         `json:"location,omitempty"`
		Site                         string         `json:"site,omitempty"`
		Tenant                       string         `json:"tenant,omitempty"`
		Platform                     string         `json:"platform,omitempty"`
		Serial                       string         `json:"serial,omitempty"`
		AssetTag                     string         `json:"asset_tag,omitempty"`
		Position                     int            `json:"position,omitempty"`
		Face                         string         `json:"face,omitempty"`
		VcPosition                   int            `json:"vc_position,omitempty"`
		VcPriority                   int            `json:"vc_priority,omitempty"`
		Comments                     string         `json:"comments,omitempty"`
		Rack                         string         `json:"rack,omitempty"`
		PrimaryIP4                   string         `json:"primary_ip4,omitempty"`
		PrimaryIP6                   string         `json:"primary_ip6,omitempty"`
		Cluster                      string         `json:"cluster,omitempty"`
		VirtualChassis               string         `json:"virtual_chassis,omitempty"`
		DeviceRedundancyGroup        string         `json:"device_redundancy_group,omitempty"`
		SoftwareVersion              string         `json:"software_version,omitempty"`
		SecretsGroup                 string         `json:"secrets_group,omitempty"`
		ControllerManagedDeviceGroup string         `json:"controller_managed_device_group,omitempty"`
		SoftwareImageFiles           []string       `json:"software_image_files,omitempty"`
		CustomFields                 map[string]any `json:"custom_fields,omitempty"`
		Tags                         []string       `json:"tags,omitempty"`
		ParentBay                    string         `json:"parent_bay,omitempty"`
	}
)

// DeviceGet : Go function to process requests for the endpoint: /api/dcim/devices/:id/
//
// https://demo.nautobot.com/api/docs/#/dcim/dcim_devices_retrieve
func (c *Client) DeviceGet(uuid string) (*Device, error) {
	req, err := c.Request(http.MethodGet, fmt.Sprintf("dcim/devices/%s/", url.PathEscape(uuid)), nil, nil)
	if err != nil {
		return nil, err
	}

	ret := new(Device)
	err = c.UnmarshalDo(req, ret)
	return ret, err
}

// DeviceFilter : Go function to process requests for the endpoint: /api/dcim/devices/
//
// https://demo.nautobot.com/api/docs/#/dcim/dcim_devices_list
func (c *Client) DeviceFilter(q *url.Values) ([]Device, error) {
	devices := make([]Device, 0)
	return devices, core.Paginate[Device](c.Client, "dcim/devices/", q, &devices)
}

// DeviceAll : Go function to process requests for the endpoint: /api/dcim/devices/
//
// https://demo.nautobot.com/api/docs/#/dcim/dcim_devices_list
func (c *Client) DeviceAll() ([]Device, error) {
	devices := make([]Device, 0)
	return devices, core.Paginate[Device](c.Client, "dcim/devices/", nil, &devices)
}

// DeviceCreate : Generate a new Device record in Nautobot.
func (c *Client) DeviceCreate(obj NewDevice) (Device, error) {
	var r Device
	req, err := c.Request(http.MethodPost, "dcim/devices/", obj, nil)
	if err != nil {
		return r, err
	}

	if err := c.UnmarshalDo(req, &r); err != nil {
		return r, fmt.Errorf("DeviceCreate.error.UnmarshalDo(%w)", err)
	}
	return r, nil
}

// DeviceDelete : DCIM method to delete a Device by UUIDv4 identifier.
func (c *Client) DeviceDelete(uuid string) error {
	if uuid == "" {
		return errors.New("DeviceDelete.error.UUID(UUIDv4 is missing)")
	}
	req, err := c.Request(http.MethodDelete, fmt.Sprintf("dcim/devices/%s/", url.PathEscape(uuid)), nil, nil)
	if err != nil {
		return err
	}
	return c.UnmarshalDo(req, nil)
}

// DeviceUpdate : Update an existing Device record in Nautobot.
func (c *Client) DeviceUpdate(uuid string, patch map[string]any) (Device, error) {
	var r Device
	req, err := c.Request(http.MethodPatch, fmt.Sprintf("dcim/devices/%s/", uuid), patch, nil)
	if err != nil {
		return r, err
	}
	if err := c.UnmarshalDo(req, &r); err != nil {
		return r, fmt.Errorf("DeviceUpdate.error.UnmarshalDo(%w)", err)
	}
	return r, nil
}
