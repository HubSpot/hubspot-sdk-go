// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/respjson"
	"github.com/stainless-sdks/hubspot-sdk-go/shared"
)

// ObjectDealSplitService contains methods and other services that help with
// interacting with the Hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObjectDealSplitService] method instead.
type ObjectDealSplitService struct {
	Options []option.RequestOption
}

// NewObjectDealSplitService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewObjectDealSplitService(opts ...option.RequestOption) (r ObjectDealSplitService) {
	r = ObjectDealSplitService{}
	r.Options = opts
	return
}

// Read a batch of deal split objects by their associated deal object internal ID
func (r *ObjectDealSplitService) BatchRead(ctx context.Context, body ObjectDealSplitBatchReadParams, opts ...option.RequestOption) (res *ObjectDealSplitBatchReadResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/objects/deals/splits/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Create or replace deal splits for deals with the provided IDs. Deal split
// percentages for each deal must sum up to 1.0 (100%) and may have up to 8 decimal
// places
func (r *ObjectDealSplitService) BatchUpsert(ctx context.Context, body ObjectDealSplitBatchUpsertParams, opts ...option.RequestOption) (res *ObjectDealSplitBatchUpsertResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "crm/v3/objects/deals/splits/batch/upsert"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ObjectDealSplitBatchReadResponse struct {
	CompletedAt time.Time                                `json:"completedAt,required" format:"date-time"`
	Results     []ObjectDealSplitBatchReadResponseResult `json:"results,required"`
	StartedAt   time.Time                                `json:"startedAt,required" format:"date-time"`
	// Any of "PENDING", "PROCESSING", "CANCELED", "COMPLETE".
	Status      ObjectDealSplitBatchReadResponseStatus `json:"status,required"`
	Links       map[string]string                      `json:"links"`
	RequestedAt time.Time                              `json:"requestedAt" format:"date-time"`
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
func (r ObjectDealSplitBatchReadResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectDealSplitBatchReadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectDealSplitBatchReadResponseResult struct {
	ID     string               `json:"id,required"`
	Splits []SimplePublicObject `json:"splits,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Splits      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectDealSplitBatchReadResponseResult) RawJSON() string { return r.JSON.raw }
func (r *ObjectDealSplitBatchReadResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectDealSplitBatchReadResponseStatus string

const (
	ObjectDealSplitBatchReadResponseStatusPending    ObjectDealSplitBatchReadResponseStatus = "PENDING"
	ObjectDealSplitBatchReadResponseStatusProcessing ObjectDealSplitBatchReadResponseStatus = "PROCESSING"
	ObjectDealSplitBatchReadResponseStatusCanceled   ObjectDealSplitBatchReadResponseStatus = "CANCELED"
	ObjectDealSplitBatchReadResponseStatusComplete   ObjectDealSplitBatchReadResponseStatus = "COMPLETE"
)

type ObjectDealSplitBatchUpsertResponse struct {
	CompletedAt time.Time                                  `json:"completedAt,required" format:"date-time"`
	Results     []ObjectDealSplitBatchUpsertResponseResult `json:"results,required"`
	StartedAt   time.Time                                  `json:"startedAt,required" format:"date-time"`
	// Any of "PENDING", "PROCESSING", "CANCELED", "COMPLETE".
	Status      ObjectDealSplitBatchUpsertResponseStatus `json:"status,required"`
	Links       map[string]string                        `json:"links"`
	RequestedAt time.Time                                `json:"requestedAt" format:"date-time"`
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
func (r ObjectDealSplitBatchUpsertResponse) RawJSON() string { return r.JSON.raw }
func (r *ObjectDealSplitBatchUpsertResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectDealSplitBatchUpsertResponseResult struct {
	ID     string               `json:"id,required"`
	Splits []SimplePublicObject `json:"splits,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Splits      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectDealSplitBatchUpsertResponseResult) RawJSON() string { return r.JSON.raw }
func (r *ObjectDealSplitBatchUpsertResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectDealSplitBatchUpsertResponseStatus string

const (
	ObjectDealSplitBatchUpsertResponseStatusPending    ObjectDealSplitBatchUpsertResponseStatus = "PENDING"
	ObjectDealSplitBatchUpsertResponseStatusProcessing ObjectDealSplitBatchUpsertResponseStatus = "PROCESSING"
	ObjectDealSplitBatchUpsertResponseStatusCanceled   ObjectDealSplitBatchUpsertResponseStatus = "CANCELED"
	ObjectDealSplitBatchUpsertResponseStatusComplete   ObjectDealSplitBatchUpsertResponseStatus = "COMPLETE"
)

type ObjectDealSplitBatchReadParams struct {
	BatchInputPublicObjectID shared.BatchInputPublicObjectIDParam
	paramObj
}

func (r ObjectDealSplitBatchReadParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPublicObjectID)
}
func (r *ObjectDealSplitBatchReadParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputPublicObjectID)
}

type ObjectDealSplitBatchUpsertParams struct {
	Inputs []ObjectDealSplitBatchUpsertParamsInput `json:"inputs,omitzero,required"`
	paramObj
}

func (r ObjectDealSplitBatchUpsertParams) MarshalJSON() (data []byte, err error) {
	type shadow ObjectDealSplitBatchUpsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectDealSplitBatchUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Splits are required.
type ObjectDealSplitBatchUpsertParamsInput struct {
	ID     int64                                        `json:"id,required"`
	Splits []ObjectDealSplitBatchUpsertParamsInputSplit `json:"splits,omitzero,required"`
	paramObj
}

func (r ObjectDealSplitBatchUpsertParamsInput) MarshalJSON() (data []byte, err error) {
	type shadow ObjectDealSplitBatchUpsertParamsInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectDealSplitBatchUpsertParamsInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties OwnerID, Percentage are required.
type ObjectDealSplitBatchUpsertParamsInputSplit struct {
	OwnerID    int64   `json:"ownerId,required"`
	Percentage float64 `json:"percentage,required"`
	paramObj
}

func (r ObjectDealSplitBatchUpsertParamsInputSplit) MarshalJSON() (data []byte, err error) {
	type shadow ObjectDealSplitBatchUpsertParamsInputSplit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectDealSplitBatchUpsertParamsInputSplit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
