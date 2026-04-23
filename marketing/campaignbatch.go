// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hubspot-sdk-go/internal/apijson"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/hubspot-sdk-go/internal/encoding/json"
	"github.com/stainless-sdks/hubspot-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/hubspot-sdk-go/option"
	"github.com/stainless-sdks/hubspot-sdk-go/packages/param"
)

// CampaignBatchService contains methods and other services that help with
// interacting with the hubspot API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCampaignBatchService] method instead.
type CampaignBatchService struct {
	options []option.RequestOption
}

// NewCampaignBatchService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCampaignBatchService(opts ...option.RequestOption) (r CampaignBatchService) {
	r = CampaignBatchService{}
	r.options = opts
	return
}

// This endpoint creates a batch of campaigns. The maximum number of items in a
// batch request is 50. The campaigns in the response are not guaranteed to be in
// the same order as they were provided in the request.
func (r *CampaignBatchService) New(ctx context.Context, body CampaignBatchNewParams, opts ...option.RequestOption) (res *BatchResponsePublicCampaign, err error) {
	opts = slices.Concat(r.options, opts)
	path := "marketing/campaigns/2026-03/batch/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// This endpoint updates a batch of campaigns based on the provided input data. The
// maximum number of items in a batch request is 50. If an empty string ("") is
// passed for any property in the Batch Update, it will reset that property's
// value.
func (r *CampaignBatchService) Update(ctx context.Context, body CampaignBatchUpdateParams, opts ...option.RequestOption) (res *BatchResponsePublicCampaign, err error) {
	opts = slices.Concat(r.options, opts)
	path := "marketing/campaigns/2026-03/batch/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// This endpoint deletes a batch of campaigns. The maximum number of items in a
// batch request is 50. The response will always be 204 No Content, regardless of
// whether the campaigns exist or not, whether they were successfully deleted or
// not, or if only some of the campaigns in the batch were deleted.
func (r *CampaignBatchService) Delete(ctx context.Context, body CampaignBatchDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "marketing/campaigns/2026-03/batch/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// This endpoint reads a batch of campaigns based on the provided input data and
// returns the campaigns along with their associated assets. The maximum number of
// items in a batch request is 50. The campaigns in the response are not guaranteed
// to be in the same order as they were provided in the request. If duplicate
// campaign IDs are provided in the request, duplicates will be ignored. The
// response will include only unique IDs and will be returned without duplicates.
func (r *CampaignBatchService) Get(ctx context.Context, params CampaignBatchGetParams, opts ...option.RequestOption) (res *BatchResponsePublicCampaignWithAssets, err error) {
	opts = slices.Concat(r.options, opts)
	path := "marketing/campaigns/2026-03/batch/read"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type CampaignBatchNewParams struct {
	BatchInputPublicCampaignInput BatchInputPublicCampaignInputParam
	paramObj
}

func (r CampaignBatchNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPublicCampaignInput)
}
func (r *CampaignBatchNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CampaignBatchUpdateParams struct {
	BatchInputPublicCampaignBatchUpdateItem BatchInputPublicCampaignBatchUpdateItemParam
	paramObj
}

func (r CampaignBatchUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPublicCampaignBatchUpdateItem)
}
func (r *CampaignBatchUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CampaignBatchDeleteParams struct {
	BatchInputPublicCampaignDeleteInput BatchInputPublicCampaignDeleteInputParam
	paramObj
}

func (r CampaignBatchDeleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPublicCampaignDeleteInput)
}
func (r *CampaignBatchDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CampaignBatchGetParams struct {
	BatchInputPublicCampaignReadInput BatchInputPublicCampaignReadInputParam
	EndDate                           param.Opt[string] `query:"endDate,omitzero" json:"-"`
	StartDate                         param.Opt[string] `query:"startDate,omitzero" json:"-"`
	Properties                        []string          `query:"properties,omitzero" json:"-"`
	paramObj
}

func (r CampaignBatchGetParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPublicCampaignReadInput)
}
func (r *CampaignBatchGetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [CampaignBatchGetParams]'s query parameters as `url.Values`.
func (r CampaignBatchGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
