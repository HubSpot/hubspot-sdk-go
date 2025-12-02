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

// AssociationV4Service contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssociationV4Service] method instead.
type AssociationV4Service struct {
	Options []option.RequestOption
	Batch   AssociationV4BatchService
	Report  AssociationV4ReportService
}

// NewAssociationV4Service generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAssociationV4Service(opts ...option.RequestOption) (r AssociationV4Service) {
	r = AssociationV4Service{}
	r.Options = opts
	r.Batch = NewAssociationV4BatchService(opts...)
	r.Report = NewAssociationV4ReportService(opts...)
	return
}

// The property Inputs is required.
type BatchInputPublicAssociationMultiArchiveParam struct {
	Inputs []PublicAssociationMultiArchiveParam `json:"inputs,omitzero,required"`
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
	Inputs []PublicAssociationMultiPostParam `json:"inputs,omitzero,required"`
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
	Inputs []PublicDefaultAssociationMultiPostParam `json:"inputs,omitzero,required"`
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
	Inputs []PublicFetchAssociationsBatchRequestParam `json:"inputs,omitzero,required"`
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
	CompletedAt time.Time                 `json:"completedAt,required" format:"date-time"`
	Results     []LabelsBetweenObjectPair `json:"results,required"`
	// The timestamp when the batch processing began, in ISO 8601 format.
	StartedAt time.Time `json:"startedAt,required" format:"date-time"`
	// The status of the batch processing request: "PENDING", "PROCESSING",
	// "CANCELLED", or "COMPLETE".
	//
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status BatchResponseLabelsBetweenObjectPairStatus `json:"status,required"`
	Errors []shared.StandardError                     `json:"errors"`
	// An object containing relevant links related to the batch request.
	Links map[string]string `json:"links"`
	// The number of errors encountered during the batch processing.
	NumErrors int64 `json:"numErrors"`
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
	CompletedAt time.Time                         `json:"completedAt,required" format:"date-time"`
	Results     []PublicAssociationMultiWithLabel `json:"results,required"`
	// The timestamp when the batch processing began, in ISO 8601 format.
	StartedAt time.Time `json:"startedAt,required" format:"date-time"`
	// The status of the batch processing request: "PENDING", "PROCESSING", "CANCELED",
	// or "COMPLETE".
	//
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status BatchResponsePublicAssociationMultiWithLabelStatus `json:"status,required"`
	Errors []shared.StandardError                             `json:"errors"`
	// An object containing relevant links related to the batch request.
	Links map[string]string `json:"links"`
	// The number of errors encountered during the batch processing.
	NumErrors int64 `json:"numErrors"`
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

type DateTime struct {
	// Indicates whether the DateTime value represents only a date without a time
	// component.
	DateOnly bool `json:"dateOnly,required"`
	// The integer value representing the shift in minutes from UTC for the DateTime
	// value.
	TimeZoneShift int64 `json:"timeZoneShift,required"`
	// The integer value representing a specific point in time.
	Value int64 `json:"value,required"`
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

// The properties From, To are required.
type PublicAssociationMultiArchiveParam struct {
	From shared.PublicObjectIDParam   `json:"from,omitzero,required"`
	To   []shared.PublicObjectIDParam `json:"to,omitzero,required"`
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
	From  shared.PublicObjectIDParam    `json:"from,omitzero,required"`
	To    shared.PublicObjectIDParam    `json:"to,omitzero,required"`
	Types []shared.AssociationSpecParam `json:"types,omitzero,required"`
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
	From   shared.PublicObjectID            `json:"from,required"`
	To     []MultiAssociatedObjectWithLabel `json:"to,required"`
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

// The properties From, To are required.
type PublicDefaultAssociationMultiPostParam struct {
	From shared.PublicObjectIDParam `json:"from,omitzero,required"`
	To   shared.PublicObjectIDParam `json:"to,omitzero,required"`
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
	ID string `json:"id,required"`
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

type ReportCreationResponse struct {
	EnqueueTime DateTime `json:"enqueueTime,required"`
	// Email of the user
	UserEmail string `json:"userEmail,required"`
	// ID of the user
	UserID int64 `json:"userId,required"`
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
