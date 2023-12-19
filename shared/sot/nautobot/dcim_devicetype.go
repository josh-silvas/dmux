package nautobot

type (
	// NestedDeviceType : defines a stub device_type entry in Nautobot
	NestedDeviceType struct {
		ID           string              `json:"id"`
		DeviceCount  int                 `json:"device_count"`
		Display      string              `json:"display"`
		Manufacturer *NestedManufacturer `json:"manufacturer"`
		Model        string              `json:"model"`
		Slug         string              `json:"slug"`
		URL          string              `json:"url"`
	}
)
