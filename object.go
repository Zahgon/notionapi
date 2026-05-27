package notionapi

import (
	"time"
)

type ObjectType string

func (ot ObjectType) String() string { _ = "STUB: not implemented"; return "" }

type ObjectID string

func (oID ObjectID) String() string { _ = "STUB: not implemented"; return "" }

type Object interface {
	GetObject() ObjectType
}

type Color string

func (c Color) String() string { _ = "STUB: not implemented"; return "" }

func (c Color) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

type RichTextType string

func (rtType RichTextType) String() string { _ = "STUB: not implemented"; return "" }

type MentionType string

func (mType MentionType) String() string { _ = "STUB: not implemented"; return "" }

type DatabaseMention struct {
	ID ObjectID `json:"id"`
}

type PageMention struct {
	ID ObjectID `json:"id"`
}

type TemplateMentionType string

func (tMType TemplateMentionType) String() string { _ = "STUB: not implemented"; return "" }

type TemplateMention struct {
	Type                TemplateMentionType `json:"type"`
	TemplateMentionUser string              `json:"template_mention_user,omitempty"`
	TemplateMentionDate string              `json:"template_mention_date,omitempty"`
}

type Mention struct {
	Type            MentionType      `json:"type,omitempty"`
	Database        *DatabaseMention `json:"database,omitempty"`
	Page            *PageMention     `json:"page,omitempty"`
	User            *User            `json:"user,omitempty"`
	Date            *DateObject      `json:"date,omitempty"`
	TemplateMention *TemplateMention `json:"template_mention,omitempty"`
}

type RichText struct {
	Type        RichTextType `json:"type,omitempty"`
	Text        *Text        `json:"text,omitempty"`
	Mention     *Mention     `json:"mention,omitempty"`
	Equation    *Equation    `json:"equation,omitempty"`
	Annotations *Annotations `json:"annotations,omitempty"`
	PlainText   string       `json:"plain_text,omitempty"`
	Href        string       `json:"href,omitempty"`
}

type Text struct {
	Content string `json:"content"`
	Link    *Link  `json:"link,omitempty"`
}

type Link struct {
	Url string `json:"url,omitempty"`
}

type Annotations struct {
	Bold          bool  `json:"bold"`
	Italic        bool  `json:"italic"`
	Strikethrough bool  `json:"strikethrough"`
	Underline     bool  `json:"underline"`
	Code          bool  `json:"code"`
	Color         Color `json:"color,omitempty"`
}

type RelationObject struct {
	Database           DatabaseID `json:"database"`
	SyncedPropertyName string     `json:"synced_property_name"`
}

type FunctionType string

func (ft FunctionType) String() string { _ = "STUB: not implemented"; return "" }

type Cursor string

func (c Cursor) String() string { _ = "STUB: not implemented"; return "" }

type Date time.Time

func (d *Date) String() string { _ = "STUB: not implemented"; return "" }

func (d Date) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *Date) UnmarshalText(data []byte) error { _ = "STUB: not implemented"; return nil }

// Because the API does not distinguish between datetime with a
// timezone and dates, we eventually have to try both.

// Date

// Still cannot parse it, nothing else to try.

type FileType string

type File struct {
	Name     string      `json:"name"`
	Type     FileType    `json:"type"`
	File     *FileObject `json:"file,omitempty"`
	External *FileObject `json:"external,omitempty"`
}

type FileObject struct {
	URL        string     `json:"url,omitempty"`
	ExpiryTime *time.Time `json:"expiry_time,omitempty"`
}

type Icon struct {
	Type        FileType     `json:"type"`
	Emoji       *Emoji       `json:"emoji,omitempty"`
	CustomEmoji *CustomEmoji `json:"custom_emoji,omitempty"`
	File        *FileObject  `json:"file,omitempty"`
	External    *FileObject  `json:"external,omitempty"`
}

// GetURL returns the external or internal URL depending on the image type.
func (i Icon) GetURL() string { _ = "STUB: not implemented"; return "" }

type Emoji string

type CustomEmoji struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PropertyID string

func (pID PropertyID) String() string { _ = "STUB: not implemented"; return "" }

type Status = Option

type UniqueID struct {
	Prefix *string `json:"prefix,omitempty"`
	Number int     `json:"number"`
}

func (uID UniqueID) String() string { _ = "STUB: not implemented"; return "" }

type VerificationState string

func (vs VerificationState) String() string {
	_ = "STUB: not implemented"

	// Verification documented here: https://developers.notion.com/reference/page-property-values#verification
	return ""
}

type Verification struct {
	State      VerificationState `json:"state"`
	VerifiedBy *User             `json:"verified_by,omitempty"`
	Date       *DateObject       `json:"date,omitempty"`
}

type Button struct {
}
