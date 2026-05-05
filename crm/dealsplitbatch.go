// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package crm

import (
	"context"
	"net/http"
	"slices"

	"github.com/HubSpot/hubspot-sdk-go/internal/apijson"
	shimjson "github.com/HubSpot/hubspot-sdk-go/internal/encoding/json"
	"github.com/HubSpot/hubspot-sdk-go/internal/requestconfig"
	"github.com/HubSpot/hubspot-sdk-go/option"
	"github.com/HubSpot/hubspot-sdk-go/shared"
)

// DealSplitBatchService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDealSplitBatchService] method instead.
type DealSplitBatchService struct {
	options []option.RequestOption
}

// NewDealSplitBatchService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDealSplitBatchService(opts ...option.RequestOption) (r DealSplitBatchService) {
	r = DealSplitBatchService{}
	r.options = opts
	return
}

// Read a batch of deal split objects by their associated deal object internal ID
func (r *DealSplitBatchService) Read(ctx context.Context, body DealSplitBatchReadParams, opts ...option.RequestOption) (res *BatchResponseDealToDealSplits, err error) {
	opts = slices.Concat(r.options, opts)
	path := "deal-splits/2026-03/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Create or replace deal splits for deals with the provided IDs. Deal split
// percentages for each deal must sum up to 1.0 (100%) and may have up to 8 decimal
// places
func (r *DealSplitBatchService) Upsert(ctx context.Context, body DealSplitBatchUpsertParams, opts ...option.RequestOption) (res *BatchResponseDealToDealSplits, err error) {
	opts = slices.Concat(r.options, opts)
	path := "deal-splits/2026-03/batch/upsert"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type DealSplitBatchReadParams struct {
	BatchInputPublicObjectID shared.BatchInputPublicObjectIDParam
	paramObj
}

func (r DealSplitBatchReadParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPublicObjectID)
}
func (r *DealSplitBatchReadParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DealSplitBatchUpsertParams struct {
	PublicDealSplitsBatchCreateRequest PublicDealSplitsBatchCreateRequestParam
	paramObj
}

func (r DealSplitBatchUpsertParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PublicDealSplitsBatchCreateRequest)
}
func (r *DealSplitBatchUpsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
