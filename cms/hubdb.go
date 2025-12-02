// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

import (
	"encoding/json"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// HubdbService contains methods and other services that help with interacting with
// the hubspot API.
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
	// The timestamp indicating when the batch processing was completed.
	CompletedAt time.Time         `json:"completedAt,required" format:"date-time"`
	Results     []HubDBTableRowV3 `json:"results,required"`
	// The timestamp indicating when the batch processing began.
	StartedAt time.Time `json:"startedAt,required" format:"date-time"`
	// The current status of the batch operation, with possible values: CANCELED,
	// COMPLETE, PENDING, PROCESSING.
	//
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status BatchResponseHubDBTableRowV3Status `json:"status,required"`
	// A collection of related links associated with the batch response.
	Links map[string]string `json:"links"`
	// The timestamp indicating when the batch request was made.
	RequestedAt time.Time `json:"requestedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt respjson.Field
		Results     respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		Links       respjson.Field
		RequestedAt respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchResponseHubDBTableRowV3) RawJSON() string { return r.JSON.raw }
func (r *BatchResponseHubDBTableRowV3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current status of the batch operation, with possible values: CANCELED,
// COMPLETE, PENDING, PROCESSING.
type BatchResponseHubDBTableRowV3Status string

const (
	BatchResponseHubDBTableRowV3StatusCanceled   BatchResponseHubDBTableRowV3Status = "CANCELED"
	BatchResponseHubDBTableRowV3StatusComplete   BatchResponseHubDBTableRowV3Status = "COMPLETE"
	BatchResponseHubDBTableRowV3StatusPending    BatchResponseHubDBTableRowV3Status = "PENDING"
	BatchResponseHubDBTableRowV3StatusProcessing BatchResponseHubDBTableRowV3Status = "PROCESSING"
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
	// Column Id
	ID          string `json:"id,required"`
	Deleted     bool   `json:"deleted,required"`
	Description string `json:"description,required"`
	// Label of the column
	Label string `json:"label,required"`
	// Name of the column
	Name string `json:"name,required"`
	// Type of the column
	//
	// Any of "BOOLEAN", "CODE", "COMPOSITE", "CTA", "CURRENCY", "DATE", "DATETIME",
	// "EMBED", "FILE", "FOREIGN_ID", "HUBSPOT_VIDEO", "IMAGE", "JSON", "LOCATION",
	// "MULTISELECT", "NULL", "NUMBER", "RICHTEXT", "SELECT", "TEXT", "URL", "VIDEO".
	Type            ColumnType `json:"type,required"`
	CreatedAt       time.Time  `json:"createdAt" format:"date-time"`
	CreatedBy       SimpleUser `json:"createdBy"`
	CreatedByUserID int64      `json:"createdByUserId"`
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
		ID               respjson.Field
		Deleted          respjson.Field
		Description      respjson.Field
		Label            respjson.Field
		Name             respjson.Field
		Type             respjson.Field
		CreatedAt        respjson.Field
		CreatedBy        respjson.Field
		CreatedByUserID  respjson.Field
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
	ColumnTypeBoolean      ColumnType = "BOOLEAN"
	ColumnTypeCode         ColumnType = "CODE"
	ColumnTypeComposite    ColumnType = "COMPOSITE"
	ColumnTypeCta          ColumnType = "CTA"
	ColumnTypeCurrency     ColumnType = "CURRENCY"
	ColumnTypeDate         ColumnType = "DATE"
	ColumnTypeDatetime     ColumnType = "DATETIME"
	ColumnTypeEmbed        ColumnType = "EMBED"
	ColumnTypeFile         ColumnType = "FILE"
	ColumnTypeForeignID    ColumnType = "FOREIGN_ID"
	ColumnTypeHubspotVideo ColumnType = "HUBSPOT_VIDEO"
	ColumnTypeImage        ColumnType = "IMAGE"
	ColumnTypeJson         ColumnType = "JSON"
	ColumnTypeLocation     ColumnType = "LOCATION"
	ColumnTypeMultiselect  ColumnType = "MULTISELECT"
	ColumnTypeNull         ColumnType = "NULL"
	ColumnTypeNumber       ColumnType = "NUMBER"
	ColumnTypeRichtext     ColumnType = "RICHTEXT"
	ColumnTypeSelect       ColumnType = "SELECT"
	ColumnTypeText         ColumnType = "TEXT"
	ColumnTypeURL          ColumnType = "URL"
	ColumnTypeVideo        ColumnType = "VIDEO"
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
	// Any of "BOOLEAN", "CODE", "COMPOSITE", "CTA", "CURRENCY", "DATE", "DATETIME",
	// "EMBED", "FILE", "FOREIGN_ID", "HUBSPOT_VIDEO", "IMAGE", "JSON", "LOCATION",
	// "MULTISELECT", "NULL", "NUMBER", "RICHTEXT", "SELECT", "TEXT", "URL", "VIDEO".
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
	ColumnRequestTypeBoolean      ColumnRequestType = "BOOLEAN"
	ColumnRequestTypeCode         ColumnRequestType = "CODE"
	ColumnRequestTypeComposite    ColumnRequestType = "COMPOSITE"
	ColumnRequestTypeCta          ColumnRequestType = "CTA"
	ColumnRequestTypeCurrency     ColumnRequestType = "CURRENCY"
	ColumnRequestTypeDate         ColumnRequestType = "DATE"
	ColumnRequestTypeDatetime     ColumnRequestType = "DATETIME"
	ColumnRequestTypeEmbed        ColumnRequestType = "EMBED"
	ColumnRequestTypeFile         ColumnRequestType = "FILE"
	ColumnRequestTypeForeignID    ColumnRequestType = "FOREIGN_ID"
	ColumnRequestTypeHubspotVideo ColumnRequestType = "HUBSPOT_VIDEO"
	ColumnRequestTypeImage        ColumnRequestType = "IMAGE"
	ColumnRequestTypeJson         ColumnRequestType = "JSON"
	ColumnRequestTypeLocation     ColumnRequestType = "LOCATION"
	ColumnRequestTypeMultiselect  ColumnRequestType = "MULTISELECT"
	ColumnRequestTypeNull         ColumnRequestType = "NULL"
	ColumnRequestTypeNumber       ColumnRequestType = "NUMBER"
	ColumnRequestTypeRichtext     ColumnRequestType = "RICHTEXT"
	ColumnRequestTypeSelect       ColumnRequestType = "SELECT"
	ColumnRequestTypeText         ColumnRequestType = "TEXT"
	ColumnRequestTypeURL          ColumnRequestType = "URL"
	ColumnRequestTypeVideo        ColumnRequestType = "VIDEO"
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
	// The id of the table row
	ID string `json:"id,required"`
	// Specifies the value for the column child table id
	ChildTableID string `json:"childTableId,required"`
	// Timestamp at which the row is created
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Specifies the value for `hs_name` column, which will be used as title in the
	// dynamic pages
	Name string `json:"name,required"`
	// Specifies the value for `hs_path` column, which will be used as slug in the
	// dynamic pages
	Path        string    `json:"path,required"`
	PublishedAt time.Time `json:"publishedAt,required" format:"date-time"`
	// Timestamp at which the row is updated last time
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	// List of key value pairs with the column name and column value
	Values map[string]any `json:"values,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		ChildTableID respjson.Field
		CreatedAt    respjson.Field
		Name         respjson.Field
		Path         respjson.Field
		PublishedAt  respjson.Field
		UpdatedAt    respjson.Field
		Values       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HubDBTableRowV3) RawJSON() string { return r.JSON.raw }
