// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cms

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/crm"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// MediaBridgePropertyService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMediaBridgePropertyService] method instead.
type MediaBridgePropertyService struct {
	Options []option.RequestOption
}

// NewMediaBridgePropertyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMediaBridgePropertyService(opts ...option.RequestOption) (r MediaBridgePropertyService) {
	r = MediaBridgePropertyService{}
	r.Options = opts
	return
}

// Create a new property for the specified media type
func (r *MediaBridgePropertyService) New(ctx context.Context, objectType string, params MediaBridgePropertyNewParams, opts ...option.RequestOption) (res *shared.Property, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AppID == "" {
		err = errors.New("missing required appId parameter")
		return
	}
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%s/properties/%s", params.AppID, objectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Update an existing property for an object type.
func (r *MediaBridgePropertyService) Update(ctx context.Context, propertyName string, params MediaBridgePropertyUpdateParams, opts ...option.RequestOption) (res *shared.Property, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AppID == "" {
		err = errors.New("missing required appId parameter")
		return
	}
	if params.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	if propertyName == "" {
		err = errors.New("missing required propertyName parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%s/properties/%s/%s", params.AppID, params.ObjectType, propertyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Get the existing properties defined for a media object type.
func (r *MediaBridgePropertyService) List(ctx context.Context, objectType string, query MediaBridgePropertyListParams, opts ...option.RequestOption) (res *MediaBridgePropertyListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AppID == "" {
		err = errors.New("missing required appId parameter")
		return
	}
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%s/properties/%s", query.AppID, objectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete an existing property for an object type.
func (r *MediaBridgePropertyService) Delete(ctx context.Context, propertyName string, body MediaBridgePropertyDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.AppID == "" {
		err = errors.New("missing required appId parameter")
		return
	}
	if body.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	if propertyName == "" {
		err = errors.New("missing required propertyName parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%s/properties/%s/%s", body.AppID, body.ObjectType, propertyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Archive a batch of existing properties for the specified types.
func (r *MediaBridgePropertyService) ArchiveBatch(ctx context.Context, objectType string, params MediaBridgePropertyArchiveBatchParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if params.AppID == "" {
		err = errors.New("missing required appId parameter")
		return
	}
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%s/properties/%s/batch/archive", params.AppID, objectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Create a batch of properties of the specified object type.
func (r *MediaBridgePropertyService) NewBatch(ctx context.Context, objectType string, params MediaBridgePropertyNewBatchParams, opts ...option.RequestOption) (res *shared.BatchResponseProperty, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AppID == "" {
		err = errors.New("missing required appId parameter")
		return
	}
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%s/properties/%s/batch/create", params.AppID, objectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get the details for an existing property by name.
func (r *MediaBridgePropertyService) Get(ctx context.Context, propertyName string, query MediaBridgePropertyGetParams, opts ...option.RequestOption) (res *shared.Property, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AppID == "" {
		err = errors.New("missing required appId parameter")
		return
	}
	if query.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	if propertyName == "" {
		err = errors.New("missing required propertyName parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%s/properties/%s/%s", query.AppID, query.ObjectType, propertyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get the details for a batch of properties for a specified object type.
func (r *MediaBridgePropertyService) GetBatch(ctx context.Context, objectType string, params MediaBridgePropertyGetBatchParams, opts ...option.RequestOption) (res *shared.BatchResponseProperty, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AppID == "" {
		err = errors.New("missing required appId parameter")
		return
	}
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	path := fmt.Sprintf("media-bridge/v1/%s/properties/%s/batch/read", params.AppID, objectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type MediaBridgePropertyListResponse struct {
	Results []MediaBridgePropertyListResponseResult `json:"results,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaBridgePropertyListResponse) RawJSON() string { return r.JSON.raw }
func (r *MediaBridgePropertyListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgePropertyListResponseResult struct {
	Description        string                                        `json:"description,required"`
	FieldType          string                                        `json:"fieldType,required"`
	GroupName          string                                        `json:"groupName,required"`
	Label              string                                        `json:"label,required"`
	Name               string                                        `json:"name,required"`
	Options            []MediaBridgePropertyListResponseResultOption `json:"options,required"`
	Type               string                                        `json:"type,required"`
	Archived           bool                                          `json:"archived"`
	ArchivedAt         time.Time                                     `json:"archivedAt" format:"date-time"`
	Calculated         bool                                          `json:"calculated"`
	CalculationFormula string                                        `json:"calculationFormula"`
	CreatedAt          time.Time                                     `json:"createdAt" format:"date-time"`
	CreatedUserID      string                                        `json:"createdUserId"`
	// Any of "non_sensitive", "sensitive", "highly_sensitive".
	DataSensitivity string `json:"dataSensitivity"`
	// Any of "absolute", "absolute_with_relative", "time_since", "time_until".
	DateDisplayHint         string                              `json:"dateDisplayHint"`
	DisplayOrder            int64                               `json:"displayOrder"`
	ExternalOptions         bool                                `json:"externalOptions"`
	FormField               bool                                `json:"formField"`
	HasUniqueValue          bool                                `json:"hasUniqueValue"`
	Hidden                  bool                                `json:"hidden"`
	HubspotDefined          bool                                `json:"hubspotDefined"`
	ModificationMetadata    shared.PropertyModificationMetadata `json:"modificationMetadata"`
	ReferencedObjectType    string                              `json:"referencedObjectType"`
	SensitiveDataCategories []string                            `json:"sensitiveDataCategories"`
	ShowCurrencySymbol      bool                                `json:"showCurrencySymbol"`
	UpdatedAt               time.Time                           `json:"updatedAt" format:"date-time"`
	UpdatedUserID           string                              `json:"updatedUserId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description             respjson.Field
		FieldType               respjson.Field
		GroupName               respjson.Field
		Label                   respjson.Field
		Name                    respjson.Field
		Options                 respjson.Field
		Type                    respjson.Field
		Archived                respjson.Field
		ArchivedAt              respjson.Field
		Calculated              respjson.Field
		CalculationFormula      respjson.Field
		CreatedAt               respjson.Field
		CreatedUserID           respjson.Field
		DataSensitivity         respjson.Field
		DateDisplayHint         respjson.Field
		DisplayOrder            respjson.Field
		ExternalOptions         respjson.Field
		FormField               respjson.Field
		HasUniqueValue          respjson.Field
		Hidden                  respjson.Field
		HubspotDefined          respjson.Field
		ModificationMetadata    respjson.Field
		ReferencedObjectType    respjson.Field
		SensitiveDataCategories respjson.Field
		ShowCurrencySymbol      respjson.Field
		UpdatedAt               respjson.Field
		UpdatedUserID           respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaBridgePropertyListResponseResult) RawJSON() string { return r.JSON.raw }
func (r *MediaBridgePropertyListResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgePropertyListResponseResultOption struct {
	Hidden       bool   `json:"hidden,required"`
	Label        string `json:"label,required"`
	Value        string `json:"value,required"`
	Description  string `json:"description"`
	DisplayOrder int64  `json:"displayOrder"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Hidden       respjson.Field
		Label        respjson.Field
		Value        respjson.Field
		Description  respjson.Field
		DisplayOrder respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaBridgePropertyListResponseResultOption) RawJSON() string { return r.JSON.raw }
func (r *MediaBridgePropertyListResponseResultOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgePropertyNewParams struct {
	AppID          string `path:"appId,required" json:"-"`
	PropertyCreate shared.PropertyCreateParam
	paramObj
}

func (r MediaBridgePropertyNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PropertyCreate)
}
func (r *MediaBridgePropertyNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PropertyCreate)
}

type MediaBridgePropertyUpdateParams struct {
	AppID              string            `path:"appId,required" json:"-"`
	ObjectType         string            `path:"objectType,required" json:"-"`
	CalculationFormula param.Opt[string] `json:"calculationFormula,omitzero"`
	Description        param.Opt[string] `json:"description,omitzero"`
	DisplayOrder       param.Opt[int64]  `json:"displayOrder,omitzero"`
	FormField          param.Opt[bool]   `json:"formField,omitzero"`
	GroupName          param.Opt[string] `json:"groupName,omitzero"`
	HasUniqueValue     param.Opt[bool]   `json:"hasUniqueValue,omitzero"`
	Hidden             param.Opt[bool]   `json:"hidden,omitzero"`
	Label              param.Opt[string] `json:"label,omitzero"`
	// Any of "booleancheckbox", "calculation_equation", "checkbox", "date", "file",
	// "html", "number", "phonenumber", "radio", "select", "text", "textarea".
	FieldType MediaBridgePropertyUpdateParamsFieldType `json:"fieldType,omitzero"`
	Options   []shared.OptionInputParam                `json:"options,omitzero"`
	// Any of "bool", "date", "datetime", "enumeration", "number", "phone_number",
	// "string".
	Type MediaBridgePropertyUpdateParamsType `json:"type,omitzero"`
	paramObj
}

func (r MediaBridgePropertyUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow MediaBridgePropertyUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MediaBridgePropertyUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaBridgePropertyUpdateParamsFieldType string

const (
	MediaBridgePropertyUpdateParamsFieldTypeBooleancheckbox     MediaBridgePropertyUpdateParamsFieldType = "booleancheckbox"
	MediaBridgePropertyUpdateParamsFieldTypeCalculationEquation MediaBridgePropertyUpdateParamsFieldType = "calculation_equation"
	MediaBridgePropertyUpdateParamsFieldTypeCheckbox            MediaBridgePropertyUpdateParamsFieldType = "checkbox"
	MediaBridgePropertyUpdateParamsFieldTypeDate                MediaBridgePropertyUpdateParamsFieldType = "date"
	MediaBridgePropertyUpdateParamsFieldTypeFile                MediaBridgePropertyUpdateParamsFieldType = "file"
	MediaBridgePropertyUpdateParamsFieldTypeHTML                MediaBridgePropertyUpdateParamsFieldType = "html"
	MediaBridgePropertyUpdateParamsFieldTypeNumber              MediaBridgePropertyUpdateParamsFieldType = "number"
	MediaBridgePropertyUpdateParamsFieldTypePhonenumber         MediaBridgePropertyUpdateParamsFieldType = "phonenumber"
	MediaBridgePropertyUpdateParamsFieldTypeRadio               MediaBridgePropertyUpdateParamsFieldType = "radio"
	MediaBridgePropertyUpdateParamsFieldTypeSelect              MediaBridgePropertyUpdateParamsFieldType = "select"
	MediaBridgePropertyUpdateParamsFieldTypeText                MediaBridgePropertyUpdateParamsFieldType = "text"
	MediaBridgePropertyUpdateParamsFieldTypeTextarea            MediaBridgePropertyUpdateParamsFieldType = "textarea"
)

type MediaBridgePropertyUpdateParamsType string

const (
	MediaBridgePropertyUpdateParamsTypeBool        MediaBridgePropertyUpdateParamsType = "bool"
	MediaBridgePropertyUpdateParamsTypeDate        MediaBridgePropertyUpdateParamsType = "date"
	MediaBridgePropertyUpdateParamsTypeDatetime    MediaBridgePropertyUpdateParamsType = "datetime"
	MediaBridgePropertyUpdateParamsTypeEnumeration MediaBridgePropertyUpdateParamsType = "enumeration"
	MediaBridgePropertyUpdateParamsTypeNumber      MediaBridgePropertyUpdateParamsType = "number"
	MediaBridgePropertyUpdateParamsTypePhoneNumber MediaBridgePropertyUpdateParamsType = "phone_number"
	MediaBridgePropertyUpdateParamsTypeString      MediaBridgePropertyUpdateParamsType = "string"
)

type MediaBridgePropertyListParams struct {
	AppID string `path:"appId,required" json:"-"`
	paramObj
}

type MediaBridgePropertyDeleteParams struct {
	AppID      string `path:"appId,required" json:"-"`
	ObjectType string `path:"objectType,required" json:"-"`
	paramObj
}

type MediaBridgePropertyArchiveBatchParams struct {
	AppID                  string `path:"appId,required" json:"-"`
	BatchInputPropertyName shared.BatchInputPropertyNameParam
	paramObj
}

func (r MediaBridgePropertyArchiveBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPropertyName)
}
func (r *MediaBridgePropertyArchiveBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputPropertyName)
}

type MediaBridgePropertyNewBatchParams struct {
	AppID                    string `path:"appId,required" json:"-"`
	BatchInputPropertyCreate shared.BatchInputPropertyCreateParam
	paramObj
}

func (r MediaBridgePropertyNewBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPropertyCreate)
}
func (r *MediaBridgePropertyNewBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputPropertyCreate)
}

type MediaBridgePropertyGetParams struct {
	AppID      string `path:"appId,required" json:"-"`
	ObjectType string `path:"objectType,required" json:"-"`
	paramObj
}

type MediaBridgePropertyGetBatchParams struct {
	AppID                      string `path:"appId,required" json:"-"`
	BatchReadInputPropertyName crm.BatchReadInputPropertyNameParam
	paramObj
}

func (r MediaBridgePropertyGetBatchParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchReadInputPropertyName)
}
func (r *MediaBridgePropertyGetBatchParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchReadInputPropertyName)
}
