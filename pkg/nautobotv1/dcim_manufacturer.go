package nautobotv1

// Manufacturer : Type definition for a Nautobot Manufacturer.
type Manufacturer struct {
	ID              string `json:"id"`
	DeviceTypeCount int    `json:"devicetype_count"`
	Display         string `json:"display"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	URL             string `json:"url"`
}
