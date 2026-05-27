package notionapi

type PropertyConfigType string

type PropertyConfig interface {
	GetType() PropertyConfigType
	GetID() PropertyID
}

type TitlePropertyConfig struct {
	ID    PropertyID         `json:"id,omitempty"`
	Type  PropertyConfigType `json:"type"`
	Title struct{}           `json:"title"`
}

func (p TitlePropertyConfig) GetType() PropertyConfigType {
	_ = "STUB: not implemented"
	return *new(PropertyConfigType)
}

func (p TitlePropertyConfig) GetID() PropertyID { _ = "STUB: not implemented"; return *new(PropertyID) }

type RichTextPropertyConfig struct {
	ID       PropertyID         `json:"id,omitempty"`
	Type     PropertyConfigType `json:"type"`
	RichText struct{}           `json:"rich_text"`
}

func (p RichTextPropertyConfig) GetType() PropertyConfigType {
	_ = "STUB: not implemented"
	return *new(PropertyConfigType)
}

func (p RichTextPropertyConfig) GetID() PropertyID {
	_ = "STUB: not implemented"
	return *new(PropertyID)
}

type NumberPropertyConfig struct {
	ID     PropertyID         `json:"id,omitempty"`
	Type   PropertyConfigType `json:"type"`
	Number NumberFormat       `json:"number"`
}

type FormatType string

func (ft FormatType) String() string { _ = "STUB: not implemented"; return "" }

type NumberFormat struct {
	Format FormatType `json:"format"`
}

func (p NumberPropertyConfig) GetType() PropertyConfigType {
	_ = "STUB: not implemented"
	return *new(PropertyConfigType)
}

func (p NumberPropertyConfig) GetID() PropertyID {
	_ = "STUB: not implemented"
	return *new(PropertyID)
}

type SelectPropertyConfig struct {
	ID     PropertyID         `json:"id,omitempty"`
	Type   PropertyConfigType `json:"type"`
	Select Select             `json:"select"`
}

func (p SelectPropertyConfig) GetType() PropertyConfigType {
	_ = "STUB: not implemented"
	return *new(PropertyConfigType)
}

func (p SelectPropertyConfig) GetID() PropertyID {
	_ = "STUB: not implemented"
	return *new(PropertyID)
}

type MultiSelectPropertyConfig struct {
	ID          PropertyID         `json:"id,omitempty"`
	Type        PropertyConfigType `json:"type"`
	MultiSelect Select             `json:"multi_select"`
}

type Select struct {
	Options []Option `json:"options"`
}

func (p MultiSelectPropertyConfig) GetType() PropertyConfigType {
	_ = "STUB: not implemented"
	return *new(PropertyConfigType)
}

func (p MultiSelectPropertyConfig) GetID() PropertyID {
	_ = "STUB: not implemented"
	return *new(PropertyID)
}

type DatePropertyConfig struct {
	ID   PropertyID         `json:"id,omitempty"`
	Type PropertyConfigType `json:"type"`
	Date struct{}           `json:"date"`
}

func (p DatePropertyConfig) GetType() PropertyConfigType {
	_ = "STUB: not implemented"
	return *new(PropertyConfigType)
}

func (p DatePropertyConfig) GetID() PropertyID { _ = "STUB: not implemented"; return *new(PropertyID) }

type PeoplePropertyConfig struct {
	ID     PropertyID         `json:"id,omitempty"`
	Type   PropertyConfigType `json:"type"`
	People struct{}           `json:"people"`
}

func (p PeoplePropertyConfig) GetType() PropertyConfigType {
	_ = "STUB: not implemented"
	return *new(PropertyConfigType)
}

func (p PeoplePropertyConfig) GetID() PropertyID {
	_ = "STUB: not implemented"
	return *new(PropertyID)
}

type FilesPropertyConfig struct {
	ID    PropertyID         `json:"id,omitempty"`
	Type  PropertyConfigType `json:"type"`
	Files struct{}           `json:"files"`
}

func (p FilesPropertyConfig) GetType() PropertyConfigType {
	_ = "STUB: not implemented"
	return *new(PropertyConfigType)
}

func (p FilesPropertyConfig) GetID() PropertyID { _ = "STUB: not implemented"; return *new(PropertyID) }

type CheckboxPropertyConfig struct {
	ID       PropertyID         `json:"id,omitempty"`
	Type     PropertyConfigType `json:"type"`
	Checkbox struct{}           `json:"checkbox"`
}

