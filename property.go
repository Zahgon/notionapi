package notionapi

import (
	"time"
)

type PropertyType string

type Property interface {
	GetID() string
	GetType() PropertyType
}

type PropertyArray []Property

func (arr *PropertyArray) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

type TitleProperty struct {
	ID    PropertyID   `json:"id,omitempty"`
	Type  PropertyType `json:"type,omitempty"`
	Title []RichText   `json:"title"`
}

func (p TitleProperty) GetID() string { _ = "STUB: not implemented"; return "" }

func (p TitleProperty) GetType() PropertyType { _ = "STUB: not implemented"; return *new(PropertyType) }

type RichTextProperty struct {
	ID       PropertyID   `json:"id,omitempty"`
	Type     PropertyType `json:"type,omitempty"`
	RichText []RichText   `json:"rich_text"`
}

func (p RichTextProperty) GetID() string { _ = "STUB: not implemented"; return "" }

func (p RichTextProperty) GetType() PropertyType {
	_ = "STUB: not implemented"
	return *new(PropertyType)
}

type TextProperty struct {
	ID   PropertyID   `json:"id,omitempty"`
	Type PropertyType `json:"type,omitempty"`
	Text []RichText   `json:"text"`
}

func (p TextProperty) GetID() string { _ = "STUB: not implemented"; return "" }

func (p TextProperty) GetType() PropertyType { _ = "STUB: not implemented"; return *new(PropertyType) }

type NumberProperty struct {
	ID     PropertyID   `json:"id,omitempty"`
	Type   PropertyType `json:"type,omitempty"`
	Number float64      `json:"number"`
}

func (p NumberProperty) GetID() string { _ = "STUB: not implemented"; return "" }

func (p NumberProperty) GetType() PropertyType {
	_ = "STUB: not implemented"
	return *new(PropertyType)
}

type SelectProperty struct {
	ID     ObjectID     `json:"id,omitempty"`
	Type   PropertyType `json:"type,omitempty"`
	Select Option       `json:"select"`
}

func (p SelectProperty) GetID() string { _ = "STUB: not implemented"; return "" }

func (p SelectProperty) GetType() PropertyType {
	_ = "STUB: not implemented"
	return *new(PropertyType)
}

type MultiSelectProperty struct {
	ID          ObjectID     `json:"id,omitempty"`
	Type        PropertyType `json:"type,omitempty"`
	MultiSelect []Option     `json:"multi_select"`
}

func (p MultiSelectProperty) GetID() string { _ = "STUB: not implemented"; return "" }

func (p MultiSelectProperty) GetType() PropertyType {
	_ = "STUB: not implemented"
	return *new(PropertyType)
}

type Option struct {
	ID    PropertyID `json:"id,omitempty"`
	Name  string     `json:"name,omitempty"`
	Color Color      `json:"color,omitempty"`
}

type DateProperty struct {
	ID   ObjectID     `json:"id,omitempty"`
	Type PropertyType `json:"type,omitempty"`
	Date *DateObject  `json:"date"`
}

type DateObject struct {
	Start *Date `json:"start"`
	End   *Date `json:"end"`
}

func (p DateProperty) GetID() string { _ = "STUB: not implemented"; return "" }

func (p DateProperty) GetType() PropertyType { _ = "STUB: not implemented"; return *new(PropertyType) }

type FormulaProperty struct {
	ID      ObjectID     `json:"id,omitempty"`
	Type    PropertyType `json:"type,omitempty"`
	Formula Formula      `json:"formula"`
}

type FormulaType string

type Formula struct {
	Type    FormulaType `json:"type,omitempty"`
	String  string      `json:"string,omitempty"`
	Number  float64     `json:"number,omitempty"`
	Boolean bool        `json:"boolean,omitempty"`
	Date    *DateObject `json:"date,omitempty"`
}

func (p FormulaProperty) GetID() string { _ = "STUB: not implemented"; return "" }

func (p FormulaProperty) GetType() PropertyType {
	_ = "STUB: not implemented"
	return *new(PropertyType)
}

type RelationProperty struct {
	ID       ObjectID     `json:"id,omitempty"`
	Type     PropertyType `json:"type,omitempty"`
	Relation []Relation   `json:"relation"`
}

type Relation struct {
	ID PageID `json:"id"`
}

func (p RelationProperty) GetID() string { _ = "STUB: not implemented"; return "" }

func (p RelationProperty) GetType() PropertyType {
	_ = "STUB: not implemented"
	return *new(PropertyType)
}

type RollupProperty struct {
	ID     ObjectID     `json:"id,omitempty"`
	Type   PropertyType `json:"type,omitempty"`
	Rollup Rollup       `json:"rollup"`
}

type RollupType string

type Rollup struct {
	Type   RollupType    `json:"type,omitempty"`
	Number float64       `json:"number,omitempty"`
	Date   *DateObject   `json:"date,omitempty"`
	Array  PropertyArray `json:"array,omitempty"`
}

func (p RollupProperty) GetID() string { _ = "STUB: not implemented"; return "" }

func (p RollupProperty) GetType() PropertyType {
	_ = "STUB: not implemented"
	return *new(PropertyType)
}

type PeopleProperty struct {
	ID     ObjectID     `json:"id,omitempty"`
	Type   PropertyType `json:"type,omitempty"`
	People []User       `json:"people"`
}

func (p PeopleProperty) GetID() string { _ = "STUB: not implemented"; return "" }

func (p PeopleProperty) GetType() PropertyType {
	_ = "STUB: not implemented"
	return *new(PropertyType)
}

