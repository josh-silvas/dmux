package nautobot

type (
	// NestedDeviceRole : defines a stub device_role entry in Nautobot
	NestedDeviceRole struct {
		ID                  string `json:"id"`
		DeviceCount         int    `json:"device_count"`
		Display             string `json:"display"`
		Model               string `json:"model"`
		Slug                string `json:"slug"`
		URL                 string `json:"url"`
		VirtualMachineCount int    `json:"virtualmachine_count"`
	}
)
