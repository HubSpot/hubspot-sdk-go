// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package marketing

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"slices"

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
	Options []option.RequestOption
}

// NewCampaignBatchService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCampaignBatchService(opts ...option.RequestOption) (r CampaignBatchService) {
	r = CampaignBatchService{}
	r.Options = opts
	return
}

// Create a batch of campaigns with specified properties. This endpoint allows for
// the creation of multiple campaigns in a single request. Note that the 'hs_goal'
// property is deprecated and will be ignored if provided.
func (r *CampaignBatchService) New(ctx context.Context, body CampaignBatchNewParams, opts ...option.RequestOption) (res *BatchResponsePublicCampaign, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "marketing/campaigns/2026-03/batch/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update a batch of marketing campaigns with specified properties. This endpoint
// allows you to modify multiple campaigns in one request. Note that the 'hs_goal'
// property is deprecated and will be ignored if provided.
func (r *CampaignBatchService) Update(ctx context.Context, body CampaignBatchUpdateParams, opts ...option.RequestOption) (res *BatchResponsePublicCampaign, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "marketing/campaigns/2026-03/batch/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Archive a batch of marketing campaigns in your HubSpot account. This operation
// permanently removes the specified campaigns, making them inaccessible. It is
// useful for cleaning up outdated or unnecessary campaigns in bulk.
func (r *CampaignBatchService) Delete(ctx context.Context, body CampaignBatchDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "marketing/campaigns/2026-03/batch/archive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return err
}

// Retrieve a batch of campaigns with specified properties and date range. This
// endpoint allows you to filter campaigns by start and end dates and specify which
// properties to include in the response.
func (r *CampaignBatchService) Get(ctx context.Context, params CampaignBatchGetParams, opts ...option.RequestOption) (res *BatchResponsePublicCampaignWithAssets, err error) {
	opts = slices.Concat(r.Options, opts)
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
	return json.Unmarshal(data, &r.BatchInputPublicCampaignInput)
}

type CampaignBatchUpdateParams struct {
	BatchInputPublicCampaignBatchUpdateItem BatchInputPublicCampaignBatchUpdateItemParam
	paramObj
}

func (r CampaignBatchUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPublicCampaignBatchUpdateItem)
}
func (r *CampaignBatchUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputPublicCampaignBatchUpdateItem)
}

type CampaignBatchDeleteParams struct {
	BatchInputPublicCampaignDeleteInput BatchInputPublicCampaignDeleteInputParam
	paramObj
}

func (r CampaignBatchDeleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPublicCampaignDeleteInput)
}
func (r *CampaignBatchDeleteParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputPublicCampaignDeleteInput)
}

type CampaignBatchGetParams struct {
	BatchInputPublicCampaignReadInput BatchInputPublicCampaignReadInputParam
	// The end date for filtering campaigns, in YYYY-MM-DD format.
	EndDate param.Opt[string] `query:"endDate,omitzero" json:"-"`
	// The start date for filtering campaigns, in YYYY-MM-DD format.
	StartDate param.Opt[string] `query:"startDate,omitzero" json:"-"`
	// A comma-separated list of property names to include in the response.
	Properties []string `query:"properties,omitzero" json:"-"`
	paramObj
}

func (r CampaignBatchGetParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.BatchInputPublicCampaignReadInput)
}
func (r *CampaignBatchGetParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.BatchInputPublicCampaignReadInput)
}

// URLQuery serializes [CampaignBatchGetParams]'s query parameters as `url.Values`.
func (r CampaignBatchGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
