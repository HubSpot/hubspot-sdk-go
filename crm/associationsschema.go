// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"time"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
	"github.com/HubSpot/hubspot-sdk-go/packages/respjson"
)

// AssociationsSchemaService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssociationsSchemaService] method instead.
type AssociationsSchemaService struct {
	options []option.RequestOption
	Labels  AssociationsSchemaLabelService
	Limits  AssociationsSchemaLimitService
}

// NewAssociationsSchemaService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAssociationsSchemaService(opts ...option.RequestOption) (r AssociationsSchemaService) {
	r = AssociationsSchemaService{}
	r.options = opts
	r.Labels = NewAssociationsSchemaLabelService(opts...)
	r.Limits = NewAssociationsSchemaLimitService(opts...)
	return
}

// The property Inputs is required.
type BatchInputPublicAssociationDefinitionConfigurationCreateRequestParam struct {
	Inputs []PublicAssociationDefinitionConfigurationCreateRequestParam `json:"inputs,omitzero" api:"required"`
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
	Inputs []PublicAssociationDefinitionConfigurationUpdateRequestParam `json:"inputs,omitzero" api:"required"`
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
	Inputs []PublicAssociationSpecParam `json:"inputs,omitzero" api:"required"`
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
	// The date and time when the batch update operation was completed.
	CompletedAt time.Time                                              `json:"completedAt" api:"required" format:"date-time"`
	Results     []PublicAssociationDefinitionConfigurationUpdateResult `json:"results" api:"required"`
	// The date and time when the batch update operation started.
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// The current status of the batch update operation, which can be CANCELED,
	// COMPLETE, PENDING, or PROCESSING.
	//
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status BatchResponsePublicAssociationDefinitionConfigurationUpdateResultStatus `json:"status" api:"required"`
	// URLs linking to documentation or resources associated with the batch update
	// operation.
	Links map[string]string `json:"links"`
	// The date and time when the batch update operation was requested.
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
func (r BatchResponsePublicAssociationDefinitionConfigurationUpdateResult) RawJSON() string {
	return r.JSON.raw
}
func (r *BatchResponsePublicAssociationDefinitionConfigurationUpdateResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current status of the batch update operation, which can be CANCELED,
// COMPLETE, PENDING, or PROCESSING.
type BatchResponsePublicAssociationDefinitionConfigurationUpdateResultStatus string

const (
	BatchResponsePublicAssociationDefinitionConfigurationUpdateResultStatusCanceled   BatchResponsePublicAssociationDefinitionConfigurationUpdateResultStatus = "CANCELED"
	BatchResponsePublicAssociationDefinitionConfigurationUpdateResultStatusComplete   BatchResponsePublicAssociationDefinitionConfigurationUpdateResultStatus = "COMPLETE"
	BatchResponsePublicAssociationDefinitionConfigurationUpdateResultStatusPending    BatchResponsePublicAssociationDefinitionConfigurationUpdateResultStatus = "PENDING"
	BatchResponsePublicAssociationDefinitionConfigurationUpdateResultStatusProcessing BatchResponsePublicAssociationDefinitionConfigurationUpdateResultStatus = "PROCESSING"
)

type BatchResponsePublicAssociationDefinitionUserConfiguration struct {
	// The date and time when the batch operation was completed.
	CompletedAt time.Time                                      `json:"completedAt" api:"required" format:"date-time"`
	Results     []PublicAssociationDefinitionUserConfiguration `json:"results" api:"required"`
	// The date and time when the batch operation started.
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// The current status of the batch operation, which can be CANCELED, COMPLETE,
	// PENDING, or PROCESSING.
	//
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status BatchResponsePublicAssociationDefinitionUserConfigurationStatus `json:"status" api:"required"`
	// A collection of URLs linking to related documentation or resources associated
	// with the batch operation.
	Links map[string]string `json:"links"`
	// The date and time when the batch operation was requested.
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
func (r BatchResponsePublicAssociationDefinitionUserConfiguration) RawJSON() string {
	return r.JSON.raw
}
func (r *BatchResponsePublicAssociationDefinitionUserConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current status of the batch operation, which can be CANCELED, COMPLETE,
// PENDING, or PROCESSING.
type BatchResponsePublicAssociationDefinitionUserConfigurationStatus string

const (
	BatchResponsePublicAssociationDefinitionUserConfigurationStatusCanceled   BatchResponsePublicAssociationDefinitionUserConfigurationStatus = "CANCELED"
	BatchResponsePublicAssociationDefinitionUserConfigurationStatusComplete   BatchResponsePublicAssociationDefinitionUserConfigurationStatus = "COMPLETE"
	BatchResponsePublicAssociationDefinitionUserConfigurationStatusPending    BatchResponsePublicAssociationDefinitionUserConfigurationStatus = "PENDING"
	BatchResponsePublicAssociationDefinitionUserConfigurationStatusProcessing BatchResponsePublicAssociationDefinitionUserConfigurationStatus = "PROCESSING"
)

type CollectionResponseAssociationSpecWithLabelNoPaging struct {
	Results []AssociationSpecWithLabel `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseAssociationSpecWithLabelNoPaging) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseAssociationSpecWithLabelNoPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponsePublicAssociationDefinitionUserConfigurationNoPaging struct {
	Results []PublicAssociationDefinitionUserConfiguration `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponsePublicAssociationDefinitionUserConfigurationNoPaging) RawJSON() string {
	return r.JSON.raw
}
func (r *CollectionResponsePublicAssociationDefinitionUserConfigurationNoPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Category, MaxToObjectIDs, TypeID are required.
type PublicAssociationDefinitionConfigurationCreateRequestParam struct {
	// Specifies the category of the association, which can be HUBSPOT_DEFINED,
	// INTEGRATOR_DEFINED, or USER_DEFINED.
	//
	// Any of "HUBSPOT_DEFINED", "INTEGRATOR_DEFINED", "USER_DEFINED", "WORK".
	Category PublicAssociationDefinitionConfigurationCreateRequestCategory `json:"category,omitzero" api:"required"`
	// The maximum number of target object IDs that can be associated with a single
	// source object.
	MaxToObjectIDs int64 `json:"maxToObjectIds" api:"required"`
	// An integer used to uniquely identify a specific association type within its
	// category.
	TypeID int64 `json:"typeId" api:"required"`
	paramObj
}

func (r PublicAssociationDefinitionConfigurationCreateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicAssociationDefinitionConfigurationCreateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicAssociationDefinitionConfigurationCreateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the category of the association, which can be HUBSPOT_DEFINED,
// INTEGRATOR_DEFINED, or USER_DEFINED.
type PublicAssociationDefinitionConfigurationCreateRequestCategory string

const (
	PublicAssociationDefinitionConfigurationCreateRequestCategoryHubSpotDefined    PublicAssociationDefinitionConfigurationCreateRequestCategory = "HUBSPOT_DEFINED"
	PublicAssociationDefinitionConfigurationCreateRequestCategoryIntegratorDefined PublicAssociationDefinitionConfigurationCreateRequestCategory = "INTEGRATOR_DEFINED"
	PublicAssociationDefinitionConfigurationCreateRequestCategoryUserDefined       PublicAssociationDefinitionConfigurationCreateRequestCategory = "USER_DEFINED"
	PublicAssociationDefinitionConfigurationCreateRequestCategoryWork              PublicAssociationDefinitionConfigurationCreateRequestCategory = "WORK"
)

// The properties Category, MaxToObjectIDs, TypeID are required.
type PublicAssociationDefinitionConfigurationUpdateRequestParam struct {
	// Specifies the category of the association, which can be HUBSPOT_DEFINED,
	// INTEGRATOR_DEFINED, or USER_DEFINED.
	//
	// Any of "HUBSPOT_DEFINED", "INTEGRATOR_DEFINED", "USER_DEFINED", "WORK".
	Category PublicAssociationDefinitionConfigurationUpdateRequestCategory `json:"category,omitzero" api:"required"`
	// Defines the maximum number of target object IDs that can be associated.
	MaxToObjectIDs int64 `json:"maxToObjectIds" api:"required"`
	// A unique identifier for the association type.
	TypeID int64 `json:"typeId" api:"required"`
	paramObj
}

func (r PublicAssociationDefinitionConfigurationUpdateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicAssociationDefinitionConfigurationUpdateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicAssociationDefinitionConfigurationUpdateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the category of the association, which can be HUBSPOT_DEFINED,
// INTEGRATOR_DEFINED, or USER_DEFINED.
type PublicAssociationDefinitionConfigurationUpdateRequestCategory string

const (
	PublicAssociationDefinitionConfigurationUpdateRequestCategoryHubSpotDefined    PublicAssociationDefinitionConfigurationUpdateRequestCategory = "HUBSPOT_DEFINED"
	PublicAssociationDefinitionConfigurationUpdateRequestCategoryIntegratorDefined PublicAssociationDefinitionConfigurationUpdateRequestCategory = "INTEGRATOR_DEFINED"
	PublicAssociationDefinitionConfigurationUpdateRequestCategoryUserDefined       PublicAssociationDefinitionConfigurationUpdateRequestCategory = "USER_DEFINED"
	PublicAssociationDefinitionConfigurationUpdateRequestCategoryWork              PublicAssociationDefinitionConfigurationUpdateRequestCategory = "WORK"
)

type PublicAssociationDefinitionConfigurationUpdateResult struct {
	// The category of the association, which can be HUBSPOT_DEFINED,
	// INTEGRATOR_DEFINED, or USER_DEFINED.
	//
	// Any of "HUBSPOT_DEFINED", "INTEGRATOR_DEFINED", "USER_DEFINED", "WORK".
	Category PublicAssociationDefinitionConfigurationUpdateResultCategory `json:"category" api:"required"`
	// An integer value used to uniquely identify a specific association type within
	// its Association Category.
	TypeID int64 `json:"typeId" api:"required"`
	// The maximum number of object IDs that a user can enforce for associations.
	UserEnforcedMaxToObjectIDs int64 `json:"userEnforcedMaxToObjectIds"`
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

// The category of the association, which can be HUBSPOT_DEFINED,
// INTEGRATOR_DEFINED, or USER_DEFINED.
type PublicAssociationDefinitionConfigurationUpdateResultCategory string

const (
	PublicAssociationDefinitionConfigurationUpdateResultCategoryHubSpotDefined    PublicAssociationDefinitionConfigurationUpdateResultCategory = "HUBSPOT_DEFINED"
	PublicAssociationDefinitionConfigurationUpdateResultCategoryIntegratorDefined PublicAssociationDefinitionConfigurationUpdateResultCategory = "INTEGRATOR_DEFINED"
	PublicAssociationDefinitionConfigurationUpdateResultCategoryUserDefined       PublicAssociationDefinitionConfigurationUpdateResultCategory = "USER_DEFINED"
	PublicAssociationDefinitionConfigurationUpdateResultCategoryWork              PublicAssociationDefinitionConfigurationUpdateResultCategory = "WORK"
)

// The properties Label, Name are required.
type PublicAssociationDefinitionCreateRequestParam struct {
	// A descriptor that provides context about the relationship between two associated
	// CRM objects.
	Label string `json:"label" api:"required"`
	// The unique identifier for the association definition.
	Name string `json:"name" api:"required"`
	// An optional descriptor that clarifies the reverse relationship in the
	// association.
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
	// The unique identifier for the association type.
	AssociationTypeID int64 `json:"associationTypeId" api:"required"`
	// A descriptor that provides context about the relationship between associated
	// records.
	Label string `json:"label" api:"required"`
	// An optional descriptor for the inverse relationship between associated records.
	InverseLabel param.Opt[string] `json:"inverseLabel,omitzero"`
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
	// The category of the association, which can be HUBSPOT_DEFINED,
	// INTEGRATOR_DEFINED, or USER_DEFINED.
	//
	// Any of "HUBSPOT_DEFINED", "INTEGRATOR_DEFINED", "USER_DEFINED", "WORK".
	Category PublicAssociationDefinitionUserConfigurationCategory `json:"category" api:"required"`
	// A unique integer identifier for the association type within its category.
	TypeID int64 `json:"typeId" api:"required"`
	// A descriptor providing context about the relationship between associated
	// records.
	Label string `json:"label"`
	// The maximum number of target object IDs that a user can enforce in an
	// association.
	UserEnforcedMaxToObjectIDs int64 `json:"userEnforcedMaxToObjectIds"`
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

// The category of the association, which can be HUBSPOT_DEFINED,
// INTEGRATOR_DEFINED, or USER_DEFINED.
type PublicAssociationDefinitionUserConfigurationCategory string

const (
	PublicAssociationDefinitionUserConfigurationCategoryHubSpotDefined    PublicAssociationDefinitionUserConfigurationCategory = "HUBSPOT_DEFINED"
	PublicAssociationDefinitionUserConfigurationCategoryIntegratorDefined PublicAssociationDefinitionUserConfigurationCategory = "INTEGRATOR_DEFINED"
	PublicAssociationDefinitionUserConfigurationCategoryUserDefined       PublicAssociationDefinitionUserConfigurationCategory = "USER_DEFINED"
	PublicAssociationDefinitionUserConfigurationCategoryWork              PublicAssociationDefinitionUserConfigurationCategory = "WORK"
)

// The properties Category, TypeID are required.
type PublicAssociationSpecParam struct {
	// Specifies the category of the association, which can be HUBSPOT_DEFINED,
	// INTEGRATOR_DEFINED, or USER_DEFINED.
	Category string `json:"category" api:"required"`
	// A unique integer identifier for the specific association type within its
	// category.
	TypeID int64 `json:"typeId" api:"required"`
	paramObj
}

func (r PublicAssociationSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicAssociationSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicAssociationSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
