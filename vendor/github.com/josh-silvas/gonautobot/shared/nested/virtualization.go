package nested

type (
	// VirtualizationCluster : Models a subset on the VirtualizationCluster model for nested responses.
	VirtualizationCluster struct {
		ID                  string `json:"id"`
		Display             string `json:"display"`
		Name                string `json:"name"`
		URL                 string `json:"url"`
		VirtualMachineCount int    `json:"virtualmachine_count"`
	}
)