type FilesProperty struct {
	ID    ObjectID     `json:"id,omitempty"`
	Type  PropertyType `json:"type,omitempty"`
	Files []File       `json:"files"`
}

func (p FilesProperty) GetID() string { _ = "STUB: not implemented"; return "" }

func (p FilesProperty) GetType() PropertyType { _ = "STUB: not implemented"; return *new(PropertyType) }

type CheckboxProperty struct {
	ID       ObjectID     `json:"id,omitempty"`
	Type     PropertyType `json:"type,omitempty"`
	Checkbox bool         `json:"checkbox"`
}

func (p CheckboxProperty) GetID() string { _ = "STUB: not implemented"; return "" }

func (p CheckboxProperty) GetType() PropertyType {
	_ = "STUB: not implemented"
	return *new(PropertyType)
}

type URLProperty struct {
	ID   ObjectID     `json:"id,omitempty"`
	Type PropertyType `json:"type,omitempty"`
	URL  string       `json:"url"`
}

func (p URLProperty) GetID() string { _ = "STUB: not implemented"; return "" }

func (p URLProperty) GetType() PropertyType { _ = "STUB: not implemented"; return *new(PropertyType) }

type EmailProperty struct {
	ID    PropertyID   `json:"id,omitempty"`
	Type  PropertyType `json:"type,omitempty"`
	Email string       `json:"email"`
}

func (p EmailProperty) GetID() string { _ = "STUB: not implemented"; return "" }

func (p EmailProperty) GetType() PropertyType { _ = "STUB: not implemented"; return *new(PropertyType) }

type PhoneNumberProperty struct {
	ID          ObjectID     `json:"id,omitempty"`
	Type        PropertyType `json:"type,omitempty"`
	PhoneNumber string       `json:"phone_number"`
}

func (p PhoneNumberProperty) GetID() string { _ = "STUB: not implemented"; return "" }

func (p PhoneNumberProperty) GetType() PropertyType {
	_ = "STUB: not implemented"
	return *new(PropertyType)
}

type CreatedTimeProperty struct {
	ID          ObjectID     `json:"id,omitempty"`
	Type        PropertyType `json:"type,omitempty"`
	CreatedTime time.Time    `json:"created_time"`
}

func (p CreatedTimeProperty) GetID() string { _ = "STUB: not implemented"; return "" }

func (p CreatedTimeProperty) GetType() PropertyType {
	_ = "STUB: not implemented"
	return *new(PropertyType)
}

type CreatedByProperty struct {
	ID        ObjectID     `json:"id,omitempty"`
	Type      PropertyType `json:"type,omitempty"`
	CreatedBy User         `json:"created_by"`
}

func (p CreatedByProperty) GetID() string { _ = "STUB: not implemented"; return "" }

func (p CreatedByProperty) GetType() PropertyType {
	_ = "STUB: not implemented"
	return *new(PropertyType)
}

type LastEditedTimeProperty struct {
	ID             ObjectID     `json:"id,omitempty"`
	Type           PropertyType `json:"type,omitempty"`
	LastEditedTime time.Time    `json:"last_edited_time"`
}

func (p LastEditedTimeProperty) GetID() string { _ = "STUB: not implemented"; return "" }

func (p LastEditedTimeProperty) GetType() PropertyType {
	_ = "STUB: not implemented"
	return *new(PropertyType)
}

type LastEditedByProperty struct {
	ID           ObjectID     `json:"id,omitempty"`
	Type         PropertyType `json:"type,omitempty"`
	LastEditedBy User         `json:"last_edited_by"`
}

func (p LastEditedByProperty) GetID() string { _ = "STUB: not implemented"; return "" }

func (p LastEditedByProperty) GetType() PropertyType {
	_ = "STUB: not implemented"
	return *new(PropertyType)
}

type StatusProperty struct {
	ID     ObjectID     `json:"id,omitempty"`
	Type   PropertyType `json:"type,omitempty"`
	Status Status       `json:"status"`
}

func (p StatusProperty) GetID() string { _ = "STUB: not implemented"; return "" }

func (p StatusProperty) GetType() PropertyType {
	_ = "STUB: not implemented"
	return *new(PropertyType)
}

type UniqueIDProperty struct {
	ID       ObjectID     `json:"id,omitempty"`
	Type     PropertyType `json:"type,omitempty"`
	UniqueID UniqueID     `json:"unique_id"`
}

func (p UniqueIDProperty) GetID() string { _ = "STUB: not implemented"; return "" }

func (p UniqueIDProperty) GetType() PropertyType {
	_ = "STUB: not implemented"
	return *new(PropertyType)
}

type VerificationProperty struct {
	ID           ObjectID     `json:"id,omitempty"`
	Type         PropertyType `json:"type,omitempty"`
	Verification Verification `json:"verification"`
}

func (p VerificationProperty) GetID() string { _ = "STUB: not implemented"; return "" }

func (p VerificationProperty) GetType() PropertyType {
	_ = "STUB: not implemented"
	return *new(PropertyType)
}

type ButtonProperty struct {
	ID     ObjectID     `json:"id,omitempty"`
	Type   PropertyType `json:"type,omitempty"`
	Button Button       `json:"button"`
}

func (p ButtonProperty) GetID() string { _ = "STUB: not implemented"; return "" }

func (p ButtonProperty) GetType() PropertyType {
	_ = "STUB: not implemented"
	return *new(PropertyType)
}

type Properties map[string]Property

func (p *Properties) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func parsePageProperties(raw map[string]interface{}) (map[string]Property, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decodeProperty(raw map[string]interface{}) (Property, error) {
	_ = "STUB: not implemented"
	return *new(Property), nil
}
