// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/marketing"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// CRMService contains methods and other services that help with interacting with
// the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCRMService] method instead.
type CRMService struct {
	Options             []option.RequestOption
	AppUninstalls       AppUninstallService
	Associations        AssociationService
	Exports             ExportService
	Extensions          ExtensionService
	FeatureFlags        FeatureFlagService
	Imports             ImportService
	Limits              LimitService
	Lists               ListService
	ObjectLibrary       ObjectLibraryService
	Objects             ObjectService
	Owners              OwnerService
	Pipelines           PipelineService
	Properties          PropertyService
	PropertyValidations PropertyValidationService
	Timeline            TimelineService
	Users               UserService
}

// NewCRMService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCRMService(opts ...option.RequestOption) (r CRMService) {
	r = CRMService{}
	r.Options = opts
	r.AppUninstalls = NewAppUninstallService(opts...)
	r.Associations = NewAssociationService(opts...)
	r.Exports = NewExportService(opts...)
	r.Extensions = NewExtensionService(opts...)
	r.FeatureFlags = NewFeatureFlagService(opts...)
	r.Imports = NewImportService(opts...)
	r.Limits = NewLimitService(opts...)
	r.Lists = NewListService(opts...)
	r.ObjectLibrary = NewObjectLibraryService(opts...)
	r.Objects = NewObjectService(opts...)
	r.Owners = NewOwnerService(opts...)
	r.Pipelines = NewPipelineService(opts...)
	r.Properties = NewPropertyService(opts...)
	r.PropertyValidations = NewPropertyValidationService(opts...)
	r.Timeline = NewTimelineService(opts...)
	r.Users = NewUserService(opts...)
	return
}

