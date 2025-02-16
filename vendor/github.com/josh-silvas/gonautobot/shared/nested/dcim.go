package nested

type (
	// Site : Models a subset on the Site model for nested responses.
	Site struct {
		ID      string `json:"id"`
		Display string `json:"display"`
		Name    string `json:"name"`
		Slug    string `json:"slug"`
		URL     string `json:"url"`
	}

	// Device : Models a subset on the Device model for nested responses.
	Device struct {
		ID      string `json:"id"`
		Display string `json:"display"`
		Name    string `json:"name"`
		URL     string `json:"url"`
	}

	// DeviceRole : Models a subset on the DeviceRole model for nested responses.
	DeviceRole struct {
		ID                  string `json:"id"`
		DeviceCount         int    `json:"device_count"`
		Display             string `json:"display"`
		Model               string `json:"model"`
		Slug                string `json:"slug"`
		URL                 string `json:"url"`
		VirtualMachineCount int    `json:"virtualmachine_count"`
	}

	// DeviceType : Models a subset on the DeviceType model for nested responses.
	DeviceType struct {
		ID           string        `json:"id"`
		DeviceCount  int           `json:"device_count"`
		Display      string        `json:"display"`
		Manufacturer *Manufacturer `json:"manufacturer"`
		Model        string        `json:"model"`
		Slug         string        `json:"slug"`
		URL          string        `json:"url"`
	}

	// Manufacturer : Models a subset on the Manufacturer model for nested responses.
	Manufacturer struct {
		ID              string `json:"id"`
		DeviceTypeCount int    `json:"devicetype_count"`
		Display         string `json:"display"`
		Name            string `json:"name"`
		Slug            string `json:"slug"`
		URL             string `json:"url"`
	}

	// Platform : Models a subset on the Platform model for nested responses.
	Platform struct {
		ID                  string `json:"id"`
		DeviceCount         int    `json:"device_count"`
		Display             string `json:"display"`
		Name                string `json:"name"`
		Slug                string `json:"slug"`
		URL                 string `json:"url"`
		VirtualMachineCount int    `json:"virtualmachine_count"`
	}

	// Rack : Models a subset on the Rack model for nested responses.
	Rack struct {
		ID          string `json:"id"`
		DeviceCount int    `json:"device_count"`
		Display     string `json:"display"`
		Name        string `json:"name"`
		URL         string `json:"url"`
	}

	// Location : Models a subset on the Location model for nested responses.
	Location struct {
		ID        string `json:"id"`
		Display   string `json:"display"`
		Name      string `json:"name"`
		Slug      string `json:"slug"`
		TreeDepth int    `json:"tree_depth"`
		URL       string `json:"url"`
	}

	// Cable : Models a subset on the Cable model for nested responses.
	Cable struct {
		ID      string `json:"id"`
		Display string `json:"display"`
		Label   string `json:"label"`
		URL     string `json:"url"`
	}

	// VirtualChassis : Models a subset on the VirtualChassis model for nested responses.
	VirtualChassis struct {
		ID          string `json:"id"`
		Display     string `json:"display"`
		Master      Device `json:"master"`
		MemberCount int    `json:"member_count"`
		Name        string `json:"name"`
		URL         string `json:"url"`
	}
)
