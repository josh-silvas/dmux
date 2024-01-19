package nautobot

import (
	"encoding/json"
	"time"
)

// Location : defines a location entry in Nautobot
type Location struct {
	ID              string          `json:"id"`
	ObjectType      string          `json:"object_type"`
	Display         string          `json:"display"`
	URL             string          `json:"url"`
	NaturalSlug     string          `json:"natural_slug"`
	TimeZone        json.RawMessage `json:"time_zone"`
	Name            string          `json:"name"`
	Description     string          `json:"description"`
	Facility        string          `json:"facility"`
	Asn             int             `json:"asn"`
	PhysicalAddress string          `json:"physical_address"`
	ShippingAddress string          `json:"shipping_address"`
	Latitude        json.RawMessage `json:"latitude"`
	Longitude       json.RawMessage `json:"longitude"`
	ContactName     string          `json:"contact_name"`
	ContactPhone    string          `json:"contact_phone"`
	ContactEmail    string          `json:"contact_email"`
	Comments        string          `json:"comments"`
	Parent          struct {
		ID         string `json:"id"`
		ObjectType string `json:"object_type"`
		URL        string `json:"url"`
	} `json:"parent"`
	LocationType struct {
		ID         string `json:"id"`
		ObjectType string `json:"object_type"`
		URL        string `json:"url"`
	} `json:"location_type"`
	Status struct {
		ID         string `json:"id"`
		ObjectType string `json:"object_type"`
		URL        string `json:"url"`
	} `json:"status"`
	Tenant struct {
		ID         string `json:"id"`
		ObjectType string `json:"object_type"`
		URL        string `json:"url"`
	} `json:"tenant"`
	Created      time.Time      `json:"created"`
	LastUpdated  time.Time      `json:"last_updated"`
	NotesURL     string         `json:"notes_url"`
	CustomFields map[string]any `json:"custom_fields"`
}
