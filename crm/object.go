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

// ObjectService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectService] method instead.
type ObjectService struct {
	options             []option.RequestOption
	Calls               ObjectCallService
	Carts               ObjectCartService
	CommercePayments    ObjectCommercePaymentService
	Communications      ObjectCommunicationService
	Companies           ObjectCompanyService
	Contacts            ObjectContactService
	Contracts           ObjectContractService
	Courses             ObjectCourseService
	Custom              ObjectCustomService
	Deals               ObjectDealService
	Discounts           ObjectDiscountService
	Emails              ObjectEmailService
	FeedbackSubmissions ObjectFeedbackSubmissionService
	Fees                ObjectFeeService
	GenericObjects      ObjectGenericObjectService
	GoalTargets         ObjectGoalTargetService
	Invoices            ObjectInvoiceService
	Leads               ObjectLeadService
	LineItems           ObjectLineItemService
	Listings            ObjectListingService
	Meetings            ObjectMeetingService
	Notes               ObjectNoteService
	Orders              ObjectOrderService
	PartnerClients      ObjectPartnerClientService
	PartnerServices     ObjectPartnerServiceService
	PostalMail          ObjectPostalMailService
	Products            ObjectProductService
	Projects            ObjectProjectService
	Quotes              ObjectQuoteService
	Services            ObjectServiceService
	Subscriptions       ObjectSubscriptionService
	Tasks               ObjectTaskService
	Taxes               ObjectTaxService
	Tickets             ObjectTicketService
	Users               ObjectUserService
}

// NewObjectService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewObjectService(opts ...option.RequestOption) (r ObjectService) {
	r = ObjectService{}
	r.options = opts
	r.Calls = NewObjectCallService(opts...)
	r.Carts = NewObjectCartService(opts...)
	r.CommercePayments = NewObjectCommercePaymentService(opts...)
	r.Communications = NewObjectCommunicationService(opts...)
	r.Companies = NewObjectCompanyService(opts...)
	r.Contacts = NewObjectContactService(opts...)
	r.Contracts = NewObjectContractService(opts...)
	r.Courses = NewObjectCourseService(opts...)
	r.Custom = NewObjectCustomService(opts...)
	r.Deals = NewObjectDealService(opts...)
	r.Discounts = NewObjectDiscountService(opts...)
	r.Emails = NewObjectEmailService(opts...)
	r.FeedbackSubmissions = NewObjectFeedbackSubmissionService(opts...)
	r.Fees = NewObjectFeeService(opts...)
	r.GenericObjects = NewObjectGenericObjectService(opts...)
	r.GoalTargets = NewObjectGoalTargetService(opts...)
	r.Invoices = NewObjectInvoiceService(opts...)
	r.Leads = NewObjectLeadService(opts...)
	r.LineItems = NewObjectLineItemService(opts...)
	r.Listings = NewObjectListingService(opts...)
	r.Meetings = NewObjectMeetingService(opts...)
	r.Notes = NewObjectNoteService(opts...)
	r.Orders = NewObjectOrderService(opts...)
	r.PartnerClients = NewObjectPartnerClientService(opts...)
	r.PartnerServices = NewObjectPartnerServiceService(opts...)
	r.PostalMail = NewObjectPostalMailService(opts...)
	r.Products = NewObjectProductService(opts...)
	r.Projects = NewObjectProjectService(opts...)
	r.Quotes = NewObjectQuoteService(opts...)
	r.Services = NewObjectServiceService(opts...)
	r.Subscriptions = NewObjectSubscriptionService(opts...)
	r.Tasks = NewObjectTaskService(opts...)
	r.Taxes = NewObjectTaxService(opts...)
	r.Tickets = NewObjectTicketService(opts...)
	r.Users = NewObjectUserService(opts...)
	return
}

