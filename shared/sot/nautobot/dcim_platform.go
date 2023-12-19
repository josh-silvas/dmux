package nautobot

type (
	// NestedPlatform : defines a stub platform entry in Nautobot
	NestedPlatform struct {
		ID                  string `json:"id"`
		DeviceCount         int    `json:"device_count"`
		Display             string `json:"display"`
		Name                string `json:"name"`
		Slug                string `json:"slug"`
		URL                 string `json:"url"`
		VirtualMachineCount int    `json:"virtualmachine_count"`
	}
)
