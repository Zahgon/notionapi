package notionapi

type ErrorCode string

type Error struct {
	Object  ObjectType `json:"object"`
	Status  int        `json:"status"`
	Code    ErrorCode  `json:"code"`
	Message string     `json:"message"`
}

func (e *Error) Error() string { _ = "STUB: not implemented"; return "" }

type RateLimitedError struct {
	Message string
}

func (e *RateLimitedError) Error() string { _ = "STUB: not implemented"; return "" }

type TokenCreateError struct {
	Code    ErrorCode `json:"error"`
	Message string    `json:"error_description"`
}

func (e *TokenCreateError) Error() string { _ = "STUB: not implemented"; return "" }
