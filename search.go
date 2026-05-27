package notionapi

import (
	"context"
)

type SearchService interface {
	Do(context.Context, *SearchRequest) (*SearchResponse, error)
}

type SearchClient struct {
	apiClient *Client
}

// Searches all parent or child pages and databases that have been shared with
// an integration.
//
// Returns all pages or databases, excluding duplicated linked databases, that
// have titles that include the query param. If no query param is provided, then
// the response contains all pages or databases that have been shared with the
// integration. The results adhere to any limitations related to an integration’s
// capabilities.

// To limit the request to search only pages or to search only databases, use
// the filter param.
//
// See https://developers.notion.com/reference/post-search
func (sc *SearchClient) Do(ctx context.Context, request *SearchRequest) (*SearchResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type SearchRequest struct {
	// The text that the API compares page and database titles against.
	Query string `json:"query,omitempty"`
	// A set of criteria, direction and timestamp keys, that orders the results.
	// The only supported timestamp value is "last_edited_time". Supported
	// direction values are "ascending" and "descending". If sort is not provided,
	// then the most recently edited results are returned first.
	Sort *SortObject `json:"sort,omitempty"`
	// A set of criteria, value and property keys, that limits the results to
	// either only pages or only databases. Possible value values are "page" or
	// "database". The only supported property value is "object".
	Filter SearchFilter `json:"filter,omitempty"`
	// A cursor value returned in a previous response that If supplied, limits the
	// response to results starting after the cursor. If not supplied, then the
	// first page of results is returned. Refer to pagination for more details.
	StartCursor Cursor `json:"start_cursor,omitempty"`
	// The number of items from the full list to include in the response. Maximum: 100.
	PageSize int `json:"page_size,omitempty"`
}

type SearchResponse struct {
	Object     ObjectType `json:"object"`
	Results    []Object   `json:"results"`
	HasMore    bool       `json:"has_more"`
	NextCursor Cursor     `json:"next_cursor"`
}

func (sr *SearchResponse) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
