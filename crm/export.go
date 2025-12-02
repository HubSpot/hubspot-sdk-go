// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// ExportService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExportService] method instead.
type ExportService struct {
	Options []option.RequestOption
}

// NewExportService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewExportService(opts ...option.RequestOption) (r ExportService) {
	r = ExportService{}
	r.Options = opts
	return
}

// Begins exporting CRM data for the portal as specified in the request body
func (r *ExportService) NewAsync(ctx context.Context, body ExportNewAsyncParams, opts ...option.RequestOption) (res *shared.TaskLocator, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/exports/export/async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve detailed information about a specific CRM export, including its current
// state and properties.
func (r *ExportService) Get(ctx context.Context, exportID int64, opts ...option.RequestOption) (res *PublicExportResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("crm/v3/exports/export/%v", exportID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns the status of the export with taskId, including the URL of the resulting
// file if the export status is COMPLETE
func (r *ExportService) GetStatus(ctx context.Context, taskID int64, opts ...option.RequestOption) (res *ActionResponseWithSingleResultUri, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("crm/v3/exports/export/async/tasks/%v/status", taskID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ActionResponseWithSingleResultUri struct {
	// The timestamp when the export was completed, in ISO 8601 format.
	CompletedAt time.Time `json:"completedAt,required" format:"date-time"`
	// The timestamp when the export process started, in ISO 8601 format.
	StartedAt time.Time `json:"startedAt,required" format:"date-time"`
	// The current status of the export, which can be PENDING, PROCESSING, COMPLETE or
	// CANCELED.
	//
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status ActionResponseWithSingleResultUriStatus `json:"status,required"`
	Errors []shared.StandardError                  `json:"errors"`
	// A collection of related links associated with the export.
	Links map[string]string `json:"links"`
	// The number of errors encountered during the export process.
	NumErrors int64 `json:"numErrors"`
	// The timestamp when the export request was made, in ISO 8601 format.
	RequestedAt time.Time `json:"requestedAt" format:"date-time"`
	// The URL of the resulting file if the export status is COMPLETE.
	Result string `json:"result"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		Errors      respjson.Field
		Links       respjson.Field
		NumErrors   respjson.Field
		RequestedAt respjson.Field
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionResponseWithSingleResultUri) RawJSON() string { return r.JSON.raw }
func (r *ActionResponseWithSingleResultUri) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current status of the export, which can be PENDING, PROCESSING, COMPLETE or
// CANCELED.
type ActionResponseWithSingleResultUriStatus string

const (
	ActionResponseWithSingleResultUriStatusCanceled   ActionResponseWithSingleResultUriStatus = "CANCELED"
	ActionResponseWithSingleResultUriStatusComplete   ActionResponseWithSingleResultUriStatus = "COMPLETE"
	ActionResponseWithSingleResultUriStatusPending    ActionResponseWithSingleResultUriStatus = "PENDING"
	ActionResponseWithSingleResultUriStatusProcessing ActionResponseWithSingleResultUriStatus = "PROCESSING"
)

// The properties FilterGroups, Filters, Sorts are required.
type PublicCrmSearchRequestParam struct {
	FilterGroups []FilterGroupParam `json:"filterGroups,omitzero,required"`
	Filters      []FilterParam      `json:"filters,omitzero,required"`
	// Defines the order in which the CRM records should be returned.
	Sorts []string `json:"sorts,omitzero,required"`
	// The search query string, to filter CRM records.
	Query param.Opt[string] `json:"query,omitzero"`
	paramObj
}

func (r PublicCrmSearchRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicCrmSearchRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicCrmSearchRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AssociatedObjectType, ExportInternalValuesOptions, ExportName,
// ExportType, Format, IncludeLabeledAssociations,
// IncludePrimaryDisplayPropertyForAssociatedObjects, Language, ListID,
// ObjectProperties, ObjectType, OverrideAssociatedObjectsPerDefinitionPerRowLimit
// are required.
type PublicExportListRequestParam struct {
	AssociatedObjectType []string `json:"associatedObjectType,omitzero,required"`
	// Any of "NAMES", "VALUES".
	ExportInternalValuesOptions []string `json:"exportInternalValuesOptions,omitzero,required"`
	ExportName                  string   `json:"exportName,required"`
	// Any of "LIST".
	ExportType PublicExportListRequestExportType `json:"exportType,omitzero,required"`
	// Any of "XLS", "XLSX", "CSV".
	Format                                            PublicExportListRequestFormat `json:"format,omitzero,required"`
	IncludeLabeledAssociations                        bool                          `json:"includeLabeledAssociations,required"`
	IncludePrimaryDisplayPropertyForAssociatedObjects bool                          `json:"includePrimaryDisplayPropertyForAssociatedObjects,required"`
	// Any of "EN", "DE", "ES", "FR", "JA", "NL", "PT_BR", "IT", "PL", "SV", "FI",
	// "ZH_TW", "DA_DK", "NO".
	Language                                          PublicExportListRequestLanguage `json:"language,omitzero,required"`
	ListID                                            string                          `json:"listId,required"`
	ObjectProperties                                  []string                        `json:"objectProperties,omitzero,required"`
	ObjectType                                        string                          `json:"objectType,required"`
	OverrideAssociatedObjectsPerDefinitionPerRowLimit bool                            `json:"overrideAssociatedObjectsPerDefinitionPerRowLimit,required"`
	paramObj
}

func (r PublicExportListRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicExportListRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicExportListRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicExportListRequestExportType string

const (
	PublicExportListRequestExportTypeList PublicExportListRequestExportType = "LIST"
)

type PublicExportListRequestFormat string

const (
	PublicExportListRequestFormatXls  PublicExportListRequestFormat = "XLS"
	PublicExportListRequestFormatXlsx PublicExportListRequestFormat = "XLSX"
	PublicExportListRequestFormatCsv  PublicExportListRequestFormat = "CSV"
)

type PublicExportListRequestLanguage string

const (
	PublicExportListRequestLanguageEn   PublicExportListRequestLanguage = "EN"
	PublicExportListRequestLanguageDe   PublicExportListRequestLanguage = "DE"
	PublicExportListRequestLanguageEs   PublicExportListRequestLanguage = "ES"
	PublicExportListRequestLanguageFr   PublicExportListRequestLanguage = "FR"
	PublicExportListRequestLanguageJa   PublicExportListRequestLanguage = "JA"
	PublicExportListRequestLanguageNl   PublicExportListRequestLanguage = "NL"
	PublicExportListRequestLanguagePtBr PublicExportListRequestLanguage = "PT_BR"
	PublicExportListRequestLanguageIt   PublicExportListRequestLanguage = "IT"
	PublicExportListRequestLanguagePl   PublicExportListRequestLanguage = "PL"
	PublicExportListRequestLanguageSv   PublicExportListRequestLanguage = "SV"
	PublicExportListRequestLanguageFi   PublicExportListRequestLanguage = "FI"
	PublicExportListRequestLanguageZhTw PublicExportListRequestLanguage = "ZH_TW"
	PublicExportListRequestLanguageDaDk PublicExportListRequestLanguage = "DA_DK"
	PublicExportListRequestLanguageNo   PublicExportListRequestLanguage = "NO"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PublicExportRequestUnionParam struct {
	OfPublicExportViewRequest *PublicExportViewRequestParam `json:",omitzero,inline"`
	OfPublicExportListRequest *PublicExportListRequestParam `json:",omitzero,inline"`
	paramUnion
}

func (u PublicExportRequestUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPublicExportViewRequest, u.OfPublicExportListRequest)
}
func (u *PublicExportRequestUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PublicExportRequestUnionParam) asAny() any {
	if !param.IsOmitted(u.OfPublicExportViewRequest) {
		return u.OfPublicExportViewRequest
	} else if !param.IsOmitted(u.OfPublicExportListRequest) {
		return u.OfPublicExportListRequest
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicExportRequestUnionParam) GetPublicCrmSearchRequest() *PublicCrmSearchRequestParam {
	if vt := u.OfPublicExportViewRequest; vt != nil {
		return &vt.PublicCrmSearchRequest
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicExportRequestUnionParam) GetListID() *string {
	if vt := u.OfPublicExportListRequest; vt != nil {
		return &vt.ListID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicExportRequestUnionParam) GetExportName() *string {
	if vt := u.OfPublicExportViewRequest; vt != nil {
		return (*string)(&vt.ExportName)
	} else if vt := u.OfPublicExportListRequest; vt != nil {
		return (*string)(&vt.ExportName)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicExportRequestUnionParam) GetExportType() *string {
	if vt := u.OfPublicExportViewRequest; vt != nil {
		return (*string)(&vt.ExportType)
	} else if vt := u.OfPublicExportListRequest; vt != nil {
		return (*string)(&vt.ExportType)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicExportRequestUnionParam) GetFormat() *string {
	if vt := u.OfPublicExportViewRequest; vt != nil {
		return (*string)(&vt.Format)
	} else if vt := u.OfPublicExportListRequest; vt != nil {
		return (*string)(&vt.Format)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicExportRequestUnionParam) GetIncludeLabeledAssociations() *bool {
	if vt := u.OfPublicExportViewRequest; vt != nil {
		return (*bool)(&vt.IncludeLabeledAssociations)
	} else if vt := u.OfPublicExportListRequest; vt != nil {
		return (*bool)(&vt.IncludeLabeledAssociations)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicExportRequestUnionParam) GetIncludePrimaryDisplayPropertyForAssociatedObjects() *bool {
	if vt := u.OfPublicExportViewRequest; vt != nil {
		return (*bool)(&vt.IncludePrimaryDisplayPropertyForAssociatedObjects)
	} else if vt := u.OfPublicExportListRequest; vt != nil {
		return (*bool)(&vt.IncludePrimaryDisplayPropertyForAssociatedObjects)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicExportRequestUnionParam) GetLanguage() *string {
	if vt := u.OfPublicExportViewRequest; vt != nil {
		return (*string)(&vt.Language)
	} else if vt := u.OfPublicExportListRequest; vt != nil {
		return (*string)(&vt.Language)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicExportRequestUnionParam) GetObjectType() *string {
	if vt := u.OfPublicExportViewRequest; vt != nil {
		return (*string)(&vt.ObjectType)
	} else if vt := u.OfPublicExportListRequest; vt != nil {
		return (*string)(&vt.ObjectType)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PublicExportRequestUnionParam) GetOverrideAssociatedObjectsPerDefinitionPerRowLimit() *bool {
	if vt := u.OfPublicExportViewRequest; vt != nil {
		return (*bool)(&vt.OverrideAssociatedObjectsPerDefinitionPerRowLimit)
	} else if vt := u.OfPublicExportListRequest; vt != nil {
		return (*bool)(&vt.OverrideAssociatedObjectsPerDefinitionPerRowLimit)
	}
	return nil
}

// Returns a pointer to the underlying variant's AssociatedObjectType property, if
// present.
func (u PublicExportRequestUnionParam) GetAssociatedObjectType() []string {
	if vt := u.OfPublicExportViewRequest; vt != nil {
		return vt.AssociatedObjectType
	} else if vt := u.OfPublicExportListRequest; vt != nil {
		return vt.AssociatedObjectType
	}
	return nil
}

// Returns a pointer to the underlying variant's ExportInternalValuesOptions
// property, if present.
func (u PublicExportRequestUnionParam) GetExportInternalValuesOptions() []string {
	if vt := u.OfPublicExportViewRequest; vt != nil {
		return vt.ExportInternalValuesOptions
	} else if vt := u.OfPublicExportListRequest; vt != nil {
		return vt.ExportInternalValuesOptions
	}
	return nil
}

// Returns a pointer to the underlying variant's ObjectProperties property, if
// present.
func (u PublicExportRequestUnionParam) GetObjectProperties() []string {
	if vt := u.OfPublicExportViewRequest; vt != nil {
		return vt.ObjectProperties
	} else if vt := u.OfPublicExportListRequest; vt != nil {
		return vt.ObjectProperties
	}
	return nil
}

type PublicExportResponse struct {
	// The unique ID of the export.
	ID string `json:"id,required"`
	// The timestamp when the export was created, in ISO 8601 format.
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// The current state of the export process.
	//
	// Any of "CANCELED", "CONFLICT", "DEFERRED", "DELETED", "DONE", "ENQUEUED",
	// "FAILED", "PENDING_APPROVAL", "PROCESSING".
	ExportState PublicExportResponseExportState `json:"exportState,required"`
	// The type of export, which can be either VIEW or LIST.
	//
	// Any of "LIST", "VIEW".
	ExportType PublicExportResponseExportType `json:"exportType,required"`
	// The list of properties exported for the associated object.
	ObjectProperties []string `json:"objectProperties,required"`
	// The associated CRM object being exported.
	ObjectType string `json:"objectType,required"`
	// The timestamp when the export was last updated, in ISO 8601 format.
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	// The name assigned to the export.
	ExportName string `json:"exportName"`
	// The total number of records included in the export.
	RecordCount int64 `json:"recordCount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		ExportState      respjson.Field
		ExportType       respjson.Field
		ObjectProperties respjson.Field
		ObjectType       respjson.Field
		UpdatedAt        respjson.Field
		ExportName       respjson.Field
		RecordCount      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicExportResponse) RawJSON() string { return r.JSON.raw }
func (r *PublicExportResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current state of the export process.
type PublicExportResponseExportState string

const (
	PublicExportResponseExportStateCanceled        PublicExportResponseExportState = "CANCELED"
	PublicExportResponseExportStateConflict        PublicExportResponseExportState = "CONFLICT"
	PublicExportResponseExportStateDeferred        PublicExportResponseExportState = "DEFERRED"
	PublicExportResponseExportStateDeleted         PublicExportResponseExportState = "DELETED"
	PublicExportResponseExportStateDone            PublicExportResponseExportState = "DONE"
	PublicExportResponseExportStateEnqueued        PublicExportResponseExportState = "ENQUEUED"
	PublicExportResponseExportStateFailed          PublicExportResponseExportState = "FAILED"
	PublicExportResponseExportStatePendingApproval PublicExportResponseExportState = "PENDING_APPROVAL"
	PublicExportResponseExportStateProcessing      PublicExportResponseExportState = "PROCESSING"
)

// The type of export, which can be either VIEW or LIST.
type PublicExportResponseExportType string

const (
	PublicExportResponseExportTypeList PublicExportResponseExportType = "LIST"
	PublicExportResponseExportTypeView PublicExportResponseExportType = "VIEW"
)

// The properties AssociatedObjectType, ExportInternalValuesOptions, ExportName,
// ExportType, Format, IncludeLabeledAssociations,
// IncludePrimaryDisplayPropertyForAssociatedObjects, Language, ObjectProperties,
// ObjectType, OverrideAssociatedObjectsPerDefinitionPerRowLimit are required.
type PublicExportViewRequestParam struct {
	AssociatedObjectType []string `json:"associatedObjectType,omitzero,required"`
	// Any of "NAMES", "VALUES".
	ExportInternalValuesOptions []string `json:"exportInternalValuesOptions,omitzero,required"`
	ExportName                  string   `json:"exportName,required"`
	// Any of "VIEW".
	ExportType PublicExportViewRequestExportType `json:"exportType,omitzero,required"`
	// Any of "XLS", "XLSX", "CSV".
	Format                                            PublicExportViewRequestFormat `json:"format,omitzero,required"`
	IncludeLabeledAssociations                        bool                          `json:"includeLabeledAssociations,required"`
	IncludePrimaryDisplayPropertyForAssociatedObjects bool                          `json:"includePrimaryDisplayPropertyForAssociatedObjects,required"`
	// Any of "EN", "DE", "ES", "FR", "JA", "NL", "PT_BR", "IT", "PL", "SV", "FI",
	// "ZH_TW", "DA_DK", "NO".
	Language                                          PublicExportViewRequestLanguage `json:"language,omitzero,required"`
	ObjectProperties                                  []string                        `json:"objectProperties,omitzero,required"`
	ObjectType                                        string                          `json:"objectType,required"`
	OverrideAssociatedObjectsPerDefinitionPerRowLimit bool                            `json:"overrideAssociatedObjectsPerDefinitionPerRowLimit,required"`
	PublicCrmSearchRequest                            PublicCrmSearchRequestParam     `json:"publicCrmSearchRequest,omitzero"`
	paramObj
}

func (r PublicExportViewRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicExportViewRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicExportViewRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicExportViewRequestExportType string

const (
	PublicExportViewRequestExportTypeView PublicExportViewRequestExportType = "VIEW"
)

type PublicExportViewRequestFormat string

const (
	PublicExportViewRequestFormatXls  PublicExportViewRequestFormat = "XLS"
	PublicExportViewRequestFormatXlsx PublicExportViewRequestFormat = "XLSX"
	PublicExportViewRequestFormatCsv  PublicExportViewRequestFormat = "CSV"
)

type PublicExportViewRequestLanguage string

const (
	PublicExportViewRequestLanguageEn   PublicExportViewRequestLanguage = "EN"
	PublicExportViewRequestLanguageDe   PublicExportViewRequestLanguage = "DE"
	PublicExportViewRequestLanguageEs   PublicExportViewRequestLanguage = "ES"
	PublicExportViewRequestLanguageFr   PublicExportViewRequestLanguage = "FR"
	PublicExportViewRequestLanguageJa   PublicExportViewRequestLanguage = "JA"
	PublicExportViewRequestLanguageNl   PublicExportViewRequestLanguage = "NL"
	PublicExportViewRequestLanguagePtBr PublicExportViewRequestLanguage = "PT_BR"
	PublicExportViewRequestLanguageIt   PublicExportViewRequestLanguage = "IT"
	PublicExportViewRequestLanguagePl   PublicExportViewRequestLanguage = "PL"
	PublicExportViewRequestLanguageSv   PublicExportViewRequestLanguage = "SV"
	PublicExportViewRequestLanguageFi   PublicExportViewRequestLanguage = "FI"
	PublicExportViewRequestLanguageZhTw PublicExportViewRequestLanguage = "ZH_TW"
	PublicExportViewRequestLanguageDaDk PublicExportViewRequestLanguage = "DA_DK"
	PublicExportViewRequestLanguageNo   PublicExportViewRequestLanguage = "NO"
)

type ExportNewAsyncParams struct {
	PublicExportRequest PublicExportRequestUnionParam
	paramObj
}

func (r ExportNewAsyncParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicExportRequest)
}
func (r *ExportNewAsyncParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PublicExportRequest)
}
