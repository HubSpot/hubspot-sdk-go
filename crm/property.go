// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// PropertyService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPropertyService] method instead.
type PropertyService struct {
	options []option.RequestOption
	Batch   PropertyBatchService
	Groups  PropertyGroupService
}

// NewPropertyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPropertyService(opts ...option.RequestOption) (r PropertyService) {
	r = PropertyService{}
	r.options = opts
	r.Batch = NewPropertyBatchService(opts...)
	r.Groups = NewPropertyGroupService(opts...)
	return
}

// Create and return a copy of a new property for the specified object type.
func (r *PropertyService) New(ctx context.Context, objectType string, body PropertyNewParams, opts ...option.RequestOption) (res *shared.BaseProperty, err error) {
	opts = slices.Concat(r.options, opts)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/properties/2026-03/%s", url.PathEscape(objectType))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Perform a partial update of a property identified by { propertyName }. Provided
// fields will be overwritten.
func (r *PropertyService) Update(ctx context.Context, propertyName string, params PropertyUpdateParams, opts ...option.RequestOption) (res *shared.BaseProperty, err error) {
	opts = slices.Concat(r.options, opts)
	if params.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return nil, err
	}
	if propertyName == "" {
		err = errors.New("missing required propertyName parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/properties/2026-03/%s/%s", url.PathEscape(params.ObjectType), url.PathEscape(propertyName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Read all existing properties for the specified object type and HubSpot account.
func (r *PropertyService) List(ctx context.Context, objectType string, query PropertyListParams, opts ...option.RequestOption) (res *CollectionResponsePropertyNoPaging, err error) {
	opts = slices.Concat(r.options, opts)
	if objectType == "" {
		err = errors.New("missing required objectType parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/properties/2026-03/%s", url.PathEscape(objectType))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Move a property identified by {propertyName} to the recycling bin.
func (r *PropertyService) Delete(ctx context.Context, propertyName string, body PropertyDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return err
	}
	if propertyName == "" {
		err = errors.New("missing required propertyName parameter")
		return err
	}
	path := fmt.Sprintf("crm/properties/2026-03/%s/%s", url.PathEscape(body.ObjectType), url.PathEscape(propertyName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Read a property identified by {propertyName}.
func (r *PropertyService) Get(ctx context.Context, propertyName string, params PropertyGetParams, opts ...option.RequestOption) (res *shared.BaseProperty, err error) {
	opts = slices.Concat(r.options, opts)
	if params.ObjectType == "" {
		err = errors.New("missing required objectType parameter")
		return nil, err
	}
	if propertyName == "" {
		err = errors.New("missing required propertyName parameter")
		return nil, err
	}
	path := fmt.Sprintf("crm/properties/2026-03/%s/%s", url.PathEscape(params.ObjectType), url.PathEscape(propertyName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type BatchResponseProperty struct {
	// The timestamp indicating when the batch operation was completed.
	CompletedAt time.Time             `json:"completedAt" api:"required" format:"date-time"`
	Results     []shared.BaseProperty `json:"results" api:"required"`
	// The timestamp indicating when the batch operation began processing.
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// The current status of the batch operation, with possible values being CANCELED,
	// COMPLETE, PENDING, or PROCESSING.
	//
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status BatchResponsePropertyStatus `json:"status" api:"required"`
	// A collection of URLs linking to documentation or resources related to the batch
	// operation.
	Links map[string]string `json:"links"`
	// The timestamp indicating when the batch operation was requested.
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
func (r BatchResponseProperty) RawJSON() string { return r.JSON.raw }
func (r *BatchResponseProperty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current status of the batch operation, with possible values being CANCELED,
// COMPLETE, PENDING, or PROCESSING.
type BatchResponsePropertyStatus string

const (
	BatchResponsePropertyStatusCanceled   BatchResponsePropertyStatus = "CANCELED"
	BatchResponsePropertyStatusComplete   BatchResponsePropertyStatus = "COMPLETE"
	BatchResponsePropertyStatusPending    BatchResponsePropertyStatus = "PENDING"
	BatchResponsePropertyStatusProcessing BatchResponsePropertyStatus = "PROCESSING"
)

type CollectionResponsePropertyNoPaging struct {
	Results []shared.BaseProperty `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponsePropertyNoPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponsePropertyNoPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PropertyUpdateParam struct {
	// Represents a formula that is used to compute a calculated property.
	CalculationFormula   param.Opt[string] `json:"calculationFormula,omitzero"`
	CurrencyPropertyName param.Opt[string] `json:"currencyPropertyName,omitzero"`
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
	Label              param.Opt[string] `json:"label,omitzero"`
	ShowCurrencySymbol param.Opt[bool]   `json:"showCurrencySymbol,omitzero"`
	// Controls how the property appears in HubSpot.
	//
	// Any of "booleancheckbox", "calculation_equation", "checkbox", "date", "file",
	// "html", "number", "phonenumber", "radio", "select", "text", "textarea".
	FieldType PropertyUpdateFieldType `json:"fieldType,omitzero"`
	// Any of "currency", "duration", "formatted", "percentage", "probability",
	// "unformatted".
	NumberDisplayHint PropertyUpdateNumberDisplayHint `json:"numberDisplayHint,omitzero"`
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

type PropertyUpdateNumberDisplayHint string

const (
	PropertyUpdateNumberDisplayHintCurrency    PropertyUpdateNumberDisplayHint = "currency"
	PropertyUpdateNumberDisplayHintDuration    PropertyUpdateNumberDisplayHint = "duration"
	PropertyUpdateNumberDisplayHintFormatted   PropertyUpdateNumberDisplayHint = "formatted"
	PropertyUpdateNumberDisplayHintPercentage  PropertyUpdateNumberDisplayHint = "percentage"
	PropertyUpdateNumberDisplayHintProbability PropertyUpdateNumberDisplayHint = "probability"
	PropertyUpdateNumberDisplayHintUnformatted PropertyUpdateNumberDisplayHint = "unformatted"
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
	return apijson.UnmarshalRoot(data, r)
}

type PropertyUpdateParams struct {
	ObjectType     string `path:"objectType" api:"required" json:"-"`
	PropertyUpdate PropertyUpdateParam
	paramObj
}

func (r PropertyUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PropertyUpdate)
}
func (r *PropertyUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PropertyListParams struct {
	// Whether to return only results that have been archived.
	Archived   param.Opt[bool]   `query:"archived,omitzero" json:"-"`
	Locale     param.Opt[string] `query:"locale,omitzero" json:"-"`
	Properties param.Opt[string] `query:"properties,omitzero" json:"-"`
	// Any of "highly_sensitive", "non_sensitive", "sensitive".
	DataSensitivity PropertyListParamsDataSensitivity `query:"dataSensitivity,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PropertyListParams]'s query parameters as `url.Values`.
func (r PropertyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PropertyListParamsDataSensitivity string

const (
	PropertyListParamsDataSensitivityHighlySensitive PropertyListParamsDataSensitivity = "highly_sensitive"
	PropertyListParamsDataSensitivityNonSensitive    PropertyListParamsDataSensitivity = "non_sensitive"
	PropertyListParamsDataSensitivitySensitive       PropertyListParamsDataSensitivity = "sensitive"
)

type PropertyDeleteParams struct {
	ObjectType string `path:"objectType" api:"required" json:"-"`
	paramObj
}

type PropertyGetParams struct {
	ObjectType string `path:"objectType" api:"required" json:"-"`
	// Whether to return only results that have been archived.
	Archived   param.Opt[bool]   `query:"archived,omitzero" json:"-"`
	Locale     param.Opt[string] `query:"locale,omitzero" json:"-"`
	Properties param.Opt[string] `query:"properties,omitzero" json:"-"`
	// Any of "highly_sensitive", "non_sensitive", "sensitive".
	DataSensitivity PropertyGetParamsDataSensitivity `query:"dataSensitivity,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PropertyGetParams]'s query parameters as `url.Values`.
func (r PropertyGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PropertyGetParamsDataSensitivity string

const (
	PropertyGetParamsDataSensitivityHighlySensitive PropertyGetParamsDataSensitivity = "highly_sensitive"
	PropertyGetParamsDataSensitivityNonSensitive    PropertyGetParamsDataSensitivity = "non_sensitive"
	PropertyGetParamsDataSensitivitySensitive       PropertyGetParamsDataSensitivity = "sensitive"
)
