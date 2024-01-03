package nautobot_v1

// Platform : Type definition for the Platform struct
type Platform struct {
	ID                  string `json:"id"`
	DeviceCount         int    `json:"device_count"`
	Display             string `json:"display"`
	Name                string `json:"name"`
	Slug                string `json:"slug"`
	URL                 string `json:"url"`
	VirtualMachineCount int    `json:"virtualmachine_count"`
}
