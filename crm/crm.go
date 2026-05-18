// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"time"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/packages/param"
	"github.com/HubSpot/hubspot-sdk-go/packages/respjson"
	"github.com/HubSpot/hubspot-sdk-go/shared"
)

// CrmService contains methods and other services that help with interacting with
// the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCrmService] method instead.
type CrmService struct {
	options               []option.RequestOption
	AppUninstalls         AppUninstallService
	Associations          AssociationService
	AssociationsSchema    AssociationsSchemaService
	DealSplits            DealSplitService
	Exports               ExportService
	Extensions            ExtensionService
	FeatureFlags          FeatureFlagService
	Imports               ImportService
	Limits                LimitService
	Lists                 ListService
	ObjectLibrary         ObjectLibraryService
	ObjectSchemas         ObjectSchemaService
	Objects               ObjectService
	Owners                OwnerService
	Pipelines             PipelineService
	Properties            PropertyService
	PropertiesValidations PropertiesValidationService
	Timeline              TimelineService
}

// NewCrmService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCrmService(opts ...option.RequestOption) (r CrmService) {
	r = CrmService{}
	r.options = opts
	r.AppUninstalls = NewAppUninstallService(opts...)
	r.Associations = NewAssociationService(opts...)
	r.AssociationsSchema = NewAssociationsSchemaService(opts...)
	r.DealSplits = NewDealSplitService(opts...)
	r.Exports = NewExportService(opts...)
	r.Extensions = NewExtensionService(opts...)
	r.FeatureFlags = NewFeatureFlagService(opts...)
	r.Imports = NewImportService(opts...)
	r.Limits = NewLimitService(opts...)
	r.Lists = NewListService(opts...)
	r.ObjectLibrary = NewObjectLibraryService(opts...)
	r.ObjectSchemas = NewObjectSchemaService(opts...)
	r.Objects = NewObjectService(opts...)
	r.Owners = NewOwnerService(opts...)
	r.Pipelines = NewPipelineService(opts...)
	r.Properties = NewPropertyService(opts...)
	r.PropertiesValidations = NewPropertiesValidationService(opts...)
	r.Timeline = NewTimelineService(opts...)
	return
}