func (p CheckboxPropertyConfig) GetType() PropertyConfigType {
	_ = "STUB: not implemented"
	return *new(PropertyConfigType)
}

func (p CheckboxPropertyConfig) GetID() PropertyID {
	_ = "STUB: not implemented"
	return *new(PropertyID)
}

type URLPropertyConfig struct {
	ID   PropertyID         `json:"id,omitempty"`
	Type PropertyConfigType `json:"type"`
	URL  struct{}           `json:"url"`
}

func (p URLPropertyConfig) GetType() PropertyConfigType {
	_ = "STUB: not implemented"
	return *new(PropertyConfigType)
}

func (p URLPropertyConfig) GetID() PropertyID { _ = "STUB: not implemented"; return *new(PropertyID) }

type EmailPropertyConfig struct {
	ID    PropertyID         `json:"id,omitempty"`
	Type  PropertyConfigType `json:"type"`
	Email struct{}           `json:"email"`
}

func (p EmailPropertyConfig) GetType() PropertyConfigType {
	_ = "STUB: not implemented"
	return *new(PropertyConfigType)
}

func (p EmailPropertyConfig) GetID() PropertyID { _ = "STUB: not implemented"; return *new(PropertyID) }

type PhoneNumberPropertyConfig struct {
	ID          PropertyID         `json:"id,omitempty"`
	Type        PropertyConfigType `json:"type"`
	PhoneNumber struct{}           `json:"phone_number"`
}

func (p PhoneNumberPropertyConfig) GetType() PropertyConfigType {
	_ = "STUB: not implemented"
	return *new(PropertyConfigType)
}

func (p PhoneNumberPropertyConfig) GetID() PropertyID {
	_ = "STUB: not implemented"
	return *new(PropertyID)
}

type FormulaPropertyConfig struct {
	ID      PropertyID         `json:"id,omitempty"`
	Type    PropertyConfigType `json:"type"`
	Formula FormulaConfig      `json:"formula"`
}

type FormulaConfig struct {
	Expression string `json:"expression"`
}

func (p FormulaPropertyConfig) GetType() PropertyConfigType {
	_ = "STUB: not implemented"
	return *new(PropertyConfigType)
}

func (p FormulaPropertyConfig) GetID() PropertyID {
	_ = "STUB: not implemented"
	return *new(PropertyID)
}

type RelationPropertyConfig struct {
	Type     PropertyConfigType `json:"type"`
	Relation RelationConfig     `json:"relation"`
}

type RelationConfigType string

func (rp RelationConfigType) String() string { _ = "STUB: not implemented"; return "" }

type SingleProperty struct{}

type DualProperty struct{}

type RelationConfig struct {
	DatabaseID         DatabaseID         `json:"database_id"`
	SyncedPropertyID   PropertyID         `json:"synced_property_id,omitempty"`
	SyncedPropertyName string             `json:"synced_property_name,omitempty"`
	Type               RelationConfigType `json:"type,omitempty"`
	SingleProperty     *SingleProperty    `json:"single_property,omitempty"`
	DualProperty       *DualProperty      `json:"dual_property,omitempty"`
}

func (p RelationPropertyConfig) GetType() PropertyConfigType {
	_ = "STUB: not implemented"
	return *new(PropertyConfigType)
}

func (p RelationPropertyConfig) GetID() PropertyID {
	_ = "STUB: not implemented"
	return *new(PropertyID)
}

type RollupPropertyConfig struct {
	ID     PropertyID         `json:"id,omitempty"`
	Type   PropertyConfigType `json:"type"`
	Rollup RollupConfig       `json:"rollup"`
}

type RollupConfig struct {
	RelationPropertyName string       `json:"relation_property_name"`
	RelationPropertyID   PropertyID   `json:"relation_property_id"`
	RollupPropertyName   string       `json:"rollup_property_name"`
	RollupPropertyID     PropertyID   `json:"rollup_property_id"`
	Function             FunctionType `json:"function"`
}

func (p RollupPropertyConfig) GetType() PropertyConfigType {
	_ = "STUB: not implemented"
	return *new(PropertyConfigType)
}

func (p RollupPropertyConfig) GetID() PropertyID {
	_ = "STUB: not implemented"
	return *new(PropertyID)
}

type CreatedTimePropertyConfig struct {
	ID          PropertyID         `json:"id,omitempty"`
	Type        PropertyConfigType `json:"type"`
	CreatedTime struct{}           `json:"created_time"`
}