// Contains the id and type of an association
type AssociatedID struct {
	// The ID for the association type.
	ID string `json:"id,required"`
	// The type of association.
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AssociatedID) RawJSON() string { return r.JSON.raw }
func (r *AssociatedID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type BatchInputSimplePublicObjectBatchInputParam struct {
	Inputs []SimplePublicObjectBatchInputParam `json:"inputs,omitzero,required"`
	paramObj
}

func (r BatchInputSimplePublicObjectBatchInputParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputSimplePublicObjectBatchInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputSimplePublicObjectBatchInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type BatchInputSimplePublicObjectBatchInputForCreateParam struct {
	Inputs []SimplePublicObjectBatchInputForCreateParam `json:"inputs,omitzero,required"`
	paramObj
}

func (r BatchInputSimplePublicObjectBatchInputForCreateParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputSimplePublicObjectBatchInputForCreateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputSimplePublicObjectBatchInputForCreateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type BatchInputSimplePublicObjectBatchInputUpsertParam struct {
	Inputs []SimplePublicObjectBatchInputUpsertParam `json:"inputs,omitzero,required"`
	paramObj
}

func (r BatchInputSimplePublicObjectBatchInputUpsertParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputSimplePublicObjectBatchInputUpsertParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputSimplePublicObjectBatchInputUpsertParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type BatchInputSimplePublicObjectIDParam struct {
	Inputs []SimplePublicObjectIDParam `json:"inputs,omitzero,required"`
	paramObj
}

func (r BatchInputSimplePublicObjectIDParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchInputSimplePublicObjectIDParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchInputSimplePublicObjectIDParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the input for reading a batch of CRM objects, including arrays of
// object IDs, requested property names (with optional history), and an optional
// unique identifying property.
//
// The properties Inputs, Properties, PropertiesWithHistory are required.
type BatchReadInputSimplePublicObjectIDParam struct {
	Inputs []SimplePublicObjectIDParam `json:"inputs,omitzero,required"`
	// Key-value pairs for setting properties for the new object.
	Properties []string `json:"properties,omitzero,required"`
	// Key-value pairs for setting properties for the new object and their histories.
	PropertiesWithHistory []string `json:"propertiesWithHistory,omitzero,required"`
	// When using a custom unique value property to retrieve records, the name of the
	// property. Do not include this parameter if retrieving by record ID.
	IDProperty param.Opt[string] `json:"idProperty,omitzero"`
	paramObj
}

func (r BatchReadInputSimplePublicObjectIDParam) MarshalJSON() (data []byte, err error) {
	type shadow BatchReadInputSimplePublicObjectIDParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchReadInputSimplePublicObjectIDParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchResponsePublicDefaultAssociation struct {
	CompletedAt time.Time                  `json:"completedAt,required" format:"date-time"`
	Results     []PublicDefaultAssociation `json:"results,required"`
	StartedAt   time.Time                  `json:"startedAt,required" format:"date-time"`
	// Any of "PENDING", "PROCESSING", "CANCELED", "COMPLETE".
	Status      BatchResponsePublicDefaultAssociationStatus `json:"status,required"`
	Errors      []StandardError1                            `json:"errors"`
	Links       map[string]string                           `json:"links"`
	NumErrors   int64                                       `json:"numErrors"`
	RequestedAt time.Time                                   `json:"requestedAt" format:"date-time"`
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

type BatchResponsePublicDefaultAssociationStatus string

const (
	BatchResponsePublicDefaultAssociationStatusPending    BatchResponsePublicDefaultAssociationStatus = "PENDING"
	BatchResponsePublicDefaultAssociationStatusProcessing BatchResponsePublicDefaultAssociationStatus = "PROCESSING"
	BatchResponsePublicDefaultAssociationStatusCanceled   BatchResponsePublicDefaultAssociationStatus = "CANCELED"
	BatchResponsePublicDefaultAssociationStatusComplete   BatchResponsePublicDefaultAssociationStatus = "COMPLETE"
)

// A public object batch response object
type BatchResponseSimplePublicObject struct {
	// The timestamp when the batch processing was completed, in ISO 8601 format.
	CompletedAt time.Time            `json:"completedAt,required" format:"date-time"`
	Results     []SimplePublicObject `json:"results,required"`
	// The timestamp when the batch processing began, in ISO 8601 format.
	StartedAt time.Time `json:"startedAt,required" format:"date-time"`
	// The status of the batch processing request: "PENDING", "PROCESSING",
	// "CANCELLED", or "COMPLETE"
	//
	// Any of "PENDING", "PROCESSING", "CANCELED", "COMPLETE".
	Status BatchResponseSimplePublicObjectStatus `json:"status,required"`
	Errors []shared.StandardError                `json:"errors"`
	// An object containing relevant links related to the batch request.
	Links     map[string]string `json:"links"`
	NumErrors int64             `json:"numErrors"`
	// The timestamp when the batch request was initially made, in ISO 8601 format.
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
func (r BatchResponseSimplePublicObject) RawJSON() string { return r.JSON.raw }
func (r *BatchResponseSimplePublicObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the batch processing request: "PENDING", "PROCESSING",
// "CANCELLED", or "COMPLETE"
type BatchResponseSimplePublicObjectStatus string

const (
	BatchResponseSimplePublicObjectStatusPending    BatchResponseSimplePublicObjectStatus = "PENDING"
	BatchResponseSimplePublicObjectStatusProcessing BatchResponseSimplePublicObjectStatus = "PROCESSING"
	BatchResponseSimplePublicObjectStatusCanceled   BatchResponseSimplePublicObjectStatus = "CANCELED"
	BatchResponseSimplePublicObjectStatusComplete   BatchResponseSimplePublicObjectStatus = "COMPLETE"
)

// Represents the result of a batch upsert operation, including the operation’s
// status, timestamps, and a list of successfully created or updated objects.
type BatchResponseSimplePublicUpsertObject struct {
	// The timestamp when the batch process was completed, in ISO 8601 format.
	CompletedAt time.Time                  `json:"completedAt,required" format:"date-time"`
	Results     []SimplePublicUpsertObject `json:"results,required"`
	// The timestamp when the batch process began execution, in ISO 8601 format.
	StartedAt time.Time `json:"startedAt,required" format:"date-time"`
	// The status of the batch processing request. Can be: "PENDING", "PROCESSING",
	// "CANCELED", or "COMPLETE".
	//
	// Any of "PENDING", "PROCESSING", "CANCELED", "COMPLETE".
	Status BatchResponseSimplePublicUpsertObjectStatus `json:"status,required"`
	Errors []shared.StandardError                      `json:"errors"`
	// An object containing relevant links related to the batch request.
	Links     map[string]string `json:"links"`
	NumErrors int64             `json:"numErrors"`
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
func (r BatchResponseSimplePublicUpsertObject) RawJSON() string { return r.JSON.raw }
func (r *BatchResponseSimplePublicUpsertObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the batch processing request. Can be: "PENDING", "PROCESSING",
// "CANCELED", or "COMPLETE".
type BatchResponseSimplePublicUpsertObjectStatus string

const (
	BatchResponseSimplePublicUpsertObjectStatusPending    BatchResponseSimplePublicUpsertObjectStatus = "PENDING"
	BatchResponseSimplePublicUpsertObjectStatusProcessing BatchResponseSimplePublicUpsertObjectStatus = "PROCESSING"
	BatchResponseSimplePublicUpsertObjectStatusCanceled   BatchResponseSimplePublicUpsertObjectStatus = "CANCELED"
	BatchResponseSimplePublicUpsertObjectStatusComplete   BatchResponseSimplePublicUpsertObjectStatus = "COMPLETE"
)

type CollectionResponseAssociatedID struct {
	Results []AssociatedID `json:"results,required"`
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
func (r CollectionResponseAssociatedID) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseAssociatedID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponseMultiAssociatedObjectWithLabel struct {
	Results []MultiAssociatedObjectWithLabel `json:"results,required"`
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
func (r CollectionResponseMultiAssociatedObjectWithLabel) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseMultiAssociatedObjectWithLabel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponseSimplePublicObjectWithAssociations struct {
	Results []SimplePublicObjectWithAssociations `json:"results,required"`
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
func (r CollectionResponseSimplePublicObjectWithAssociations) RawJSON() string { return r.JSON.raw }
func (r *CollectionResponseSimplePublicObjectWithAssociations) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionResponseWithTotalSimplePublicObject struct {
	Results []SimplePublicObject `json:"results,required"`
	// The number of available results
	Total int64 `json:"total,required"`
	// Contains information pagination of results.
	Paging marketing.Paging `json:"paging"`
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

type CreatedResponseLabelsBetweenObjectPair struct {
	CreatedResourceID string                  `json:"createdResourceId,required"`
	Entity            LabelsBetweenObjectPair `json:"entity,required"`
	Location          string                  `json:"location"`
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
func (r CreatedResponseLabelsBetweenObjectPair) RawJSON() string { return r.JSON.raw }
func (r *CreatedResponseLabelsBetweenObjectPair) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreatedResponseSimplePublicObject struct {
	CreatedResourceID string `json:"createdResourceId,required"`
	// A simple public object.
	Entity   SimplePublicObject `json:"entity,required"`
	Location string             `json:"location"`
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
func (r CreatedResponseSimplePublicObject) RawJSON() string { return r.JSON.raw }
func (r *CreatedResponseSimplePublicObject) UnmarshalJSON(data []byte) error {
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
	// Any of "EQ", "NEQ", "LT", "LTE", "GT", "GTE", "BETWEEN", "IN", "NOT_IN",
	// "HAS_PROPERTY", "NOT_HAS_PROPERTY".
	Operator FilterOperator `json:"operator,omitzero,required"`
	// The name of the property to apply the filter to.
	PropertyName string `json:"propertyName,required"`
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
	FilterOperatorEq             FilterOperator = "EQ"
	FilterOperatorNeq            FilterOperator = "NEQ"
	FilterOperatorLt             FilterOperator = "LT"
	FilterOperatorLte            FilterOperator = "LTE"
	FilterOperatorGt             FilterOperator = "GT"
	FilterOperatorGte            FilterOperator = "GTE"
	FilterOperatorBetween        FilterOperator = "BETWEEN"
	FilterOperatorIn             FilterOperator = "IN"
	FilterOperatorNotIn          FilterOperator = "NOT_IN"
	FilterOperatorHasProperty    FilterOperator = "HAS_PROPERTY"
	FilterOperatorNotHasProperty FilterOperator = "NOT_HAS_PROPERTY"
)

// The property Filters is required.
type FilterGroupParam struct {
	Filters []FilterParam `json:"filters,omitzero,required"`
	paramObj
}

func (r FilterGroupParam) MarshalJSON() (data []byte, err error) {
	type shadow FilterGroupParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FilterGroupParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LabelsBetweenObjectPair struct {
	FromObjectID     string   `json:"fromObjectId,required"`
	FromObjectTypeID string   `json:"fromObjectTypeId,required"`
	Labels           []string `json:"labels,required"`
	ToObjectID       string   `json:"toObjectId,required"`
	ToObjectTypeID   string   `json:"toObjectTypeId,required"`
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

type MultiAssociatedObjectWithLabel struct {
	AssociationTypes []AssociationSpecWithLabel1 `json:"associationTypes,required"`
	ToObjectID       string                      `json:"toObjectId,required"`
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

// The properties To, Types are required.
type PublicAssociationsForObjectParam struct {
	To    shared.PublicObjectIDParam    `json:"to,omitzero,required"`
	Types []shared.AssociationSpecParam `json:"types,omitzero,required"`
	paramObj
}

func (r PublicAssociationsForObjectParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicAssociationsForObjectParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicAssociationsForObjectParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicDefaultAssociation struct {
	// Defines the type, direction, and details of the relationship between two CRM
	// objects.
	AssociationSpec AssociationSpec1      `json:"associationSpec,required"`
	From            shared.PublicObjectID `json:"from,required"`
	To              shared.PublicObjectID `json:"to,required"`
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

// The property ObjectID is required.
type PublicGdprDeleteInputParam struct {
	ObjectID string `json:"objectId,required"`
	// The name of a property whose values are unique for this object
	IDProperty param.Opt[string] `json:"idProperty,omitzero"`
	paramObj
}

func (r PublicGdprDeleteInputParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicGdprDeleteInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicGdprDeleteInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ObjectIDToMerge, PrimaryObjectID are required.
type PublicMergeInputParam struct {
	ObjectIDToMerge string `json:"objectIdToMerge,required"`
	PrimaryObjectID string `json:"primaryObjectId,required"`
	paramObj
}

func (r PublicMergeInputParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicMergeInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicMergeInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Describes a search request
type PublicObjectSearchRequestParam struct {
	// A paging cursor token for retrieving subsequent pages.
	After param.Opt[string] `json:"after,omitzero"`
	// The maximum results to return, up to 200 objects.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// The search query string, up to 3000 characters.
	Query param.Opt[string] `json:"query,omitzero"`
	// Up to 6 groups of filters defining additional query criteria.
	FilterGroups []FilterGroupParam `json:"filterGroups,omitzero"`
	// A list of property names to include in the response.
	Properties []string `json:"properties,omitzero"`
	// Specifies sorting order based on object properties.
	Sorts []string `json:"sorts,omitzero"`
	paramObj
}

func (r PublicObjectSearchRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicObjectSearchRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicObjectSearchRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A simple public object.
type SimplePublicObject struct {
	// The unique ID of the object.
	ID string `json:"id,required"`
	// The timestamp when the object was created, in ISO 8601 format.
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Key-value pairs representing the properties of the object.
	Properties map[string]string `json:"properties,required"`
	// The timestamp when the object was last updated, in ISO 8601 format.
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	// Whether the object is archived.
	Archived bool `json:"archived"`
	// The timestamp when the object was archived, in ISO 8601 format.
	ArchivedAt         time.Time `json:"archivedAt" format:"date-time"`
	ObjectWriteTraceID string    `json:"objectWriteTraceId"`
	// Key-value pairs representing the properties of the object along with their
	// history.
	PropertiesWithHistory map[string][]ValueWithTimestamp `json:"propertiesWithHistory"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		CreatedAt             respjson.Field
		Properties            respjson.Field
		UpdatedAt             respjson.Field
		Archived              respjson.Field
		ArchivedAt            respjson.Field
		ObjectWriteTraceID    respjson.Field
		PropertiesWithHistory respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimplePublicObject) RawJSON() string { return r.JSON.raw }
func (r *SimplePublicObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contains an array of CRM object records to be processed in a batch operation,
// each defined by their ID and properties.
//
// The properties ID, Properties are required.
type SimplePublicObjectBatchInputParam struct {
	// The id to be updated. This can be the object id, or the unique property value of
	// the idProperty property
	ID string `json:"id,required"`
	// Key-value pairs representing the properties of the object.
	Properties map[string]string `json:"properties,omitzero,required"`
	// The name of a property whose values are unique for this object
	IDProperty param.Opt[string] `json:"idProperty,omitzero"`
	// A unique identifier for tracing the request.
	ObjectWriteTraceID param.Opt[string] `json:"objectWriteTraceId,omitzero"`
	paramObj
}

func (r SimplePublicObjectBatchInputParam) MarshalJSON() (data []byte, err error) {
	type shadow SimplePublicObjectBatchInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SimplePublicObjectBatchInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Properties is required.
type SimplePublicObjectBatchInputForCreateParam struct {
	Properties         map[string]string                  `json:"properties,omitzero,required"`
	ObjectWriteTraceID param.Opt[string]                  `json:"objectWriteTraceId,omitzero"`
	Associations       []PublicAssociationsForObjectParam `json:"associations,omitzero"`
	paramObj
}

func (r SimplePublicObjectBatchInputForCreateParam) MarshalJSON() (data []byte, err error) {
	type shadow SimplePublicObjectBatchInputForCreateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SimplePublicObjectBatchInputForCreateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an object used in batch upsert operations, containing an object’s
// unique identifier, its properties, and optionally the unique property name and a
// write trace ID.
//
// The properties ID, Properties are required.
type SimplePublicObjectBatchInputUpsertParam struct {
	// The unique ID of the object.
	ID string `json:"id,required"`
	// Key value pairs representing the properties of the object.
	Properties map[string]string `json:"properties,omitzero,required"`
	// The name of a property whose values are unique for this object
	IDProperty param.Opt[string] `json:"idProperty,omitzero"`
	// An identifier for tracing the creation request.
	ObjectWriteTraceID param.Opt[string] `json:"objectWriteTraceId,omitzero"`
	paramObj
}

func (r SimplePublicObjectBatchInputUpsertParam) MarshalJSON() (data []byte, err error) {
	type shadow SimplePublicObjectBatchInputUpsertParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SimplePublicObjectBatchInputUpsertParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ID is required.
type SimplePublicObjectIDParam struct {
	ID string `json:"id,required"`
	paramObj
}

func (r SimplePublicObjectIDParam) MarshalJSON() (data []byte, err error) {
	type shadow SimplePublicObjectIDParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SimplePublicObjectIDParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents the input required to create or update a CRM object, containing an
// object with property names and their corresponding values.
//
// The property Properties is required.
type SimplePublicObjectInputParam struct {
	// Key value pairs representing the properties of the object.
	Properties map[string]string `json:"properties,omitzero,required"`
	paramObj
}

func (r SimplePublicObjectInputParam) MarshalJSON() (data []byte, err error) {
	type shadow SimplePublicObjectInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SimplePublicObjectInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Is the input object used to create a new CRM object, containing the properties
// to be set and optional associations to link the new record with other CRM
// objects.
//
// The property Properties is required.
type SimplePublicObjectInputForCreateParam struct {
	// Key-value pairs for setting properties for the new object.
	Properties   map[string]string                  `json:"properties,omitzero,required"`
	Associations []PublicAssociationsForObjectParam `json:"associations,omitzero"`
	paramObj
}

func (r SimplePublicObjectInputForCreateParam) MarshalJSON() (data []byte, err error) {
	type shadow SimplePublicObjectInputForCreateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SimplePublicObjectInputForCreateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a CRM object along with its properties, timestamps, and a set of
// associated object IDs grouped by association type.
type SimplePublicObjectWithAssociations struct {
	// The unique ID of the object.
	ID string `json:"id,required"`
	// The timestamp when the object was created, in ISO 8601 format.
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Key value pairs representing the properties of the object.
	Properties map[string]string `json:"properties,required"`
	// The timestamp when the object was last updated, in ISO 8601 format.
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	// Whether the object is archived.
	Archived bool `json:"archived"`
	// The timestamp when the object was archived, in ISO 8601 format.
	ArchivedAt time.Time `json:"archivedAt" format:"date-time"`
	// A list defining relationships with other objects.
	Associations       map[string]CollectionResponseAssociatedID `json:"associations"`
	ObjectWriteTraceID string                                    `json:"objectWriteTraceId"`
	// Key-value pairs representing the properties of the object along with their
	// history.
	PropertiesWithHistory map[string][]ValueWithTimestamp `json:"propertiesWithHistory"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		CreatedAt             respjson.Field
		Properties            respjson.Field
		UpdatedAt             respjson.Field
		Archived              respjson.Field
		ArchivedAt            respjson.Field
		Associations          respjson.Field
		ObjectWriteTraceID    respjson.Field
		PropertiesWithHistory respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimplePublicObjectWithAssociations) RawJSON() string { return r.JSON.raw }
func (r *SimplePublicObjectWithAssociations) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a CRM object that has either been created or updated (upserted)
type SimplePublicUpsertObject struct {
	// The unique ID of the object.
	ID string `json:"id,required"`
	// The timestamp when the object was created, in ISO 8601 format.
	CreatedAt time.Time `json:"createdAt,required" format:"date-time"`
	// Whether the property is new.
	New bool `json:"new,required"`
	// Key value pairs representing the properties of the object.
	Properties map[string]string `json:"properties,required"`
	// The timestamp when the object was last updated, in ISO 8601 format.
	UpdatedAt time.Time `json:"updatedAt,required" format:"date-time"`
	// Whether the object is archived.
	Archived bool `json:"archived"`
	// The timestamp when the object was archived, in ISO 8601 format.
	ArchivedAt         time.Time `json:"archivedAt" format:"date-time"`
	ObjectWriteTraceID string    `json:"objectWriteTraceId"`
	// Key-value pairs representing the properties of the object along with their
	// history.
	PropertiesWithHistory map[string][]ValueWithTimestamp `json:"propertiesWithHistory"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		CreatedAt             respjson.Field
		New                   respjson.Field
		Properties            respjson.Field
		UpdatedAt             respjson.Field
		Archived              respjson.Field
		ArchivedAt            respjson.Field
		ObjectWriteTraceID    respjson.Field
		PropertiesWithHistory respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimplePublicUpsertObject) RawJSON() string { return r.JSON.raw }
func (r *SimplePublicUpsertObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Property model that includes timestamp.
type ValueWithTimestamp struct {
	// The property type.
	SourceType string `json:"sourceType,required"`
	// The timestamp when the property was updated, in ISO 8601 format.
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// The property value.
	Value string `json:"value,required"`
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
