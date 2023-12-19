package nautobot

type (
	// NestedManufacturer : defines a stub manufacturer entry in Nautobot
	NestedManufacturer struct {
		ID              string `json:"id"`
		DeviceTypeCount int    `json:"devicetype_count"`
		Display         string `json:"display"`
		Name            string `json:"name"`
		Slug            string `json:"slug"`
		URL             string `json:"url"`
	}
)
