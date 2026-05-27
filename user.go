package notionapi

import (
	"context"
)

type UserID string

func (uID UserID) String() string { _ = "STUB: not implemented"; return "" }

type UserService interface {
	List(context.Context, *Pagination) (*UsersListResponse, error)
	Get(context.Context, UserID) (*User, error)
	Me(context.Context) (*User, error)
}

type UserClient struct {
	apiClient *Client
}

// Returns a paginated list of Users for the workspace. The response may contain
// fewer than page_size of results.
//
// See https://developers.notion.com/reference/get-users
func (uc *UserClient) List(ctx context.Context, pagination *Pagination) (*UsersListResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves a User using the ID specified.
//
// See https://developers.notion.com/reference/get-user
func (uc *UserClient) Get(ctx context.Context, id UserID) (*User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Retrieves the bot User associated with the API token provided in the
// authorization header. The bot will have an owner field with information about
// the person who authorized the integration.
//
// See https://developers.notion.com/reference/get-self
func (uc *UserClient) Me(ctx context.Context) (*User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type UserType string

type User struct {
	Object    ObjectType `json:"object,omitempty"`
	ID        UserID     `json:"id"`
	Type      UserType   `json:"type,omitempty"`
	Name      string     `json:"name,omitempty"`
	AvatarURL string     `json:"avatar_url,omitempty"`
	Person    *Person    `json:"person,omitempty"`
	Bot       *Bot       `json:"bot,omitempty"`
}

type Person struct {
	Email string `json:"email"`
}

type Bot struct {
	Owner         Owner  `json:"owner"`
	WorkspaceName string `json:"workspace_name"`
}

type Owner struct {
	Type      string `json:"type"`
	Workspace bool   `json:"workspace"`
}

type UsersListResponse struct {
	Object     ObjectType `json:"object"`
	Results    []User     `json:"results"`
	HasMore    bool       `json:"has_more"`
	NextCursor Cursor     `json:"next_cursor"`
}
