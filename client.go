package notionapi

import (
	"context"
	"net/http"
	"net/url"
)

const (
	apiURL        = "https://api.notion.com"
	apiVersion    = "v1"
	notionVersion = "2022-06-28"
	maxRetries    = 3
)

type Token string

type errJsonDecodeFunc func(data []byte) error

func (it Token) String() string {
	_ = "STUB: not implemented"

	// ClientOption to configure API client
	return ""
}

type ClientOption func(*Client)

type Client struct {
	httpClient    *http.Client
	baseUrl       *url.URL
	apiVersion    string
	notionVersion string

	maxRetries int

	Token Token

	// used in Authorization header only for requests that require Basic authentication.
	oauthID     string
	oauthSecret string

	Database       DatabaseService
	Block          BlockService
	Page           PageService
	User           UserService
	Search         SearchService
	Comment        CommentService
	Authentication AuthenticationService
}

func NewClient(token Token, opts ...ClientOption) *Client { _ = "STUB: not implemented"; return nil }

// WithHTTPClient overrides the default http.Client
func WithHTTPClient(client *http.Client) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

// WithVersion overrides the Notion API version
func WithVersion(version string) ClientOption { _ = "STUB: not implemented"; return *new(ClientOption) }

// WithRetry overrides the default number of max retry attempts on 429 errors
func WithRetry(retries int) ClientOption { _ = "STUB: not implemented"; return *new(ClientOption) }

// WithOAuthAppCredentials sets the OAuth app ID and secret to use when fetching a token from Notion.
func WithOAuthAppCredentials(id, secret string) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func (c *Client) request(ctx context.Context, method string, urlStr string, queryParams map[string]string, requestBody interface{}) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) requestImpl(ctx context.Context, method string, urlStr string, queryParams map[string]string, requestBody interface{}, basicAuth bool, errDecoder errJsonDecodeFunc) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// https://developers.notion.com/reference/request-limits#rate-limits

// should not happen

func decodeClientError(data []byte) error { _ = "STUB: not implemented"; return nil }

type Pagination struct {
	StartCursor Cursor
	PageSize    int
}

func (p *Pagination) ToQuery() map[string]string { _ = "STUB: not implemented"; return nil }
