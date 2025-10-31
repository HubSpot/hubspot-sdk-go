// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

import (
	"encoding/json"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/marketing"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// HubdbService contains methods and other services that help with interacting with
// the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHubdbService] method instead.
type HubdbService struct {
	Options []option.RequestOption
	Rows    HubdbRowService
	Tables  HubdbTableService
}

// NewHubdbService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHubdbService(opts ...option.RequestOption) (r HubdbService) {
	r = HubdbService{}
	r.Options = opts
	r.Rows = NewHubdbRowService(opts...)
	r.Tables = NewHubdbTableService(opts...)
	return
}

// The property Inputs is required.
type BatchInputHubDBTableRowBatchCloneRequestParam struct {
	Inputs []HubDBTableRowBatchCloneRequestParam `json:"inputs,omitzero,required"`
	paramObj
}

func (r BatchInputHubDBTableRowBatchCloneRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputHubDBTableRowBatchCloneRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputHubDBTableRowBatchCloneRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type BatchInputHubDBTableRowV3BatchUpdateRequestParam struct {
	Inputs []HubDBTableRowV3BatchUpdateRequestParam `json:"inputs,omitzero,required"`
	paramObj
}

func (r BatchInputHubDBTableRowV3BatchUpdateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputHubDBTableRowV3BatchUpdateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputHubDBTableRowV3BatchUpdateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type BatchInputHubDBTableRowV3RequestParam struct {
	Inputs []HubDBTableRowV3RequestParam `json:"inputs,omitzero,required"`
	paramObj
}

func (r BatchInputHubDBTableRowV3RequestParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputHubDBTableRowV3RequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputHubDBTableRowV3RequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponseHubDBTableRowV3 struct {
	CompletedAt time.Time         `json:"completedAt" format:"date-time"`
	Links       map[string]string `json:"links"`
	RequestedAt time.Time         `json:"requestedAt" format:"date-time"`
	Results     []HubDBTableRowV3 `json:"results"`
	StartedAt   time.Time         `json:"startedAt" format:"date-time"`
	// Any of "PENDING", "PROCESSING", "CANCELED", "COMPLETE".
	Status BatchResponseHubDBTableRowV3Status `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt respjson.Field
		Links       respjson.Field
		RequestedAt respjson.Field
		Results     respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchResponseHubDBTableRowV3) RawJSON() string { return r.JSON.raw }
func (r *BatchResponseHubDBTableRowV3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponseHubDBTableRowV3Status string

const (
	BatchResponseHubDBTableRowV3StatusPending    BatchResponseHubDBTableRowV3Status = "PENDING"
	BatchResponseHubDBTableRowV3StatusProcessing BatchResponseHubDBTableRowV3Status = "PROCESSING"
	BatchResponseHubDBTableRowV3StatusCanceled   BatchResponseHubDBTableRowV3Status = "CANCELED"
	BatchResponseHubDBTableRowV3StatusComplete   BatchResponseHubDBTableRowV3Status = "COMPLETE"
)

type BoundedNextPage struct {
	Offset int64  `json:"offset,required"`
	Link   string `json:"link"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Offset      respjson.Field
		Link        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BoundedNextPage) RawJSON() string { return r.JSON.raw }
func (r *BoundedNextPage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BoundedPaging struct {
	Next BoundedNextPage `json:"next"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Next        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BoundedPaging) RawJSON() string { return r.JSON.raw }
func (r *BoundedPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponseWithTotalHubDBTableV3ForwardPaging struct {
	Results []HubDBTableV3       `json:"results,required"`
	Total   int64                `json:"total,required"`
	Paging  shared.ForwardPaging `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Total       respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseWithTotalHubDBTableV3ForwardPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseWithTotalHubDBTableV3ForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Column struct {
	// Label of the column
	Label string `json:"label,required"`
	// Name of the column
	Name string `json:"name,required"`
	// Type of the column
	//
	// Any of "NULL", "TEXT", "NUMBER", "URL", "IMAGE", "SELECT", "MULTISELECT",
	// "BOOLEAN", "LOCATION", "DATE", "DATETIME", "CURRENCY", "RICHTEXT", "FOREIGN_ID",
	// "VIDEO", "CTA", "FILE", "JSON", "COMPOSITE", "CODE", "HUBSPOT_VIDEO", "EMBED".
	Type ColumnType `json:"type,required"`
	// Column Id
	ID              string     `json:"id"`
	CreatedAt       time.Time  `json:"createdAt" format:"date-time"`
	CreatedBy       SimpleUser `json:"createdBy"`
	CreatedByUserID int64      `json:"createdByUserId"`
	Deleted         bool       `json:"deleted"`
	Description     string     `json:"description"`
	// Foreign Column id
	ForeignColumnID int64 `json:"foreignColumnId"`
	// Foreign Ids
	ForeignIDs []ForeignID `json:"foreignIds"`
	// Foreign ids
	ForeignIDsByID map[string]ForeignID `json:"foreignIdsById"`
	// Foreign ids by name
	ForeignIDsByName map[string]ForeignID `json:"foreignIdsByName"`
	// Foreign table id referenced
	ForeignTableID int64 `json:"foreignTableId"`
	// Number of options available
	OptionCount int64 `json:"optionCount"`
	// Options to choose for select and multi-select columns
	Options         []shared.Option `json:"options"`
	UpdatedAt       time.Time       `json:"updatedAt" format:"date-time"`
	UpdatedBy       SimpleUser      `json:"updatedBy"`
	UpdatedByUserID int64           `json:"updatedByUserId"`
	// Column width for HubDB UI
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label            respjson.Field
		Name             respjson.Field
		Type             respjson.Field
		ID               respjson.Field
		CreatedAt        respjson.Field
		CreatedBy        respjson.Field
		CreatedByUserID  respjson.Field
		Deleted          respjson.Field
		Description      respjson.Field
		ForeignColumnID  respjson.Field
		ForeignIDs       respjson.Field
		ForeignIDsByID   respjson.Field
		ForeignIDsByName respjson.Field
		ForeignTableID   respjson.Field
		OptionCount      respjson.Field
		Options          respjson.Field
		UpdatedAt        respjson.Field
		UpdatedBy        respjson.Field
		UpdatedByUserID  respjson.Field
		Width            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Column) RawJSON() string { return r.JSON.raw }
func (r *Column) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the column
type ColumnType string

const (
	ColumnTypeNull         ColumnType = "NULL"
	ColumnTypeText         ColumnType = "TEXT"
	ColumnTypeNumber       ColumnType = "NUMBER"
	ColumnTypeURL          ColumnType = "URL"
	ColumnTypeImage        ColumnType = "IMAGE"
	ColumnTypeSelect       ColumnType = "SELECT"
	ColumnTypeMultiselect  ColumnType = "MULTISELECT"
	ColumnTypeBoolean      ColumnType = "BOOLEAN"
	ColumnTypeLocation     ColumnType = "LOCATION"
	ColumnTypeDate         ColumnType = "DATE"
	ColumnTypeDatetime     ColumnType = "DATETIME"
	ColumnTypeCurrency     ColumnType = "CURRENCY"
	ColumnTypeRichtext     ColumnType = "RICHTEXT"
	ColumnTypeForeignID    ColumnType = "FOREIGN_ID"
	ColumnTypeVideo        ColumnType = "VIDEO"
	ColumnTypeCta          ColumnType = "CTA"
	ColumnTypeFile         ColumnType = "FILE"
	ColumnTypeJson         ColumnType = "JSON"
	ColumnTypeComposite    ColumnType = "COMPOSITE"
	ColumnTypeCode         ColumnType = "CODE"
	ColumnTypeHubspotVideo ColumnType = "HUBSPOT_VIDEO"
	ColumnTypeEmbed        ColumnType = "EMBED"
)

// The properties ID, Label, Name, Options, Type are required.
type ColumnRequestParam struct {
	// Column Id
	ID int64 `json:"id,required"`
	// Label of the column
	Label string `json:"label,required"`
	// Name of the column
	Name string `json:"name,required"`
	// Options to choose for select and multi-select columns
	Options []shared.OptionParam `json:"options,omitzero,required"`
	// Type of the column
	//
	// Any of "NULL", "TEXT", "NUMBER", "URL", "IMAGE", "SELECT", "MULTISELECT",
	// "BOOLEAN", "LOCATION", "DATE", "DATETIME", "CURRENCY", "RICHTEXT", "FOREIGN_ID",
	// "VIDEO", "CTA", "FILE", "JSON", "COMPOSITE", "CODE", "HUBSPOT_VIDEO", "EMBED".
	Type ColumnRequestType `json:"type,omitzero,required"`
	// The id of the column from another table to which the column refers/points to.
	ForeignColumnID param.Opt[int64] `json:"foreignColumnId,omitzero"`
	// The id of another table to which the column refers/points to.
	ForeignTableID        param.Opt[int64] `json:"foreignTableId,omitzero"`
	MaxNumberOfCharacters param.Opt[int64] `json:"maxNumberOfCharacters,omitzero"`
	MaxNumberOfOptions    param.Opt[int64] `json:"maxNumberOfOptions,omitzero"`
	paramObj
}

func (r ColumnRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ColumnRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ColumnRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the column
type ColumnRequestType string

const (
	ColumnRequestTypeNull         ColumnRequestType = "NULL"
	ColumnRequestTypeText         ColumnRequestType = "TEXT"
	ColumnRequestTypeNumber       ColumnRequestType = "NUMBER"
	ColumnRequestTypeURL          ColumnRequestType = "URL"
	ColumnRequestTypeImage        ColumnRequestType = "IMAGE"
	ColumnRequestTypeSelect       ColumnRequestType = "SELECT"
	ColumnRequestTypeMultiselect  ColumnRequestType = "MULTISELECT"
	ColumnRequestTypeBoolean      ColumnRequestType = "BOOLEAN"
	ColumnRequestTypeLocation     ColumnRequestType = "LOCATION"
	ColumnRequestTypeDate         ColumnRequestType = "DATE"
	ColumnRequestTypeDatetime     ColumnRequestType = "DATETIME"
	ColumnRequestTypeCurrency     ColumnRequestType = "CURRENCY"
	ColumnRequestTypeRichtext     ColumnRequestType = "RICHTEXT"
	ColumnRequestTypeForeignID    ColumnRequestType = "FOREIGN_ID"
	ColumnRequestTypeVideo        ColumnRequestType = "VIDEO"
	ColumnRequestTypeCta          ColumnRequestType = "CTA"
	ColumnRequestTypeFile         ColumnRequestType = "FILE"
	ColumnRequestTypeJson         ColumnRequestType = "JSON"
	ColumnRequestTypeComposite    ColumnRequestType = "COMPOSITE"
	ColumnRequestTypeCode         ColumnRequestType = "CODE"
	ColumnRequestTypeHubspotVideo ColumnRequestType = "HUBSPOT_VIDEO"
	ColumnRequestTypeEmbed        ColumnRequestType = "EMBED"
)

type ForeignID struct {
	ID   string `json:"id,required"`
	Name string `json:"name,required"`
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ForeignID) RawJSON() string { return r.JSON.raw }
func (r *ForeignID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties CopyRows, IsHubspotDefined are required.
type HubDBTableCloneRequestParam struct {
	// Specifies whether to copy the rows during clone
	CopyRows         bool `json:"copyRows,required"`
	IsHubspotDefined bool `json:"isHubspotDefined,required"`
	// The new label for the cloned table
	NewLabel param.Opt[string] `json:"newLabel,omitzero"`
	// The new name for the cloned table
	NewName param.Opt[string] `json:"newName,omitzero"`
	paramObj
}

func (r HubDBTableCloneRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow HubDBTableCloneRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *HubDBTableCloneRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type HubDBTableRowBatchCloneRequestParam struct {
	ID   string            `json:"id,required"`
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r HubDBTableRowBatchCloneRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow HubDBTableRowBatchCloneRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *HubDBTableRowBatchCloneRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HubDBTableRowV3 struct {
	// List of key value pairs with the column name and column value
	Values map[string]any `json:"values,required"`
	// The id of the table row
	ID string `json:"id"`
	// Specifies the value for the column child table id
	ChildTableID string `json:"childTableId"`
	// Timestamp at which the row is created
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Specifies the value for `hs_name` column, which will be used as title in the
	// dynamic pages
	Name string `json:"name"`
	// Specifies the value for `hs_path` column, which will be used as slug in the
	// dynamic pages
	Path        string    `json:"path"`
	PublishedAt time.Time `json:"publishedAt" format:"date-time"`
	// Timestamp at which the row is updated last time
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Values       respjson.Field
		ID           respjson.Field
		ChildTableID respjson.Field
		CreatedAt    respjson.Field
		Name         respjson.Field
		Path         respjson.Field
		PublishedAt  respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HubDBTableRowV3) RawJSON() string { return r.JSON.raw }
func (r *HubDBTableRowV3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Values are required.
type HubDBTableRowV3BatchUpdateRequestParam struct {
	// The id of the table row
	ID string `json:"id,required"`
	// List of key value pairs with the column name and column value
	Values map[string]Variant `json:"values,omitzero,required"`
	// Specifies the value for the column child table id
	ChildTableID param.Opt[int64] `json:"childTableId,omitzero"`
	DisplayIndex param.Opt[int64] `json:"displayIndex,omitzero"`
	// Specifies the value for `hs_name` column, which will be used as title in the
	// dynamic pages
	Name param.Opt[string] `json:"name,omitzero"`
	// Specifies the value for `hs_path` column, which will be used as slug in the
	// dynamic pages
	Path param.Opt[string] `json:"path,omitzero"`
	paramObj
}

func (r HubDBTableRowV3BatchUpdateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow HubDBTableRowV3BatchUpdateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *HubDBTableRowV3BatchUpdateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Values is required.
type HubDBTableRowV3RequestParam struct {
	// List of key value pairs with the column name and column value
	Values map[string]Variant `json:"values,omitzero,required"`
	// Specifies the value for the column child table id
	ChildTableID param.Opt[int64] `json:"childTableId,omitzero"`
	DisplayIndex param.Opt[int64] `json:"displayIndex,omitzero"`
	// Specifies the value for `hs_name` column, which will be used as title in the
	// dynamic pages
	Name param.Opt[string] `json:"name,omitzero"`
	// Specifies the value for `hs_path` column, which will be used as slug in the
	// dynamic pages
	Path param.Opt[string] `json:"path,omitzero"`
	paramObj
}

func (r HubDBTableRowV3RequestParam) MarshalJSON() (data []byte, err error) {
	type shadow HubDBTableRowV3RequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *HubDBTableRowV3RequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HubDBTableV3 struct {
	DeletedAt time.Time `json:"deletedAt,required" format:"date-time"`
	// Label of the table
	Label string `json:"label,required"`
	// Name of the table
	Name string `json:"name,required"`
	// Id of the table
	ID string `json:"id"`
	// Specifies whether child tables can be created
	AllowChildTables bool `json:"allowChildTables"`
	// Specifies whether the table can be read by public without authorization
	AllowPublicAPIAccess bool `json:"allowPublicApiAccess"`
	// Number of columns including deleted
	ColumnCount int64 `json:"columnCount"`
	// List of columns in the table
	Columns []Column `json:"columns"`
	// Timestamp at which the table is created
	CreatedAt time.Time  `json:"createdAt" format:"date-time"`
	CreatedBy SimpleUser `json:"createdBy"`
	Deleted   bool       `json:"deleted"`
	// Specifies the key value pairs of the
	// [metadata fields](https://developers.hubspot.com/docs/cms/guides/dynamic-pages/hubdb#dynamic-pages)
	// with the associated column IDs.
	DynamicMetaTags map[string]int64 `json:"dynamicMetaTags"`
	// Specifies creation of multi-level dynamic pages using child tables
	EnableChildTablePages bool `json:"enableChildTablePages"`
	IsOrderedManually     bool `json:"isOrderedManually"`
	Published             bool `json:"published"`
	// Timestamp at which the table is published recently
	PublishedAt time.Time `json:"publishedAt" format:"date-time"`
	// Number of rows in the table
	RowCount int64 `json:"rowCount"`
	// Timestamp at which the table is updated recently
	UpdatedAt time.Time  `json:"updatedAt" format:"date-time"`
	UpdatedBy SimpleUser `json:"updatedBy"`
	// Specifies whether the table can be used for creation of dynamic pages
	UseForPages bool `json:"useForPages"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeletedAt             respjson.Field
		Label                 respjson.Field
		Name                  respjson.Field
		ID                    respjson.Field
		AllowChildTables      respjson.Field
		AllowPublicAPIAccess  respjson.Field
		ColumnCount           respjson.Field
		Columns               respjson.Field
		CreatedAt             respjson.Field
		CreatedBy             respjson.Field
		Deleted               respjson.Field
		DynamicMetaTags       respjson.Field
		EnableChildTablePages respjson.Field
		IsOrderedManually     respjson.Field
		Published             respjson.Field
		PublishedAt           respjson.Field
		RowCount              respjson.Field
		UpdatedAt             respjson.Field
		UpdatedBy             respjson.Field
		UseForPages           respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HubDBTableV3) RawJSON() string { return r.JSON.raw }
func (r *HubDBTableV3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Label, Name are required.
type HubDBTableV3RequestParam struct {
	// Label of the table
	Label string `json:"label,required"`
	// Name of the table
	Name string `json:"name,required"`
	// Specifies whether child tables can be created
	AllowChildTables param.Opt[bool] `json:"allowChildTables,omitzero"`
	// Specifies whether the table can be read by public without authorization
	AllowPublicAPIAccess param.Opt[bool] `json:"allowPublicApiAccess,omitzero"`
	// Specifies creation of multi-level dynamic pages using child tables
	EnableChildTablePages param.Opt[bool] `json:"enableChildTablePages,omitzero"`
	// Specifies whether the table can be used for creation of dynamic pages
	UseForPages param.Opt[bool] `json:"useForPages,omitzero"`
	// List of columns in the table
	Columns []ColumnRequestParam `json:"columns,omitzero"`
	// Specifies the key value pairs of the
	// [metadata fields](https://developers.hubspot.com/docs/cms/guides/dynamic-pages/hubdb#dynamic-pages)
	// with the associated column IDs.
	DynamicMetaTags map[string]int64 `json:"dynamicMetaTags,omitzero"`
	paramObj
}

func (r HubDBTableV3RequestParam) MarshalJSON() (data []byte, err error) {
	type shadow HubDBTableV3RequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *HubDBTableV3RequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ImportResult struct {
	// Specifies number of duplicate rows
	DuplicateRows int64 `json:"duplicateRows,required"`
	// List of errors during import
	Errors []shared.Error `json:"errors,required"`
	// Specifies whether row limit exceeded during import
	RowLimitExceeded bool `json:"rowLimitExceeded,required"`
	// Specifies number of rows imported
	RowsImported int64 `json:"rowsImported,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DuplicateRows    respjson.Field
		Errors           respjson.Field
		RowLimitExceeded respjson.Field
		RowsImported     respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ImportResult) RawJSON() string { return r.JSON.raw }
func (r *ImportResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RandomAccessCollectionResponseWithTotalHubDBTableRowV3 struct {
	Results []shared.HubDBTableRowV3Wrapper `json:"results,required"`
	Total   int64                           `json:"total,required"`
	// Any of "RANDOM_ACCESS".
	Type   RandomAccessCollectionResponseWithTotalHubDBTableRowV3Type `json:"type,required"`
	Paging BoundedPaging                                              `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Total       respjson.Field
		Type        respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RandomAccessCollectionResponseWithTotalHubDBTableRowV3) RawJSON() string { return r.JSON.raw }
func (r *RandomAccessCollectionResponseWithTotalHubDBTableRowV3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RandomAccessCollectionResponseWithTotalHubDBTableRowV3Type string

const (
	RandomAccessCollectionResponseWithTotalHubDBTableRowV3TypeRandomAccess RandomAccessCollectionResponseWithTotalHubDBTableRowV3Type = "RANDOM_ACCESS"
)

type SimpleUser struct {
	ID        string `json:"id,required"`
	Email     string `json:"email,required"`
	FirstName string `json:"firstName,required"`
	LastName  string `json:"lastName,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Email       respjson.Field
		FirstName   respjson.Field
		LastName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimpleUser) RawJSON() string { return r.JSON.raw }
func (r *SimpleUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StreamingCollectionResponseWithTotalHubDBTableRowV3 struct {
	Results []shared.HubDBTableRowV3Wrapper `json:"results,required"`
	Total   int64                           `json:"total,required"`
	// Any of "STREAMING".
	Type StreamingCollectionResponseWithTotalHubDBTableRowV3Type `json:"type,required"`
	// Contains information pagination of results.
	Paging marketing.Paging `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Total       respjson.Field
		Type        respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StreamingCollectionResponseWithTotalHubDBTableRowV3) RawJSON() string { return r.JSON.raw }
func (r *StreamingCollectionResponseWithTotalHubDBTableRowV3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StreamingCollectionResponseWithTotalHubDBTableRowV3Type string

const (
	StreamingCollectionResponseWithTotalHubDBTableRowV3TypeStreaming StreamingCollectionResponseWithTotalHubDBTableRowV3Type = "STREAMING"
)

// UnifiedCollectionResponseWithTotalBaseHubDBTableRowV3Union contains all possible
// properties and values from
// [RandomAccessCollectionResponseWithTotalHubDBTableRowV3],
// [StreamingCollectionResponseWithTotalHubDBTableRowV3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type UnifiedCollectionResponseWithTotalBaseHubDBTableRowV3Union struct {
	Results []shared.HubDBTableRowV3Wrapper `json:"results"`
	Total   int64                           `json:"total"`
	Type    string                          `json:"type"`
	// This field is a union of [BoundedPaging], [marketing.Paging]
	Paging UnifiedCollectionResponseWithTotalBaseHubDBTableRowV3UnionPaging `json:"paging"`
	JSON   struct {
		Results respjson.Field
		Total   respjson.Field
		Type    respjson.Field
		Paging  respjson.Field
		raw     string
	} `json:"-"`
}

func (u UnifiedCollectionResponseWithTotalBaseHubDBTableRowV3Union) AsRandomAccess() (v RandomAccessCollectionResponseWithTotalHubDBTableRowV3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnifiedCollectionResponseWithTotalBaseHubDBTableRowV3Union) AsStreaming() (v StreamingCollectionResponseWithTotalHubDBTableRowV3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u UnifiedCollectionResponseWithTotalBaseHubDBTableRowV3Union) RawJSON() string {
	return u.JSON.raw
}

func (r *UnifiedCollectionResponseWithTotalBaseHubDBTableRowV3Union) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnifiedCollectionResponseWithTotalBaseHubDBTableRowV3UnionPaging is an implicit
// subunion of [UnifiedCollectionResponseWithTotalBaseHubDBTableRowV3Union].
// UnifiedCollectionResponseWithTotalBaseHubDBTableRowV3UnionPaging provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnifiedCollectionResponseWithTotalBaseHubDBTableRowV3Union].
type UnifiedCollectionResponseWithTotalBaseHubDBTableRowV3UnionPaging struct {
	// This field is a union of [BoundedNextPage], [shared.NextPage]
	Next UnifiedCollectionResponseWithTotalBaseHubDBTableRowV3UnionPagingNext `json:"next"`
	// This field is from variant [marketing.Paging].
	Prev shared.PreviousPage `json:"prev"`
	JSON struct {
		Next respjson.Field
		Prev respjson.Field
		raw  string
	} `json:"-"`
}

func (r *UnifiedCollectionResponseWithTotalBaseHubDBTableRowV3UnionPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnifiedCollectionResponseWithTotalBaseHubDBTableRowV3UnionPagingNext is an
// implicit subunion of
// [UnifiedCollectionResponseWithTotalBaseHubDBTableRowV3Union].
// UnifiedCollectionResponseWithTotalBaseHubDBTableRowV3UnionPagingNext provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnifiedCollectionResponseWithTotalBaseHubDBTableRowV3Union].
type UnifiedCollectionResponseWithTotalBaseHubDBTableRowV3UnionPagingNext struct {
	// This field is from variant [BoundedNextPage].
	Offset int64  `json:"offset"`
	Link   string `json:"link"`
	// This field is from variant [shared.NextPage].
	After string `json:"after"`
	JSON  struct {
		Offset respjson.Field
		Link   respjson.Field
		After  respjson.Field
		raw    string
	} `json:"-"`
}

func (r *UnifiedCollectionResponseWithTotalBaseHubDBTableRowV3UnionPagingNext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
