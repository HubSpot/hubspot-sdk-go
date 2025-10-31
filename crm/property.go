// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/marketing"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// PropertyService contains methods and other services that help with interacting
// with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPropertyService] method instead.
type PropertyService struct {
	Options []option.RequestOption
	Batch   PropertyBatchService
	Groups  PropertyGroupService
}

// NewPropertyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPropertyService(opts ...option.RequestOption) (r PropertyService) {
	r = PropertyService{}
	r.Options = opts
	r.Batch = NewPropertyBatchService(opts...)
	r.Groups = NewPropertyGroupService(opts...)
	return
}

// Create and return a copy of a new property for the specified object type.
func (r *PropertyService) New(ctx context.Context, objectType string, body PropertyNewParams, opts ...option.RequestOption) (res *CreatedResponseProperty, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/properties/%s", objectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Perform a partial update of a property identified by { propertyName }. Provided
// fields will be overwritten.
func (r *PropertyService) Update(ctx context.Context, propertyName string, params PropertyUpdateParams, opts ...option.RequestOption) (res *shared.Property, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	if propertyName == "" {
		err = errors.New("missing required propertyName parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/properties/%s/%s", params.ObjectType, propertyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Read all existing properties for the specified object type and HubSpot account.
func (r *PropertyService) List(ctx context.Context, objectType string, query PropertyListParams, opts ...option.RequestOption) (res *CollectionResponseProperty, err error) {
	opts = slices.Concat(r.Options, opts)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/properties/%s", objectType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Move a property identified by {propertyName} to the recycling bin.
func (r *PropertyService) Delete(ctx context.Context, propertyName string, body PropertyDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	if propertyName == "" {
		err = errors.New("missing required propertyName parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/properties/%s/%s", body.ObjectType, propertyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Read a property identified by {propertyName}.
func (r *PropertyService) Get(ctx context.Context, propertyName string, params PropertyGetParams, opts ...option.RequestOption) (res *shared.Property, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return
	}
	if propertyName == "" {
		err = errors.New("missing required propertyName parameter")
		return
	}
	path := fmt.Sprintf("crm/v3/properties/%s/%s", params.ObjectType, propertyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// The properties Archived, Inputs are required.
type BatchReadInputPropertyNameParam struct {
	Archived bool                       `json:"archived,required"`
	Inputs   []shared.PropertyNameParam `json:"inputs,omitzero,required"`
	// Any of "non_sensitive", "sensitive", "highly_sensitive".
	DataSensitivity BatchReadInputPropertyNameDataSensitivity `json:"dataSensitivity,omitzero"`
	paramObj
}

func (r BatchReadInputPropertyNameParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchReadInputPropertyNameParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchReadInputPropertyNameParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchReadInputPropertyNameDataSensitivity string

const (
	BatchReadInputPropertyNameDataSensitivityNonSensitive    BatchReadInputPropertyNameDataSensitivity = "non_sensitive"
	BatchReadInputPropertyNameDataSensitivitySensitive       BatchReadInputPropertyNameDataSensitivity = "sensitive"
	BatchReadInputPropertyNameDataSensitivityHighlySensitive BatchReadInputPropertyNameDataSensitivity = "highly_sensitive"
)

type CollectionResponseProperty struct {
	Results []shared.Property `json:"results,required"`
	// Contains information pagination of results.
	Paging marketing.Paging `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseProperty) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseProperty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponsePropertyGroup struct {
	Results []PropertyGroup `json:"results,required"`
	// Contains information pagination of results.
	Paging marketing.Paging `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponsePropertyGroup) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponsePropertyGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreatedResponseProperty struct {
	CreatedResourceID string `json:"createdResourceId,required"`
	// Defines a property
	Entity   shared.Property `json:"entity,required"`
	Location string          `json:"location"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedResourceID respjson.Field
		Entity            respjson.Field
		Location          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreatedResponseProperty) RawJSON() string { return r.JSON.raw }
func (r *CreatedResponseProperty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreatedResponsePropertyGroup struct {
	CreatedResourceID string `json:"createdResourceId,required"`
	// An ID for a group of properties
	Entity   PropertyGroup `json:"entity,required"`
	Location string        `json:"location"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedResourceID respjson.Field
		Entity            respjson.Field
		Location          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreatedResponsePropertyGroup) RawJSON() string { return r.JSON.raw }
func (r *CreatedResponsePropertyGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An ID for a group of properties
type PropertyGroup struct {
	Archived bool `json:"archived,required"`
	// Property groups are displayed in order starting with the lowest positive integer
	// value. Values of -1 will cause the property group to be displayed after any
	// positive values.
	DisplayOrder int64 `json:"displayOrder,required"`
	// A human-readable label that will be shown in HubSpot.
	Label string `json:"label,required"`
	// The internal property group name, which must be used when referencing the
	// property group via the API.
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Archived     respjson.Field
		DisplayOrder respjson.Field
		Label        respjson.Field
		Name         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PropertyGroup) RawJSON() string { return r.JSON.raw }
func (r *PropertyGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PropertyUpdateParam struct {
	// Represents a formula that is used to compute a calculated property.
	CalculationFormula param.Opt[string] `json:"calculationFormula,omitzero"`
	// A description of the property that will be shown as help text in HubSpot.
	Description param.Opt[string] `json:"description,omitzero"`
	// Properties are displayed in order starting with the lowest positive integer
	// value. Values of -1 will cause the Property to be displayed after any positive
	// values.
	DisplayOrder param.Opt[int64] `json:"displayOrder,omitzero"`
	// Whether or not the property can be used in a HubSpot form.
	FormField param.Opt[bool] `json:"formField,omitzero"`
	// The name of the property group the property belongs to.
	GroupName param.Opt[string] `json:"groupName,omitzero"`
	// If true, the property won't be visible and can't be used in HubSpot.
	Hidden param.Opt[bool] `json:"hidden,omitzero"`
	// A human-readable property label that will be shown in HubSpot.
	Label param.Opt[string] `json:"label,omitzero"`
	// Controls how the property appears in HubSpot.
	//
	// Any of "booleancheckbox", "calculation_equation", "checkbox", "date", "file",
	// "html", "number", "phonenumber", "radio", "select", "text", "textarea".
	FieldType PropertyUpdateFieldType `json:"fieldType,omitzero"`
	// A list of valid options for the property.
	Options []shared.OptionInputParam `json:"options,omitzero"`
	// The data type of the property.
	//
	// Any of "bool", "date", "datetime", "enumeration", "number", "phone_number",
	// "string".
	Type PropertyUpdateType `json:"type,omitzero"`
	paramObj
}

func (r PropertyUpdateParam) MarshalJSON() (data []byte, err error) {
	type shadow PropertyUpdateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PropertyUpdateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls how the property appears in HubSpot.
type PropertyUpdateFieldType string

const (
	PropertyUpdateFieldTypeBooleancheckbox     PropertyUpdateFieldType = "booleancheckbox"
	PropertyUpdateFieldTypeCalculationEquation PropertyUpdateFieldType = "calculation_equation"
	PropertyUpdateFieldTypeCheckbox            PropertyUpdateFieldType = "checkbox"
	PropertyUpdateFieldTypeDate                PropertyUpdateFieldType = "date"
	PropertyUpdateFieldTypeFile                PropertyUpdateFieldType = "file"
	PropertyUpdateFieldTypeHTML                PropertyUpdateFieldType = "html"
	PropertyUpdateFieldTypeNumber              PropertyUpdateFieldType = "number"
	PropertyUpdateFieldTypePhonenumber         PropertyUpdateFieldType = "phonenumber"
	PropertyUpdateFieldTypeRadio               PropertyUpdateFieldType = "radio"
	PropertyUpdateFieldTypeSelect              PropertyUpdateFieldType = "select"
	PropertyUpdateFieldTypeText                PropertyUpdateFieldType = "text"
	PropertyUpdateFieldTypeTextarea            PropertyUpdateFieldType = "textarea"
)

// The data type of the property.
type PropertyUpdateType string

const (
	PropertyUpdateTypeBool        PropertyUpdateType = "bool"
	PropertyUpdateTypeDate        PropertyUpdateType = "date"
	PropertyUpdateTypeDatetime    PropertyUpdateType = "datetime"
	PropertyUpdateTypeEnumeration PropertyUpdateType = "enumeration"
	PropertyUpdateTypeNumber      PropertyUpdateType = "number"
	PropertyUpdateTypePhoneNumber PropertyUpdateType = "phone_number"
	PropertyUpdateTypeString      PropertyUpdateType = "string"
)

type PropertyNewParams struct {
	PropertyCreate shared.PropertyCreateParam
	paramObj
}

func (r PropertyNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PropertyCreate)
}
func (r *PropertyNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PropertyCreate)
}

type PropertyUpdateParams struct {
	ObjectType     string `path:"objectType,required" json:"-"`
	PropertyUpdate PropertyUpdateParam
	paramObj
}

func (r PropertyUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PropertyUpdate)
}
func (r *PropertyUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PropertyUpdate)
}

type PropertyListParams struct {
	// Whether to return only results that have been archived.
	Archived   param.Opt[bool]   `query:"archived,omitzero" json:"-"`
	Properties param.Opt[string] `query:"properties,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PropertyListParams]'s query parameters as `url.Values`.
func (r PropertyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PropertyDeleteParams struct {
	ObjectType string `path:"objectType,required" json:"-"`
	paramObj
}

type PropertyGetParams struct {
	ObjectType string `path:"objectType,required" json:"-"`
	// Whether to return only results that have been archived.
	Archived   param.Opt[bool]   `query:"archived,omitzero" json:"-"`
	Properties param.Opt[string] `query:"properties,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PropertyGetParams]'s query parameters as `url.Values`.
func (r PropertyGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
