package extras

import (
	"errors"
	"fmt"
	"github.com/josh-silvas/gonautobot/core"
	"github.com/josh-silvas/gonautobot/shared"
	"net/http"
	"net/url"
)

type (
	// Tag : Data type entry for a Tag in Nautobot.
	Tag struct {
		ID          string `json:"id"`
		Color       string `json:"color"`
		Created     string `json:"created"`
		Display     string `json:"display"`
		LastUpdated string `json:"last_updated"`
		Name        string `json:"name"`
		Slug        string `json:"slug"`
		URL         string `json:"url"`
	}
)

// TagGet : Fetches a tag from Nautobot by the slug name.
func (c *Client) TagGet(uuid string) (Tag, error) {
	req, err := c.Request(http.MethodGet, fmt.Sprintf("extras/tags/%s/", url.PathEscape(uuid)), nil, nil)
	if err != nil {
		return Tag{}, err
	}

	ret := new(Tag)
	if err = c.UnmarshalDo(req, ret); err != nil {
		return Tag{}, err
	}

	return *ret, nil
}

// TagFilter : Returns an array of tags from Nautobot filtered by the q (query values)
func (c *Client) TagFilter(q *url.Values) ([]Tag, error) {
	resp := make([]Tag, 0)
	if err := core.Paginate[Tag](c.Client, "extras/tags/", q, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// TagCreate : Generate a new Tag in Nautobot.
func (c *Client) TagCreate(tag Tag) (Tag, error) {
	req, err := c.Request(http.MethodPost, "extras/tags/", tag, nil)
	if err != nil {
		return tag, err
	}

	var r Tag
	if err := c.UnmarshalDo(req, &r); err != nil {
		return tag, fmt.Errorf("TagCreate.error.UnmarshalDo(%w)", err)
	}
	return r, nil
}

// TagDelete : Extras method to delete a tag by UUIDv4 identifier.
func (c *Client) TagDelete(uuid string) error {
	if uuid == "" {
		return errors.New("TagDelete.error.UUID(UUIDv4 is missing)")
	}
	req, err := c.Request(http.MethodDelete, "extras/tags/", shared.BulkDelete{{ID: uuid}}, nil)
	if err != nil {
		return err
	}
	return c.UnmarshalDo(req, nil)
}
