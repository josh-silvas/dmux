package nautobot

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	// GraphQLQuery : the required JSON request payload for
	// Nautobot's GraphQL API endpoint
	GraphQLQuery struct {
		Query     string                 `json:"query"`
		Variables map[string]interface{} `json:"variables"`
	}

	graphQLResponse struct {
		Data json.RawMessage `json:"data"`
	}
)

// GraphQL : method for API endpoint /api/graphql/
//
// This method requires a GraphQLQuery and a struct pointer with fields that mirror the query.
// See the examples/graphql.go script for an example.
//
// https://demo.nautobot.com/api/docs/#/graphql/graphql_create
func (c *Client) GraphQL(query *GraphQLQuery, v any) error {
	req, err := c.Request(
		http.MethodPost,
		"graphql/",
		query,
		nil,
	)
	if err != nil {
		return err
	}

	ret := new(graphQLResponse)
	if err = c.Do(req, ret); err != nil {
		return fmt.Errorf("error unmarshalling api response: %w", err)
	}

	if err = json.Unmarshal(ret.Data, v); err != nil {
		return fmt.Errorf("error unmarshalling into provided type: %w", err)
	}
	return nil
}
