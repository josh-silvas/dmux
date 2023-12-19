package nautobot

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type (
	// NestedTag : defines a stub tag entry in Nautobot
	NestedTag struct {
		ID          string `json:"id"`
		Color       string `json:"color"`
		Created     string `json:"created"`
		Display     string `json:"display"`
		LastUpdated string `json:"last_updated"`
		Name        string `json:"name"`
		Slug        string `json:"slug"`
		URL         string `json:"url"`
	}
	// NestedTagsList : defines a list of tags with a count
	NestedTagsList struct {
		Count int64        `json:"count"`
		Tags  []*NestedTag `json:"results"`
	}

	// NestedTagNotFoundError : return this error when provided tag is not found in nautobot
	NestedTagNotFoundError struct {
		tag NestedTag
	}
)

func (tnf NestedTagNotFoundError) Error() string {
	return fmt.Sprintf("Did not find tag in nautobot: %v", tnf.tag)
}

// TagsList returns a list of all nautobot tags
func (c *Client) TagsList(query *url.Values) (*NestedTagsList, error) {
	req, err := c.Request(http.MethodGet, "extras/tags/", nil, query)
	if err != nil {
		return nil, err
	}

	var result NestedTagsList

	err = c.Do(req, &result)
	if err != nil {
		return nil, fmt.Errorf("return code error: %w", err)
	}

	return &result, nil
}

// TagGet returns the tag from nautobot if it is found, including the ID
func (c *Client) TagGet(tag NestedTag) (*NestedTag, error) {
	v := url.Values{}
	v.Add("id", strings.ToLower(tag.ID))

	req, err := c.Request("GET", "extras/tags/", v, nil)

	if err != nil {
		return nil, err
	}

	var result NestedTagsList

	err = c.Do(req, &result)
	if err != nil {
		return nil, fmt.Errorf("return code error: %w", err)
	}

	if len(result.Tags) == 0 {
		return nil, NestedTagNotFoundError{tag: tag}
	}

	return result.Tags[0], nil
}

// TagExists returns true if a specific tag exists on the nautobot server
func (c *Client) TagExists(tag NestedTag) (bool, error) {
	_, err := c.TagGet(tag)

	var notFoundErr NestedTagNotFoundError
	if notFound := errors.As(err, &notFoundErr); notFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

// TagsExist returns true if multiple tags exist on the nautobot server
func (c *Client) TagsExist(tags []NestedTag) (bool, error) {
	for _, tag := range tags {
		exists, err := c.TagExists(tag)
		if err != nil {
			return false, err
		}

		if !exists {
			return false, nil
		}
	}
	return true, nil
}

// TagUpsert creates a tag if and only if it does not exist.
func (c *Client) TagUpsert(tag NestedTag) (*NestedTag, error) {
	tag.Name = strings.ToLower(tag.Name)
	t, err := c.TagGet(tag)

	var notFoundErr NestedTagNotFoundError
	if notFound := errors.As(err, &notFoundErr); notFound {
		return c.TagCreate(tag)
	}
	return t, err
}

// TagCreate creates a new tag if it doesn't exist and returns an error if it does.
// It does not check if the tag exists. That can be done with TagExists().
func (c *Client) TagCreate(tag NestedTag) (*NestedTag, error) {

	req, err := c.Request(http.MethodPost, "extras/tags/", tag, nil)

	if err != nil {
		return nil, err
	}

	var result NestedTag

	resp, err := c.RawDo(req)
	if err != nil {
		return nil, fmt.Errorf("return code error: %w", err)
	}
	if resp.StatusCode == 204 {
		return nil, fmt.Errorf("return code error: http 204: %w", err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	decErr := json.NewDecoder(resp.Body).Decode(&result)
	if errors.Is(decErr, io.EOF) {
		decErr = nil // ignore EOF errors caused by empty response body
	}
	if decErr != nil {
		return nil, fmt.Errorf("failure decoding payload: %w", decErr)
	}

	return &result, nil
}

// TagDelete deletes a tag based on it's id
func (c *Client) TagDelete(uuid string) error {
	if uuid == "" {
		return fmt.Errorf("cannot delete tag with empty ID")
	}

	type Item struct {
		ID string `json:"id"`
	}

	// delete method on extras/tags/ is a bulk operation
	items := []Item{
		{ID: uuid},
	}

	req, err := c.Request(http.MethodDelete, "extras/tags/", items, nil)
	if err != nil {
		return err
	}

	resp, err := c.RawDo(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == 204 {
		return nil
	}
	return fmt.Errorf("return code error: %w", err)
}

func slugify(name string) string {
	slug := name

	// 1. Converting to lowercase.
	slug = strings.ToLower(slug)

	// 2. Removing characters that aren’t alphanumerics, underscores, hyphens, or whitespace.
	re := regexp.MustCompile(`[^\w\s-]`)
	slug = re.ReplaceAllString(slug, "")

	// 3. Removing leading and trailing whitespace.
	slug = strings.TrimSpace(slug)

	// 4. Replacing any whitespace or repeated dashes with single dashes.
	re = regexp.MustCompile(`[-\s]+`)
	slug = re.ReplaceAllString(slug, "-")

	return slug
}

// NewTag creates a NestedTag object with the provided name
func NewTag(name string) NestedTag {
	return NestedTag{
		Name: name,
		Slug: slugify(name),
	}
}

// TagNames extracts names from a list of tags
func TagNames(tags []NestedTag) []string {
	names := make([]string, len(tags))
	for i, tag := range tags {
		names[i] = tag.Name
	}
	return names
}

// NamesToTags creates a NestedTag object for each name in the list
func NamesToTags(names []string) []NestedTag {
	tags := make([]NestedTag, len(names))
	for i, name := range names {
		tags[i] = NewTag(name)
	}
	return tags
}