func (r *HubDBTableRowV3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChildTableID, DisplayIndex, Values are required.
type HubDBTableRowV3BatchUpdateRequestParam struct {
	// Specifies the value for the column child table id
	ChildTableID int64 `json:"childTableId,required"`
	DisplayIndex int64 `json:"displayIndex,required"`
	// List of key value pairs with the column name and column value
	Values map[string]Variant `json:"values,omitzero,required"`
	// The id of the table row
	ID param.Opt[string] `json:"id,omitzero"`
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

// The properties ChildTableID, DisplayIndex, Values are required.
type HubDBTableRowV3RequestParam struct {
	// Specifies the value for the column child table id
	ChildTableID int64 `json:"childTableId,required"`
	DisplayIndex int64 `json:"displayIndex,required"`
	// List of key value pairs with the column name and column value
	Values map[string]Variant `json:"values,omitzero,required"`
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
	// Id of the table
	ID string `json:"id,required"`
	// Specifies whether child tables can be created
	AllowChildTables bool `json:"allowChildTables,required"`
	// Specifies whether the table can be read by public without authorization
	AllowPublicAPIAccess bool `json:"allowPublicApiAccess,required"`
	// Number of columns including deleted
	ColumnCount int64 `json:"columnCount,required"`
	// List of columns in the table
	Columns []Column `json:"columns,required"`
	// Timestamp at which the table is created
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	Deleted   bool      `json:"deleted,required"`
	DeletedAt time.Time `json:"deletedAt,required" format:"date-time"`
	// Specifies the key value pairs of the
	// [metadata fields](https://developers.hubspot.com/docs/cms/guides/dynamic-pages/hubdb#dynamic-pages)
	// with the associated column IDs.
	DynamicMetaTags map[string]int64 `json:"dynamicMetaTags,required"`
	// Specifies creation of multi-level dynamic pages using child tables
	EnableChildTablePages bool `json:"enableChildTablePages,required"`
	// Label of the table
	Label string `json:"label,required"`
	// Name of the table
	Name      string `json:"name,required"`
	Published bool   `json:"published,required"`
	// Timestamp at which the table is published recently
	PublishedAt time.Time `json:"publishedAt,required" format:"date-time"`
	// Number of rows in the table
	RowCount int64 `json:"rowCount,required"`
	// Timestamp at which the table is updated recently
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	// Specifies whether the table can be used for creation of dynamic pages
	UseForPages       bool       `json:"useForPages,required"`
	CreatedBy         SimpleUser `json:"createdBy"`
	IsOrderedManually bool       `json:"isOrderedManually"`
	UpdatedBy         SimpleUser `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		AllowChildTables      respjson.Field
		AllowPublicAPIAccess  respjson.Field
		ColumnCount           respjson.Field
		Columns               respjson.Field
		CreatedAt             respjson.Field
		Deleted               respjson.Field
		DeletedAt             respjson.Field
		DynamicMetaTags       respjson.Field
		EnableChildTablePages respjson.Field
		Label                 respjson.Field
		Name                  respjson.Field
		Published             respjson.Field
		PublishedAt           respjson.Field
		RowCount              respjson.Field
		UpdatedAt             respjson.Field
		UseForPages           respjson.Field
		CreatedBy             respjson.Field
		IsOrderedManually     respjson.Field
		UpdatedBy             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HubDBTableV3) RawJSON() string { return r.JSON.raw }
func (r *HubDBTableV3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AllowChildTables, AllowPublicAPIAccess, Columns, DynamicMetaTags,
// EnableChildTablePages, Label, Name, UseForPages are required.
type HubDBTableV3RequestParam struct {
	// Specifies whether child tables can be created
	AllowChildTables bool `json:"allowChildTables,required"`
	// Specifies whether the table can be read by public without authorization
	AllowPublicAPIAccess bool `json:"allowPublicApiAccess,required"`
	// List of columns in the table
	Columns []ColumnRequestParam `json:"columns,omitzero,required"`
	// Specifies the key value pairs of the
	// [metadata fields](https://developers.hubspot.com/docs/cms/guides/dynamic-pages/hubdb#dynamic-pages)
	// with the associated column IDs.
	DynamicMetaTags map[string]int64 `json:"dynamicMetaTags,omitzero,required"`
	// Specifies creation of multi-level dynamic pages using child tables
	EnableChildTablePages bool `json:"enableChildTablePages,required"`
	// Label of the table
	Label string `json:"label,required"`
	// Name of the table
	Name string `json:"name,required"`
	// Specifies whether the table can be used for creation of dynamic pages
	UseForPages bool `json:"useForPages,required"`
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
	// The total number of rows available in the collection.
	Total int64 `json:"total,required"`
	// Indicates the type of response, which is 'STREAMING' by default.
	//
	// Any of "STREAMING".
	Type   StreamingCollectionResponseWithTotalHubDBTableRowV3Type `json:"type,required"`
	Paging shared.Paging                                           `json:"paging"`
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

// Indicates the type of response, which is 'STREAMING' by default.
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
	// This field is a union of [BoundedPaging], [shared.Paging]
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
	// This field is from variant [shared.Paging].
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