// Defines the type, direction, and details of the relationship between two CRM
// objects.
type AssociationSpecWithLabel struct {
	// Association category. Can be HUBSPOT_DEFINED, USER_DEFINED, INTEGRATOR_DEFINED
	// or WORK
	//
	// Any of "HUBSPOT_DEFINED", "INTEGRATOR_DEFINED", "USER_DEFINED", "WORK".
	Category AssociationSpecWithLabelCategory `json:"category" api:"required"`
	// An integer value used to uniquely identify a specific association type within
	// its Association Category.
	TypeID int64 `json:"typeId" api:"required"`
	// An optional descriptor that provides additional context about the relationship
	// between associated records, such as "Mentor" and "Mentee".
	Label string `json:"label"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Category    respjson.Field
		TypeID      respjson.Field
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssociationSpecWithLabel) RawJSON() string { return r.JSON.raw }
func (r *AssociationSpecWithLabel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Association category. Can be HUBSPOT_DEFINED, USER_DEFINED, INTEGRATOR_DEFINED
// or WORK
type AssociationSpecWithLabelCategory string

const (
	AssociationSpecWithLabelCategoryHubSpotDefined    AssociationSpecWithLabelCategory = "HUBSPOT_DEFINED"
	AssociationSpecWithLabelCategoryIntegratorDefined AssociationSpecWithLabelCategory = "INTEGRATOR_DEFINED"
	AssociationSpecWithLabelCategoryUserDefined       AssociationSpecWithLabelCategory = "USER_DEFINED"
	AssociationSpecWithLabelCategoryWork              AssociationSpecWithLabelCategory = "WORK"
)

// The property Inputs is required.
type BatchInputPublicAssociationMultiArchiveParam struct {
	Inputs []PublicAssociationMultiArchiveParam `json:"inputs,omitzero" api:"required"`
	paramObj
}

func (r BatchInputPublicAssociationMultiArchiveParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputPublicAssociationMultiArchiveParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputPublicAssociationMultiArchiveParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type BatchInputPublicAssociationMultiPostParam struct {
	Inputs []PublicAssociationMultiPostParam `json:"inputs,omitzero" api:"required"`
	paramObj
}

func (r BatchInputPublicAssociationMultiPostParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputPublicAssociationMultiPostParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputPublicAssociationMultiPostParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type BatchInputPublicDefaultAssociationMultiPostParam struct {
	Inputs []PublicDefaultAssociationMultiPostParam `json:"inputs,omitzero" api:"required"`
	paramObj
}

func (r BatchInputPublicDefaultAssociationMultiPostParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputPublicDefaultAssociationMultiPostParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputPublicDefaultAssociationMultiPostParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type BatchInputPublicFetchAssociationsBatchRequestParam struct {
	Inputs []PublicFetchAssociationsBatchRequestParam `json:"inputs,omitzero" api:"required"`
	paramObj
}

func (r BatchInputPublicFetchAssociationsBatchRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputPublicFetchAssociationsBatchRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputPublicFetchAssociationsBatchRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponseLabelsBetweenObjectPair struct {
	// The timestamp when the batch processing was completed, in ISO 8601 format.
	CompletedAt time.Time                 `json:"completedAt" api:"required" format:"date-time"`
	Results     []LabelsBetweenObjectPair `json:"results" api:"required"`
	// The timestamp when the batch processing began, in ISO 8601 format.
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// The status of the batch processing request: "PENDING", "PROCESSING",
	// "CANCELLED", or "COMPLETE".
	//
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status BatchResponseLabelsBetweenObjectPairStatus `json:"status" api:"required"`
	// An object containing relevant links related to the batch request.
	Links map[string]string `json:"links"`
	// The timestamp when the batch request was initially made, in ISO 8601 format.
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
func (r BatchResponseLabelsBetweenObjectPair) RawJSON() string { return r.JSON.raw }
func (r *BatchResponseLabelsBetweenObjectPair) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the batch processing request: "PENDING", "PROCESSING",
// "CANCELLED", or "COMPLETE".
type BatchResponseLabelsBetweenObjectPairStatus string

const (
	BatchResponseLabelsBetweenObjectPairStatusCanceled   BatchResponseLabelsBetweenObjectPairStatus = "CANCELED"
	BatchResponseLabelsBetweenObjectPairStatusComplete   BatchResponseLabelsBetweenObjectPairStatus = "COMPLETE"
	BatchResponseLabelsBetweenObjectPairStatusPending    BatchResponseLabelsBetweenObjectPairStatus = "PENDING"
	BatchResponseLabelsBetweenObjectPairStatusProcessing BatchResponseLabelsBetweenObjectPairStatus = "PROCESSING"
)

type BatchResponsePublicAssociationMultiWithLabel struct {
	// The timestamp when the batch processing was completed, in ISO 8601 format.
	CompletedAt time.Time                         `json:"completedAt" api:"required" format:"date-time"`
	Results     []PublicAssociationMultiWithLabel `json:"results" api:"required"`
	// The timestamp when the batch processing began, in ISO 8601 format.
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// The status of the batch processing request: "PENDING", "PROCESSING", "CANCELED",
	// or "COMPLETE".
	//
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status BatchResponsePublicAssociationMultiWithLabelStatus `json:"status" api:"required"`
	// An object containing relevant links related to the batch request.
	Links map[string]string `json:"links"`
	// The timestamp when the batch request was initially made, in ISO 8601 format.
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
func (r BatchResponsePublicAssociationMultiWithLabel) RawJSON() string { return r.JSON.raw }
func (r *BatchResponsePublicAssociationMultiWithLabel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the batch processing request: "PENDING", "PROCESSING", "CANCELED",
// or "COMPLETE".
type BatchResponsePublicAssociationMultiWithLabelStatus string

const (
	BatchResponsePublicAssociationMultiWithLabelStatusCanceled   BatchResponsePublicAssociationMultiWithLabelStatus = "CANCELED"
	BatchResponsePublicAssociationMultiWithLabelStatusComplete   BatchResponsePublicAssociationMultiWithLabelStatus = "COMPLETE"
	BatchResponsePublicAssociationMultiWithLabelStatusPending    BatchResponsePublicAssociationMultiWithLabelStatus = "PENDING"
	BatchResponsePublicAssociationMultiWithLabelStatusProcessing BatchResponsePublicAssociationMultiWithLabelStatus = "PROCESSING"
)

// The response returned after performing a batch operation on associations.
type BatchResponsePublicDefaultAssociation struct {
	// The timestamp when the batch process was completed, in ISO 8601 format.
	CompletedAt time.Time                  `json:"completedAt" api:"required" format:"date-time"`
	Results     []PublicDefaultAssociation `json:"results" api:"required"`
	// The timestamp when the batch process began execution, in ISO 8601 format.
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// The status of the batch processing request. Can be: "PENDING", "PROCESSING",
	// "CANCELED", or "COMPLETE".
	//
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status BatchResponsePublicDefaultAssociationStatus `json:"status" api:"required"`
	Errors []shared.StandardError                      `json:"errors"`
	// An object containing relevant links related to the batch request.
	Links map[string]string `json:"links"`
	// The total number of errors that occurred during the operation.
	NumErrors int64 `json:"numErrors"`
	// The timestamp when the batch process was initiated, in ISO 8601 format.
	RequestedAt time.Time `json:"requestedAt" format:"date-time"`
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
func (r BatchResponsePublicDefaultAssociation) RawJSON() string { return r.JSON.raw }
func (r *BatchResponsePublicDefaultAssociation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the batch processing request. Can be: "PENDING", "PROCESSING",
// "CANCELED", or "COMPLETE".
type BatchResponsePublicDefaultAssociationStatus string

const (
	BatchResponsePublicDefaultAssociationStatusCanceled   BatchResponsePublicDefaultAssociationStatus = "CANCELED"
	BatchResponsePublicDefaultAssociationStatusComplete   BatchResponsePublicDefaultAssociationStatus = "COMPLETE"
	BatchResponsePublicDefaultAssociationStatusPending    BatchResponsePublicDefaultAssociationStatus = "PENDING"
	BatchResponsePublicDefaultAssociationStatusProcessing BatchResponsePublicDefaultAssociationStatus = "PROCESSING"
)

type CollectionResponseMultiAssociatedObjectWithLabelForwardPaging struct {
	Results []MultiAssociatedObjectWithLabel `json:"results" api:"required"`
	Paging  shared.ForwardPaging             `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseMultiAssociatedObjectWithLabelForwardPaging) RawJSON() string {
	return r.JSON.raw
}
func (r *CollectionResponseMultiAssociatedObjectWithLabelForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a list of simple objects returned from an API request, along with the
// total count of objects available.
type CollectionResponseWithTotalSimplePublicObject struct {
	Results []SimplePublicObject `json:"results" api:"required"`
	// The number of available results
	Total  int64         `json:"total" api:"required"`
	Paging shared.Paging `json:"paging"`
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
func (r CollectionResponseWithTotalSimplePublicObject) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseWithTotalSimplePublicObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DateTime struct {
	// Indicates whether the DateTime value represents only a date without a time
	// component.
	DateOnly bool `json:"dateOnly" api:"required"`
	// The integer value representing the shift in minutes from UTC for the DateTime
	// value.
	TimeZoneShift int64 `json:"timeZoneShift" api:"required"`
	// The integer value representing a specific point in time.
	Value int64 `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DateOnly      respjson.Field
		TimeZoneShift respjson.Field
		Value         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DateTime) RawJSON() string { return r.JSON.raw }
func (r *DateTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines a single condition for searching CRM objects, specifying the property to
// filter on, the operator to use (such as equals, greater than, or contains), and
// the value(s) to compare against.
//
// The properties Operator, PropertyName are required.
type FilterParam struct {
	// The comparison operator used in the filter, such as "EQ" or "GT".
	//
	// Any of "BETWEEN", "CONTAINS_TOKEN", "EQ", "GT", "GTE", "HAS_PROPERTY", "IN",
	// "LT", "LTE", "NEQ", "NOT_CONTAINS_TOKEN", "NOT_HAS_PROPERTY", "NOT_IN".
	Operator FilterOperator `json:"operator,omitzero" api:"required"`
	// The name of the property to apply the filter to.
	PropertyName string `json:"propertyName" api:"required"`
	// The upper boundary value when using ranged-based filters.
	HighValue param.Opt[string] `json:"highValue,omitzero"`
	// The value to match against the property.
	Value param.Opt[string] `json:"value,omitzero"`
	// The values to match against the property.
	Values []string `json:"values,omitzero"`
	paramObj
}

func (r FilterParam) MarshalJSON() (data []byte, err error) {
	type shadow FilterParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FilterParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The comparison operator used in the filter, such as "EQ" or "GT".
type FilterOperator string

const (
	FilterOperatorBetween          FilterOperator = "BETWEEN"
	FilterOperatorContainsToken    FilterOperator = "CONTAINS_TOKEN"
	FilterOperatorEq               FilterOperator = "EQ"
	FilterOperatorGt               FilterOperator = "GT"
	FilterOperatorGte              FilterOperator = "GTE"
	FilterOperatorHasProperty      FilterOperator = "HAS_PROPERTY"
	FilterOperatorIn               FilterOperator = "IN"
	FilterOperatorLt               FilterOperator = "LT"
	FilterOperatorLte              FilterOperator = "LTE"
	FilterOperatorNeq              FilterOperator = "NEQ"
	FilterOperatorNotContainsToken FilterOperator = "NOT_CONTAINS_TOKEN"
	FilterOperatorNotHasProperty   FilterOperator = "NOT_HAS_PROPERTY"
	FilterOperatorNotIn            FilterOperator = "NOT_IN"
)

// The property Filters is required.
type FilterGroupParam struct {
	Filters []FilterParam `json:"filters,omitzero" api:"required"`
	paramObj
}

func (r FilterGroupParam) MarshalJSON() (data []byte, err error) {
	type shadow FilterGroupParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FilterGroupParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The relationship descriptors applicable between two object types.
type LabelsBetweenObjectPair struct {
	// Source unique ID of the object.
	FromObjectID string `json:"fromObjectId" api:"required"`
	// Source object type.
	FromObjectTypeID string   `json:"fromObjectTypeId" api:"required"`
	Labels           []string `json:"labels" api:"required"`
	// Target unique ID of the object.
	ToObjectID string `json:"toObjectId" api:"required"`
	// Target object type.
	ToObjectTypeID string `json:"toObjectTypeId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FromObjectID     respjson.Field
		FromObjectTypeID respjson.Field
		Labels           respjson.Field
		ToObjectID       respjson.Field
		ToObjectTypeID   respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LabelsBetweenObjectPair) RawJSON() string { return r.JSON.raw }
func (r *LabelsBetweenObjectPair) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an object that is associated with multiple other objects, with
// optional context.
type MultiAssociatedObjectWithLabel struct {
	AssociationTypes []AssociationSpecWithLabel `json:"associationTypes" api:"required"`
	// Target unique ID of the object.
	ToObjectID string `json:"toObjectId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssociationTypes respjson.Field
		ToObjectID       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MultiAssociatedObjectWithLabel) RawJSON() string { return r.JSON.raw }
func (r *MultiAssociatedObjectWithLabel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties From, To are required.
type PublicAssociationMultiArchiveParam struct {
	// Contains the Id of a Public Object
	From shared.PublicObjectIDParam   `json:"from,omitzero" api:"required"`
	To   []shared.PublicObjectIDParam `json:"to,omitzero" api:"required"`
	paramObj
}

func (r PublicAssociationMultiArchiveParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicAssociationMultiArchiveParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicAssociationMultiArchiveParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties From, To, Types are required.
type PublicAssociationMultiPostParam struct {
	// Contains the Id of a Public Object
	From shared.PublicObjectIDParam `json:"from,omitzero" api:"required"`
	// Contains the Id of a Public Object
	To    shared.PublicObjectIDParam    `json:"to,omitzero" api:"required"`
	Types []shared.AssociationSpecParam `json:"types,omitzero" api:"required"`
	paramObj
}

func (r PublicAssociationMultiPostParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicAssociationMultiPostParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicAssociationMultiPostParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicAssociationMultiWithLabel struct {
	// Contains the Id of a Public Object
	From   shared.PublicObjectID            `json:"from" api:"required"`
	To     []MultiAssociatedObjectWithLabel `json:"to" api:"required"`
	Paging shared.Paging                    `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From        respjson.Field
		To          respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicAssociationMultiWithLabel) RawJSON() string { return r.JSON.raw }
func (r *PublicAssociationMultiWithLabel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicDefaultAssociation struct {
	// Defines the type, direction, and details of the relationship between two CRM
	// objects.
	AssociationSpec shared.AssociationSpec `json:"associationSpec" api:"required"`
	// Contains the Id of a Public Object
	From shared.PublicObjectID `json:"from" api:"required"`
	// Contains the Id of a Public Object
	To shared.PublicObjectID `json:"to" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssociationSpec respjson.Field
		From            respjson.Field
		To              respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicDefaultAssociation) RawJSON() string { return r.JSON.raw }
func (r *PublicDefaultAssociation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties From, To are required.
type PublicDefaultAssociationMultiPostParam struct {
	// Contains the Id of a Public Object
	From shared.PublicObjectIDParam `json:"from,omitzero" api:"required"`
	// Contains the Id of a Public Object
	To shared.PublicObjectIDParam `json:"to,omitzero" api:"required"`
	paramObj
}

func (r PublicDefaultAssociationMultiPostParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicDefaultAssociationMultiPostParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicDefaultAssociationMultiPostParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type PublicFetchAssociationsBatchRequestParam struct {
	// The unique identifier for the object whose associations are being fetched.
	ID string `json:"id" api:"required"`
	// A paging cursor token used to retrieve the next set of results in a paginated
	// response.
	After param.Opt[string] `json:"after,omitzero"`
	paramObj
}

func (r PublicFetchAssociationsBatchRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicFetchAssociationsBatchRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicFetchAssociationsBatchRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes a search request
//
// The properties After, FilterGroups, Limit, Properties, Sorts are required.
type PublicObjectSearchRequestParam struct {
	// A paging cursor token for retrieving subsequent pages.
	After string `json:"after" api:"required"`
	// Up to 6 groups of filters defining additional query criteria.
	FilterGroups []FilterGroupParam `json:"filterGroups,omitzero" api:"required"`
	// The maximum results to return, up to 200 objects.
	Limit int64 `json:"limit" api:"required"`
	// A list of property names to include in the response.
	Properties []string `json:"properties,omitzero" api:"required"`
	// Specifies sorting order based on object properties.
	Sorts []string `json:"sorts,omitzero" api:"required"`
	// The search query string, up to 3000 characters.
	Query param.Opt[string] `json:"query,omitzero"`
	paramObj
}

func (r PublicObjectSearchRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicObjectSearchRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicObjectSearchRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReportCreationResponse struct {
	EnqueueTime DateTime `json:"enqueueTime" api:"required"`
	// Email of the user
	UserEmail string `json:"userEmail" api:"required"`
	// ID of the user
	UserID int64 `json:"userId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnqueueTime respjson.Field
		UserEmail   respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReportCreationResponse) RawJSON() string { return r.JSON.raw }
func (r *ReportCreationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A simple public object.
type SimplePublicObject struct {
	// The unique ID of the object.
	ID string `json:"id" api:"required"`
	// Whether the object is archived.
	Archived bool `json:"archived" api:"required"`
	// The timestamp when the object was created, in ISO 8601 format.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Key-value pairs representing the properties of the object.
	Properties map[string]string `json:"properties" api:"required"`
	// The timestamp when the object was last updated, in ISO 8601 format.
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// The timestamp when the object was archived, in ISO 8601 format.
	ArchivedAt time.Time `json:"archivedAt" format:"date-time"`
	// An identifier used for tracing the write request for the object.
	ObjectWriteTraceID string `json:"objectWriteTraceId"`
	// Key-value pairs representing the properties of the object along with their
	// history.
	PropertiesWithHistory map[string][]ValueWithTimestamp `json:"propertiesWithHistory"`
	// The URL associated with the object.
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		Archived              respjson.Field
		CreatedAt             respjson.Field
		Properties            respjson.Field
		UpdatedAt             respjson.Field
		ArchivedAt            respjson.Field
		ObjectWriteTraceID    respjson.Field
		PropertiesWithHistory respjson.Field
		URL                   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimplePublicObject) RawJSON() string { return r.JSON.raw }
func (r *SimplePublicObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Property model that includes timestamp.
type ValueWithTimestamp struct {
	// The property type.
	SourceType string `json:"sourceType" api:"required"`
	// The timestamp when the property was updated, in ISO 8601 format.
	Timestamp time.Time `json:"timestamp" api:"required" format:"date-time"`
	// The property value.
	Value string `json:"value" api:"required"`
	// The unique ID of the property.
	SourceID string `json:"sourceId"`
	// A human-readable label.
	SourceLabel string `json:"sourceLabel"`
	// The ID of the user who last updated the property.
	UpdatedByUserID int64 `json:"updatedByUserId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SourceType      respjson.Field
		Timestamp       respjson.Field
		Value           respjson.Field
		SourceID        respjson.Field
		SourceLabel     respjson.Field
		UpdatedByUserID respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ValueWithTimestamp) RawJSON() string { return r.JSON.raw }
func (r *ValueWithTimestamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
