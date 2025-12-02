// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// AssociationSchemaV4Service contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssociationSchemaV4Service] method instead.
type AssociationSchemaV4Service struct {
	Options        []option.RequestOption
	Configurations AssociationSchemaV4ConfigurationService
	Definitions    AssociationSchemaV4DefinitionService
}

// NewAssociationSchemaV4Service generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAssociationSchemaV4Service(opts ...option.RequestOption) (r AssociationSchemaV4Service) {
	r = AssociationSchemaV4Service{}
	r.Options = opts
	r.Configurations = NewAssociationSchemaV4ConfigurationService(opts...)
	r.Definitions = NewAssociationSchemaV4DefinitionService(opts...)
	return
}

// The property Inputs is required.
type BatchInputPublicAssociationDefinitionConfigurationCreateRequestParam struct {
	Inputs []PublicAssociationDefinitionConfigurationCreateRequestParam `json:"inputs,omitzero,required"`
	paramObj
}

func (r BatchInputPublicAssociationDefinitionConfigurationCreateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputPublicAssociationDefinitionConfigurationCreateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputPublicAssociationDefinitionConfigurationCreateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type BatchInputPublicAssociationDefinitionConfigurationUpdateRequestParam struct {
	Inputs []PublicAssociationDefinitionConfigurationUpdateRequestParam `json:"inputs,omitzero,required"`
	paramObj
}

func (r BatchInputPublicAssociationDefinitionConfigurationUpdateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputPublicAssociationDefinitionConfigurationUpdateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputPublicAssociationDefinitionConfigurationUpdateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type BatchInputPublicAssociationSpecParam struct {
	Inputs []PublicAssociationSpecParam `json:"inputs,omitzero,required"`
	paramObj
}

func (r BatchInputPublicAssociationSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputPublicAssociationSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputPublicAssociationSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponsePublicAssociationDefinitionConfigurationUpdateResult struct {
	CompletedAt time.Time                                              `json:"completedAt,required" format:"date-time"`
	Results     []PublicAssociationDefinitionConfigurationUpdateResult `json:"results,required"`
	StartedAt   time.Time                                              `json:"startedAt,required" format:"date-time"`
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status      BatchResponsePublicAssociationDefinitionConfigurationUpdateResultStatus `json:"status,required"`
	Errors      []shared.StandardError                                                  `json:"errors"`
	Links       map[string]string                                                       `json:"links"`
	NumErrors   int64                                                                   `json:"numErrors"`
	RequestedAt time.Time                                                               `json:"requestedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt respjson.Field
		Results     respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		Errors      respjson.Field
		Links       respjson.Field
		NumErrors   respjson.Field
		RequestedAt respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchResponsePublicAssociationDefinitionConfigurationUpdateResult) RawJSON() string {
	return r.JSON.raw
}
func (r *BatchResponsePublicAssociationDefinitionConfigurationUpdateResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponsePublicAssociationDefinitionConfigurationUpdateResultStatus string

const (
	BatchResponsePublicAssociationDefinitionConfigurationUpdateResultStatusCanceled   BatchResponsePublicAssociationDefinitionConfigurationUpdateResultStatus = "CANCELED"
	BatchResponsePublicAssociationDefinitionConfigurationUpdateResultStatusComplete   BatchResponsePublicAssociationDefinitionConfigurationUpdateResultStatus = "COMPLETE"
	BatchResponsePublicAssociationDefinitionConfigurationUpdateResultStatusPending    BatchResponsePublicAssociationDefinitionConfigurationUpdateResultStatus = "PENDING"
	BatchResponsePublicAssociationDefinitionConfigurationUpdateResultStatusProcessing BatchResponsePublicAssociationDefinitionConfigurationUpdateResultStatus = "PROCESSING"
)

type BatchResponsePublicAssociationDefinitionUserConfiguration struct {
	CompletedAt time.Time                                      `json:"completedAt,required" format:"date-time"`
	Results     []PublicAssociationDefinitionUserConfiguration `json:"results,required"`
	StartedAt   time.Time                                      `json:"startedAt,required" format:"date-time"`
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status      BatchResponsePublicAssociationDefinitionUserConfigurationStatus `json:"status,required"`
	Errors      []shared.StandardError                                          `json:"errors"`
	Links       map[string]string                                               `json:"links"`
	NumErrors   int64                                                           `json:"numErrors"`
	RequestedAt time.Time                                                       `json:"requestedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt respjson.Field
		Results     respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		Errors      respjson.Field
		Links       respjson.Field
		NumErrors   respjson.Field
		RequestedAt respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchResponsePublicAssociationDefinitionUserConfiguration) RawJSON() string {
	return r.JSON.raw
}
func (r *BatchResponsePublicAssociationDefinitionUserConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponsePublicAssociationDefinitionUserConfigurationStatus string

const (
	BatchResponsePublicAssociationDefinitionUserConfigurationStatusCanceled   BatchResponsePublicAssociationDefinitionUserConfigurationStatus = "CANCELED"
	BatchResponsePublicAssociationDefinitionUserConfigurationStatusComplete   BatchResponsePublicAssociationDefinitionUserConfigurationStatus = "COMPLETE"
	BatchResponsePublicAssociationDefinitionUserConfigurationStatusPending    BatchResponsePublicAssociationDefinitionUserConfigurationStatus = "PENDING"
	BatchResponsePublicAssociationDefinitionUserConfigurationStatusProcessing BatchResponsePublicAssociationDefinitionUserConfigurationStatus = "PROCESSING"
)

type CollectionResponseAssociationSpecWithLabel struct {
	Results []AssociationSpecWithLabel `json:"results,required"`
	Paging  shared.Paging              `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseAssociationSpecWithLabel) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseAssociationSpecWithLabel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponsePublicAssociationDefinitionUserConfiguration struct {
	Results []PublicAssociationDefinitionUserConfiguration `json:"results,required"`
	Paging  shared.Paging                                  `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponsePublicAssociationDefinitionUserConfiguration) RawJSON() string {
	return r.JSON.raw
}
func (r *CollectionResponsePublicAssociationDefinitionUserConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Category, MaxToObjectIDs, TypeID are required.
type PublicAssociationDefinitionConfigurationCreateRequestParam struct {
	// Any of "HUBSPOT_DEFINED", "INTEGRATOR_DEFINED", "USER_DEFINED".
	Category       PublicAssociationDefinitionConfigurationCreateRequestCategory `json:"category,omitzero,required"`
	MaxToObjectIDs int64                                                         `json:"maxToObjectIds,required"`
	TypeID         int64                                                         `json:"typeId,required"`
	paramObj
}

func (r PublicAssociationDefinitionConfigurationCreateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicAssociationDefinitionConfigurationCreateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicAssociationDefinitionConfigurationCreateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicAssociationDefinitionConfigurationCreateRequestCategory string

const (
	PublicAssociationDefinitionConfigurationCreateRequestCategoryHubspotDefined    PublicAssociationDefinitionConfigurationCreateRequestCategory = "HUBSPOT_DEFINED"
	PublicAssociationDefinitionConfigurationCreateRequestCategoryIntegratorDefined PublicAssociationDefinitionConfigurationCreateRequestCategory = "INTEGRATOR_DEFINED"
	PublicAssociationDefinitionConfigurationCreateRequestCategoryUserDefined       PublicAssociationDefinitionConfigurationCreateRequestCategory = "USER_DEFINED"
)

// The properties Category, MaxToObjectIDs, TypeID are required.
type PublicAssociationDefinitionConfigurationUpdateRequestParam struct {
	// Any of "HUBSPOT_DEFINED", "INTEGRATOR_DEFINED", "USER_DEFINED".
	Category       PublicAssociationDefinitionConfigurationUpdateRequestCategory `json:"category,omitzero,required"`
	MaxToObjectIDs int64                                                         `json:"maxToObjectIds,required"`
	TypeID         int64                                                         `json:"typeId,required"`
	paramObj
}

func (r PublicAssociationDefinitionConfigurationUpdateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicAssociationDefinitionConfigurationUpdateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicAssociationDefinitionConfigurationUpdateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicAssociationDefinitionConfigurationUpdateRequestCategory string

const (
	PublicAssociationDefinitionConfigurationUpdateRequestCategoryHubspotDefined    PublicAssociationDefinitionConfigurationUpdateRequestCategory = "HUBSPOT_DEFINED"
	PublicAssociationDefinitionConfigurationUpdateRequestCategoryIntegratorDefined PublicAssociationDefinitionConfigurationUpdateRequestCategory = "INTEGRATOR_DEFINED"
	PublicAssociationDefinitionConfigurationUpdateRequestCategoryUserDefined       PublicAssociationDefinitionConfigurationUpdateRequestCategory = "USER_DEFINED"
)

type PublicAssociationDefinitionConfigurationUpdateResult struct {
	// Any of "HUBSPOT_DEFINED", "INTEGRATOR_DEFINED", "USER_DEFINED".
	Category                   PublicAssociationDefinitionConfigurationUpdateResultCategory `json:"category,required"`
	TypeID                     int64                                                        `json:"typeId,required"`
	UserEnforcedMaxToObjectIDs int64                                                        `json:"userEnforcedMaxToObjectIds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Category                   respjson.Field
		TypeID                     respjson.Field
		UserEnforcedMaxToObjectIDs respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicAssociationDefinitionConfigurationUpdateResult) RawJSON() string { return r.JSON.raw }
func (r *PublicAssociationDefinitionConfigurationUpdateResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicAssociationDefinitionConfigurationUpdateResultCategory string

const (
	PublicAssociationDefinitionConfigurationUpdateResultCategoryHubspotDefined    PublicAssociationDefinitionConfigurationUpdateResultCategory = "HUBSPOT_DEFINED"
	PublicAssociationDefinitionConfigurationUpdateResultCategoryIntegratorDefined PublicAssociationDefinitionConfigurationUpdateResultCategory = "INTEGRATOR_DEFINED"
	PublicAssociationDefinitionConfigurationUpdateResultCategoryUserDefined       PublicAssociationDefinitionConfigurationUpdateResultCategory = "USER_DEFINED"
)

// The properties Label, Name are required.
type PublicAssociationDefinitionCreateRequestParam struct {
	Label        string            `json:"label,required"`
	Name         string            `json:"name,required"`
	InverseLabel param.Opt[string] `json:"inverseLabel,omitzero"`
	paramObj
}

func (r PublicAssociationDefinitionCreateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicAssociationDefinitionCreateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicAssociationDefinitionCreateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AssociationTypeID, Label are required.
type PublicAssociationDefinitionUpdateRequestParam struct {
	AssociationTypeID int64             `json:"associationTypeId,required"`
	Label             string            `json:"label,required"`
	InverseLabel      param.Opt[string] `json:"inverseLabel,omitzero"`
	paramObj
}

func (r PublicAssociationDefinitionUpdateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicAssociationDefinitionUpdateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicAssociationDefinitionUpdateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicAssociationDefinitionUserConfiguration struct {
	// Any of "HUBSPOT_DEFINED", "INTEGRATOR_DEFINED", "USER_DEFINED".
	Category                   PublicAssociationDefinitionUserConfigurationCategory `json:"category,required"`
	TypeID                     int64                                                `json:"typeId,required"`
	Label                      string                                               `json:"label"`
	UserEnforcedMaxToObjectIDs int64                                                `json:"userEnforcedMaxToObjectIds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Category                   respjson.Field
		TypeID                     respjson.Field
		Label                      respjson.Field
		UserEnforcedMaxToObjectIDs respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicAssociationDefinitionUserConfiguration) RawJSON() string { return r.JSON.raw }
func (r *PublicAssociationDefinitionUserConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicAssociationDefinitionUserConfigurationCategory string

const (
	PublicAssociationDefinitionUserConfigurationCategoryHubspotDefined    PublicAssociationDefinitionUserConfigurationCategory = "HUBSPOT_DEFINED"
	PublicAssociationDefinitionUserConfigurationCategoryIntegratorDefined PublicAssociationDefinitionUserConfigurationCategory = "INTEGRATOR_DEFINED"
	PublicAssociationDefinitionUserConfigurationCategoryUserDefined       PublicAssociationDefinitionUserConfigurationCategory = "USER_DEFINED"
)

// The properties Category, TypeID are required.
type PublicAssociationSpecParam struct {
	Category string `json:"category,required"`
	TypeID   int64  `json:"typeId,required"`
	paramObj
}

func (r PublicAssociationSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicAssociationSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicAssociationSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