// Contains the id and type of an association
type AssociatedID struct {
	// The ID for the association type.
	ID string `json:"id" api:"required"`
	// The type of associations.
	Type string `json:"type" api:"required"`
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
	Inputs []SimplePublicObjectBatchInputParam `json:"inputs,omitzero" api:"required"`
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
	Inputs []SimplePublicObjectBatchInputForCreateParam `json:"inputs,omitzero" api:"required"`
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
	Inputs []SimplePublicObjectBatchInputUpsertParam `json:"inputs,omitzero" api:"required"`
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
	Inputs []SimplePublicObjectIDParam `json:"inputs,omitzero" api:"required"`
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
	Inputs []SimplePublicObjectIDParam `json:"inputs,omitzero" api:"required"`
	// Key-value pairs for setting properties for the new object.
	Properties []string `json:"properties,omitzero" api:"required"`
	// Key-value pairs for setting properties for the new object and their histories.
	PropertiesWithHistory []string `json:"propertiesWithHistory,omitzero" api:"required"`
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

// A public object batch response object
type BatchResponseSimplePublicObject struct {
	// The timestamp when the batch processing was completed, in ISO 8601 format.
	CompletedAt time.Time            `json:"completedAt" api:"required" format:"date-time"`
	Results     []SimplePublicObject `json:"results" api:"required"`
	// The timestamp when the batch processing began, in ISO 8601 format.
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// The status of the batch processing request: "PENDING", "PROCESSING",
	// "CANCELLED", or "COMPLETE"
	//
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status BatchResponseSimplePublicObjectStatus `json:"status" api:"required"`
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
func (r BatchResponseSimplePublicObject) RawJSON() string { return r.JSON.raw }
func (r *BatchResponseSimplePublicObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the batch processing request: "PENDING", "PROCESSING",
// "CANCELLED", or "COMPLETE"
type BatchResponseSimplePublicObjectStatus string

const (
	BatchResponseSimplePublicObjectStatusCanceled   BatchResponseSimplePublicObjectStatus = "CANCELED"
	BatchResponseSimplePublicObjectStatusComplete   BatchResponseSimplePublicObjectStatus = "COMPLETE"
	BatchResponseSimplePublicObjectStatusPending    BatchResponseSimplePublicObjectStatus = "PENDING"
	BatchResponseSimplePublicObjectStatusProcessing BatchResponseSimplePublicObjectStatus = "PROCESSING"
)

// Represents the result of a batch upsert operation, including the operation’s
// status, timestamps, and a list of successfully created or updated objects.
type BatchResponseSimplePublicUpsertObject struct {
	// The timestamp when the batch process was completed, in ISO 8601 format.
	CompletedAt time.Time                  `json:"completedAt" api:"required" format:"date-time"`
	Results     []SimplePublicUpsertObject `json:"results" api:"required"`
	// The timestamp when the batch process began execution, in ISO 8601 format.
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// The status of the batch processing request. Can be: "PENDING", "PROCESSING",
	// "CANCELED", or "COMPLETE".
	//
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status BatchResponseSimplePublicUpsertObjectStatus `json:"status" api:"required"`
	// An object containing relevant links related to the batch request.
	Links map[string]string `json:"links"`
	// The timestamp when the batch process was initiated, in ISO 8601 format.
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
func (r BatchResponseSimplePublicUpsertObject) RawJSON() string { return r.JSON.raw }
func (r *BatchResponseSimplePublicUpsertObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the batch processing request. Can be: "PENDING", "PROCESSING",
// "CANCELED", or "COMPLETE".
type BatchResponseSimplePublicUpsertObjectStatus string

const (
	BatchResponseSimplePublicUpsertObjectStatusCanceled   BatchResponseSimplePublicUpsertObjectStatus = "CANCELED"
	BatchResponseSimplePublicUpsertObjectStatusComplete   BatchResponseSimplePublicUpsertObjectStatus = "COMPLETE"
	BatchResponseSimplePublicUpsertObjectStatusPending    BatchResponseSimplePublicUpsertObjectStatus = "PENDING"
	BatchResponseSimplePublicUpsertObjectStatusProcessing BatchResponseSimplePublicUpsertObjectStatus = "PROCESSING"
)

type CollectionResponseAssociatedID struct {
	Results []AssociatedID `json:"results" api:"required"`
	Paging  shared.Paging  `json:"paging"`
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

type CollectionResponseSimplePublicObjectWithAssociationsForwardPaging struct {
	Results []SimplePublicObjectWithAssociations `json:"results" api:"required"`
	Paging  shared.ForwardPaging                 `json:"paging"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Paging      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CollectionResponseSimplePublicObjectWithAssociationsForwardPaging) RawJSON() string {
	return r.JSON.raw
}
func (r *CollectionResponseSimplePublicObjectWithAssociationsForwardPaging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties To, Types are required.
type PublicAssociationsForObjectParam struct {
	// Contains the Id of a Public Object
	To    shared.PublicObjectIDParam    `json:"to,omitzero" api:"required"`
	Types []shared.AssociationSpecParam `json:"types,omitzero" api:"required"`
	paramObj
}

func (r PublicAssociationsForObjectParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicAssociationsForObjectParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicAssociationsForObjectParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Input data for merging two records.
//
// The properties ObjectIDToMerge, PrimaryObjectID are required.
type PublicMergeInputParam struct {
	// The ID of the company to merge into the primary.
	ObjectIDToMerge string `json:"objectIdToMerge" api:"required"`
	// The ID of the primary company, which the other will merge into.
	PrimaryObjectID string `json:"primaryObjectId" api:"required"`
	paramObj
}

func (r PublicMergeInputParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicMergeInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicMergeInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Contains an array of CRM object records to be processed in a batch operation,
// each defined by their ID and properties.
//
// The properties ID, Properties are required.
type SimplePublicObjectBatchInputParam struct {
	// The id to be updated. This can be the object id, or the unique property value of
	// the idProperty property
	ID string `json:"id" api:"required"`
	// Key-value pairs representing the properties of the object.
	Properties map[string]string `json:"properties,omitzero" api:"required"`
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

// An input object that contains a collection of objects to be created together in
// a batch.
//
// The properties Associations, Properties are required.
type SimplePublicObjectBatchInputForCreateParam struct {
	Associations []PublicAssociationsForObjectParam `json:"associations,omitzero" api:"required"`
	// Key value pairs representing the properties of the object.
	Properties map[string]string `json:"properties,omitzero" api:"required"`
	// An identifier used for tracing the creation or update request of the object.
	ObjectWriteTraceID param.Opt[string] `json:"objectWriteTraceId,omitzero"`
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
	ID string `json:"id" api:"required"`
	// Key value pairs representing the properties of the object.
	Properties map[string]string `json:"properties,omitzero" api:"required"`
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

// Contains the Id of a Public Object
//
// The property ID is required.
type SimplePublicObjectIDParam struct {
	// The unique ID of the object.
	ID string `json:"id" api:"required"`
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
	Properties map[string]string `json:"properties,omitzero" api:"required"`
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
// The properties Associations, Properties are required.
type SimplePublicObjectInputForCreateParam struct {
	Associations []PublicAssociationsForObjectParam `json:"associations,omitzero" api:"required"`
	// Key-value pairs for setting properties for the new object.
	Properties map[string]string `json:"properties,omitzero" api:"required"`
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
	ID string `json:"id" api:"required"`
	// Whether the object is archived.
	Archived bool `json:"archived" api:"required"`
	// The timestamp when the object was created, in ISO 8601 format.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Key value pairs representing the properties of the object.
	Properties map[string]string `json:"properties" api:"required"`
	// The timestamp when the object was last updated, in ISO 8601 format.
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// The timestamp when the object was archived, in ISO 8601 format.
	ArchivedAt time.Time `json:"archivedAt" format:"date-time"`
	// A list defining relationships with other objects.
	Associations map[string]CollectionResponseAssociatedID `json:"associations"`
	// An identifier used for tracing the creation or update request of the object.
	ObjectWriteTraceID string `json:"objectWriteTraceId"`
	// Key-value pairs representing the properties of the object along with their
	// history.
	PropertiesWithHistory map[string][]ValueWithTimestamp `json:"propertiesWithHistory"`
	// The URL on the API that provide direct navigation to the corresponding UI pages
	// for the connectors.
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		Archived              respjson.Field
		CreatedAt             respjson.Field
		Properties            respjson.Field
		UpdatedAt             respjson.Field
		ArchivedAt            respjson.Field
		Associations          respjson.Field
		ObjectWriteTraceID    respjson.Field
		PropertiesWithHistory respjson.Field
		URL                   respjson.Field
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
	ID string `json:"id" api:"required"`
	// Whether the object is archived.
	Archived bool `json:"archived" api:"required"`
	// The timestamp when the object was created, in ISO 8601 format.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Whether the property is new.
	New bool `json:"new" api:"required"`
	// Key value pairs representing the properties of the object.
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
		New                   respjson.Field
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
func (r SimplePublicUpsertObject) RawJSON() string { return r.JSON.raw }
func (r *SimplePublicUpsertObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
