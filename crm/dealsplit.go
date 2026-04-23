// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
)

// DealSplitService contains methods and other services that help with interacting
// with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDealSplitService] method instead.
type DealSplitService struct {
	options []option.RequestOption
	Batch   DealSplitBatchService
}

// NewDealSplitService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDealSplitService(opts ...option.RequestOption) (r DealSplitService) {
	r = DealSplitService{}
	r.options = opts
	r.Batch = NewDealSplitBatchService(opts...)
	return
}

type BatchResponseDealToDealSplits struct {
	// The timestamp indicating when the batch operation was completed, in date-time
	// format.
	CompletedAt time.Time `json:"completedAt" api:"required" format:"date-time"`
	// An array of deal-to-deal split objects representing the results of the batch
	// operation.
	Results []DealToDealSplits `json:"results" api:"required"`
	// The timestamp indicating when the batch operation started, in date-time format.
	StartedAt time.Time `json:"startedAt" api:"required" format:"date-time"`
	// The current status of the batch operation, with possible values: CANCELED,
	// COMPLETE, PENDING, PROCESSING.
	//
	// Any of "CANCELED", "COMPLETE", "PENDING", "PROCESSING".
	Status BatchResponseDealToDealSplitsStatus `json:"status" api:"required"`
	// A map of link names to associated URIs for additional resources or
	// documentation.
	Links map[string]string `json:"links"`
	// The timestamp indicating when the batch operation was requested, in date-time
	// format.
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
func (r BatchResponseDealToDealSplits) RawJSON() string { return r.JSON.raw }
func (r *BatchResponseDealToDealSplits) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current status of the batch operation, with possible values: CANCELED,
// COMPLETE, PENDING, PROCESSING.
type BatchResponseDealToDealSplitsStatus string

const (
	BatchResponseDealToDealSplitsStatusCanceled   BatchResponseDealToDealSplitsStatus = "CANCELED"
	BatchResponseDealToDealSplitsStatusComplete   BatchResponseDealToDealSplitsStatus = "COMPLETE"
	BatchResponseDealToDealSplitsStatusPending    BatchResponseDealToDealSplitsStatus = "PENDING"
	BatchResponseDealToDealSplitsStatusProcessing BatchResponseDealToDealSplitsStatus = "PROCESSING"
)

type DealToDealSplits struct {
	// The unique identifier for the deal associated with the deal splits.
	ID string `json:"id" api:"required"`
	// An array of deal split objects, each representing a portion of the deal assigned
	// to an owner.
	Splits []SimplePublicObject `json:"splits" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Splits      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DealToDealSplits) RawJSON() string { return r.JSON.raw }
func (r *DealToDealSplits) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties OwnerID, Percentage are required.
type PublicDealSplitInputParam struct {
	// The unique identifier of the owner receiving the deal split.
	OwnerID int64 `json:"ownerId" api:"required"`
	// The portion of the deal assigned to the owner, expressed as a percentage. The
	// total percentage for all splits in a deal must sum up to 1.0 (100%) and can have
	// up to 8 decimal places.
	Percentage float64 `json:"percentage" api:"required"`
	paramObj
}

func (r PublicDealSplitInputParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicDealSplitInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicDealSplitInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Inputs is required.
type PublicDealSplitsBatchCreateRequestParam struct {
	// An array of deal split inputs
	Inputs []PublicDealSplitsCreateRequestParam `json:"inputs,omitzero" api:"required"`
	paramObj
}

func (r PublicDealSplitsBatchCreateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicDealSplitsBatchCreateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicDealSplitsBatchCreateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Splits are required.
type PublicDealSplitsCreateRequestParam struct {
	// The unique identifier for the deal.
	ID int64 `json:"id" api:"required"`
	// An array of deal split inputs, each containing an owner ID and a percentage of
	// the deal split.
	Splits []PublicDealSplitInputParam `json:"splits,omitzero" api:"required"`
	paramObj
}

func (r PublicDealSplitsCreateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow PublicDealSplitsCreateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PublicDealSplitsCreateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
