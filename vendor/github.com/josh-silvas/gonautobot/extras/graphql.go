package extras

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	// GraphQLQuery : Request payload for a GraphQL query in Nautobot.
	GraphQLQuery struct {
		Query     string                 `json:"query"`
		Variables map[string]interface{} `json:"variables"`
	}

	gql struct {
		Data json.RawMessage `json:"data"`
	}
)

// GraphQL : Go function to process requests for the endpoint: /api/graphql/
//
// https://demo.nautobot.com/api/docs/#/graphql/graphql_create
func (c *Client) GraphQL(q *GraphQLQuery, v any) error {
	req, err := c.Request(http.MethodPost, "graphql/", q, nil)
	if err != nil {
		return err
	}
	ret := new(gql)
	if err = c.UnmarshalDo(req, ret); err != nil {
		return fmt.Errorf("GraphQL.error.UnmarshalDo(%w)", err)
	}

	if err = json.Unmarshal(ret.Data, v); err != nil {
		return fmt.Errorf("GraphQL.error.json.Unmarshal(%w)", err)
	}
	return nil
}