func (p CreatedTimePropertyConfig) GetType() PropertyConfigType {
	_ = "STUB: not implemented"
	return *new(PropertyConfigType)
}

func (p CreatedTimePropertyConfig) GetID() PropertyID {
	_ = "STUB: not implemented"
	return *new(PropertyID)
}

type CreatedByPropertyConfig struct {
	ID        PropertyID         `json:"id"`
	Type      PropertyConfigType `json:"type"`
	CreatedBy struct{}           `json:"created_by"`
}

func (p CreatedByPropertyConfig) GetType() PropertyConfigType {
	_ = "STUB: not implemented"
	return *new(PropertyConfigType)
}

func (p CreatedByPropertyConfig) GetID() PropertyID {
	_ = "STUB: not implemented"
	return *new(PropertyID)
}

type LastEditedTimePropertyConfig struct {
	ID             PropertyID         `json:"id"`
	Type           PropertyConfigType `json:"type"`
	LastEditedTime struct{}           `json:"last_edited_time"`
}

func (p LastEditedTimePropertyConfig) GetType() PropertyConfigType {
	_ = "STUB: not implemented"
	return *new(PropertyConfigType)
}

func (p LastEditedTimePropertyConfig) GetID() PropertyID {
	_ = "STUB: not implemented"
	return *new(PropertyID)
}

type LastEditedByPropertyConfig struct {
	ID           PropertyID         `json:"id"`
	Type         PropertyConfigType `json:"type"`
	LastEditedBy struct{}           `json:"last_edited_by"`
}

func (p LastEditedByPropertyConfig) GetType() PropertyConfigType {
	_ = "STUB: not implemented"
	return *new(PropertyConfigType)
}

func (p LastEditedByPropertyConfig) GetID() PropertyID {
	_ = "STUB: not implemented"
	return *new(PropertyID)
}

type StatusPropertyConfig struct {
	ID     PropertyID         `json:"id"`
	Type   PropertyConfigType `json:"type"`
	Status StatusConfig       `json:"status"`
}

func (p StatusPropertyConfig) GetType() PropertyConfigType {
	_ = "STUB: not implemented"
	return *new(PropertyConfigType)
}

func (p StatusPropertyConfig) GetID() PropertyID {
	_ = "STUB: not implemented"
	return *new(PropertyID)
}

type StatusConfig struct {
	Options []Option      `json:"options"`
	Groups  []GroupConfig `json:"groups"`
}

type GroupConfig struct {
	ID        ObjectID   `json:"id"`
	Name      string     `json:"name"`
	Color     string     `json:"color"`
	OptionIDs []ObjectID `json:"option_ids"`
}

type UniqueIDPropertyConfig struct {
	ID       PropertyID         `json:"id,omitempty"`
	Type     PropertyConfigType `json:"type"`
	UniqueID UniqueIDConfig     `json:"unique_id"`
}

type UniqueIDConfig struct {
	Prefix string `json:"prefix"`
}

func (p UniqueIDPropertyConfig) GetType() PropertyConfigType {
	_ = "STUB: not implemented"
	return *new(PropertyConfigType)
}

func (p UniqueIDPropertyConfig) GetID() PropertyID {
	_ = "STUB: not implemented"
	return *new(PropertyID)
}

type VerificationPropertyConfig struct {
	ID           PropertyID         `json:"id,omitempty"`
	Type         PropertyConfigType `json:"type,omitempty"`
	Verification Verification       `json:"verification"`
}

func (p VerificationPropertyConfig) GetType() PropertyConfigType {
	_ = "STUB: not implemented"
	return *new(PropertyConfigType)
}

func (p VerificationPropertyConfig) GetID() PropertyID {
	_ = "STUB: not implemented"
	return *new(PropertyID)
}

type ButtonPropertyConfig struct {
	ID     PropertyID         `json:"id"`
	Type   PropertyConfigType `json:"type"`
	Button struct{}           `json:"button"`
}

func (p ButtonPropertyConfig) GetType() PropertyConfigType {
	_ = "STUB: not implemented"
	return *new(PropertyConfigType)
}

func (p ButtonPropertyConfig) GetID() PropertyID {
	_ = "STUB: not implemented"
	return *new(PropertyID)
}

type PropertyConfigs map[string]PropertyConfig

func (p *PropertyConfigs) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func parsePropertyConfigs(raw map[string]interface{}) (PropertyConfigs, error) {
	_ = "STUB: not implemented"
	return *new(PropertyConfigs), nil
}
